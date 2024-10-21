package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptionPassWord(pwd, uname string) (string, error) {
	encryptionPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(encryptionPwd), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
