package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedNumberGenerator_Generate(t *testing.T) {
	fixedNumberGenerator := &FixedNumberGenerator{}

	generatedNumber := fixedNumberGenerator.Generate()

	assert.Equal(t, generatedNumber, fixedNumberGenerator.Generate())
}
