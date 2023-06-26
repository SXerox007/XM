package jwtauth

import (
	"crypto/rsa"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	EXP = "exp"
	SUB = "sub"
)

// Generate the Access token
func GenerateToken(key *rsa.PublicKey, username string, uuid string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	claims["sub"] = uuid
	token.Claims = claims

	// Sign the token with the RSA private key
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logout the user
func Logout(tokenString string, token *jwt.Token) error {
	return nil
}

// Blacklist the token can't access
func BlackListToken(tokenString string) error {
	return nil
}

// parse the token
func ParseToken(myToken string, key []byte) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
