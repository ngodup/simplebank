package token

import "time"

// Maker is an interface for creating and verifying tokens, allow easily switching between diff type token maker
type Maker interface {
	//Create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	//Check if input token is valid or not
	VerifyToken(token string) (*Payload, error)
}