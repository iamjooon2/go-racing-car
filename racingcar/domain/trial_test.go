package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrial_Success(t *testing.T) {
	values := []int{1, 2, 3, 4, 100}

	for _, value := range values {
		trial, _ := NewTrial(value)

		assert.NotNil(t, trial)
	}
}

func TestNewTrial_Fail(t *testing.T) {
	values := []int{0, -1, -999}

	for _, value := range values {
		_, err := NewTrial(value)

		assert.NotNil(t, err)
	}
}

func TestTrial_UseOneTime(t *testing.T) {
	value := 1

	trial, _ := NewTrial(value)
	trial.UseOneTime()

	assert.Equal(t, trial.value, value-1)
}
