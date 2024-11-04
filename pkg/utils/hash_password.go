package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pas string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pas), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(pas, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pas))
	return err == nil
}
