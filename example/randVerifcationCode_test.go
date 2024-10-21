package example

import (
	"fmt"
	"goPro4/utils"
	"testing"
)

func TestRandVerificationCode(t *testing.T) {
	code := utils.RandomGenerateVerificationCode(5)
	fmt.Println(code)
}
