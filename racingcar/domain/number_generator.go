package domain

import (
	"math/rand"
	"time"
)

const (
	MinRandomRange = 0
	MaxRandomRange = 9
)

type NumberGenerator interface {
	Generate() int
}

type RandomNumberGenerator struct{}

type FixedNumberGenerator struct{}

type IncreasingNumberGenerator struct {
	value int
}

func (r *RandomNumberGenerator) Generate() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MinRandomRange + MaxRandomRange + 1)
}

func (f *FixedNumberGenerator) Generate() int {
	return 1
}

func (i *IncreasingNumberGenerator) Generate() int {
	i.value++
	return i.value
}
