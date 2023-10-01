package helpers

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("an7cryu9q85n89032ncl1awjqbpzc")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
