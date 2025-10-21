package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCar_Success(t *testing.T) {
	validNames := []string{"junhee", "moritz", "ali", "shaheen"}

	for _, validName := range validNames {
		car, _ := NewCar(validName)

		assert.NotNil(t, car)
	}
}

func TestNewCar_Fail(t *testing.T) {
	invalidNames := []string{"", "this is the invalid length name"}

	for _, invalidName := range invalidNames {
		_, err := NewCar(invalidName)

		assert.NotNil(t, err)
	}
}

func TestMove_WhenInputIsGreaterThanOrEqualToThrottle_PositionIncreases(t *testing.T) {
	inputs := []int{4, 5, 6, 7, 8, 9}

	car, _ := NewCar("junhee")
	for _, input := range inputs {
		car.Move(input)
	}

	assert.Equal(t, car.position, len(inputs))
}

func TestMove_WhenInputIsSmallerThanThrottle_PositionNotChange(t *testing.T) {
	inputs := []int{0, 1, 2, 3}

	car, _ := NewCar("junhee")
	for _, input := range inputs {
		car.Move(input)
	}

	assert.Zero(t, car.position)
}
