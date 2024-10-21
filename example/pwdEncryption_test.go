package example

import (
	"fmt"
	"goPro4/utils"
	"testing"
)

func TestPwdEncryption(t *testing.T) {
	word, err := utils.EncryptionPassWord("123456", "xiaoming")
	fmt.Println(err)
	result := utils.CheckPassword(word, "12456")
	fmt.Println(result)
}
