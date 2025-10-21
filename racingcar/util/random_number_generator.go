package util

import (
	"math/rand"
	"time"
)

const (
	MinRandomRange = 0
	MaxRandomRange = 9
)

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MinRandomRange + MaxRandomRange + 1)
}
