package utils

import "golang.org/x/crypto/bcrypt"

// func CreatePasswordHash(password string) string {

// }

func CheckPasswordHash(password, hash string) bool {
    // Implement bcrypt hash comparison
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}