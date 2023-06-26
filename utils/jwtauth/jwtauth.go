package jwtauth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	EXP = "exp"
	SUB = "sub"
)

var secretKey = []byte("/7fjzEeQOd762BnRCDSL3QDCzhlqq4gSAnlnAdYSPRY=")

// Generate the Access token
func GenerateToken(username string, uuid string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims[EXP] = time.Now().Add(time.Minute * 1).Unix()
	claims[SUB] = uuid
	token.Claims = claims
	return token.SignedString(secretKey)
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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
