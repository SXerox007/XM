package jwtauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndVerifyToken(t *testing.T) {
	username := "testuser"
	uuid := "123456789"

	// Generate the token
	tokenString, err := GenerateToken(username, uuid)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Verify the token
	token, err := VerifyToken(tokenString)
	assert.NoError(t, err)
	assert.True(t, token.Valid)
}

func TestParseToken(t *testing.T) {
	username := "testuser"
	uuid := "123456789"

	// Generate the token
	tokenString, err := GenerateToken(username, uuid)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse and validate the token
	valid := ParseToken(tokenString, secretKey)
	assert.True(t, valid)
}
