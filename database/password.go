package database

import "golang.org/x/crypto/bcrypt"

// VerifyPasswordForClient compares password and the hashed password
func VerifyPasswordForClient(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

// HashPasswordForClient creates a bcrypt hash
func HashPasswordForClient(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 3)
}
