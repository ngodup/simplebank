package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer within the given range.
func RandomInt(min, max int) int {
	return min + r.Intn(max-min+1)
}

// RandomString generates a random string of the given length.
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name of length 6.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money.
func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

// RandomCurrency generates a random currency code.
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP"}
	n := len(currencies)
	return currencies[r.Intn(n)]
}
