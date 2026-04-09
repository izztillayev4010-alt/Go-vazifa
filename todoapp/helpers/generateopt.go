package helpers

import (
	"fmt"
	"math/rand"
)

func GenerateOTP() string {
	code := rand.Intn(10000)
	return fmt.Sprintf("%04d", code)
}
