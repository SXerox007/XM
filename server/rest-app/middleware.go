package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SXerox007/XM/utils/jwtauth"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: Load the RSA public key from a file or a key store
		key, _ := loadPublicKey("secure-keys/public_key.pem")

		// Parse and validate the token
		if !jwtauth.ParseToken(tokenString, key) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loadPublicKeyFromFile(filepath string) ([]byte, error) {
	key, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func loadPublicKey(path string) ([]byte, error) {
	key, err := loadPublicKeyFromFile(path)
	if err != nil {
		return nil, err
	}
	return key, nil
}
