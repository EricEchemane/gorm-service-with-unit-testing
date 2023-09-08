package identity

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	const cost = 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(plainText, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	return err == nil
}
