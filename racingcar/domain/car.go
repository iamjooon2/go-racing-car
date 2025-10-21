package domain

import (
	"errors"
	"strconv"
)

const (
	INITIAL_POSITION = 0

	MAXIMUM_NAME_LENGTH = 7

	THROTTLE = 4
)

type Car struct {
	name     string
	position int
}

func NewCar(name string) (*Car, error) {
	err := validate(name)
	if err != nil {
		return nil, err
	}

	return &Car{
		name:     name,
		position: INITIAL_POSITION,
	}, nil
}

func validate(input string) error {
	if input == "" || len(input) > MAXIMUM_NAME_LENGTH {
		return errors.New("invalid car name input. length must be at most " + strconv.Itoa(MAXIMUM_NAME_LENGTH))
	}
	return nil
}

func (c *Car) Move(randomValue int) {
	if randomValue >= THROTTLE {
		c.position++
	}
}
