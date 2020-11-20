package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Check password compare password and hash
func CheckPassword(stringHash, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(stringHash), []byte(stringPassword))
}
