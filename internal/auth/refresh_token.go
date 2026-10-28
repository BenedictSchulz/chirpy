package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	randomBytes := make([]byte, 32)

	rand.Read(randomBytes)

	refreshToken := hex.EncodeToString(randomBytes)

	return refreshToken

}
