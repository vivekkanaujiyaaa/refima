package password

import "golang.org/x/crypto/bcrypt"

// Verify compares password and the hashed password
func Verify(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

// Hash creates a bcrypt hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 3)
}
