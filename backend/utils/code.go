package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(9000) + 1000
	return fmt.Sprintf("SC%d", number)
}
