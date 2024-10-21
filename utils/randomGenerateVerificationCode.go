package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomGenerateVerificationCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := 1
	for i := 0; i < length; i++ {
		start *= 10
	}
	end := start*10 - 1
	code := strconv.Itoa(r.Intn(end-start+1) + start)
	return code
}
