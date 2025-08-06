package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ngodup/simplebank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

//Function it self is not a middleware it return a middleware function
//authMiddleware handles authentication and authorization for API requests.
// It is not a handler but a higher-order function that takes a token.Maker and returns a gin.HandlerFunc and return middleware function

func authMiddleware(tokenMAker token.Maker) gin.HandlerFunc {
	//this anonymous function is a authentication middleware
	return func(ctx *gin.Context) {
		//Step 1: Extract authorization header
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		//Step 2: Check type of authorization form it should be prefix with Bearer
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("unsupported authorization type %s" + authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		//Step 3: Get access token & verify access token
		accessToken := fields[1]
		payload, err := tokenMAker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		//Store 4: Store payload in gin context and pass it to the next handler
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
