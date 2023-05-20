package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(passwordstr string) string {
	salt := 8
	password := []byte(passwordstr)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(hashstr, passStr []byte) bool {
	hash, pass := []byte(hashstr), []byte(passStr)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
	// true kalo pass nya sama
}
