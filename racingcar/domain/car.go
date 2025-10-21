package domain

import (
	"errors"
	"strconv"
)

const (
	InitialPosition = 0

	MaximumNameLength = 8

	Throttle = 4
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
		position: InitialPosition,
	}, nil
}

func validate(input string) error {
	if input == "" || len(input) > MaximumNameLength {
		return errors.New("invalid car name input. name length must be at smaller than " + strconv.Itoa(MaximumNameLength))
	}
	return nil
}

func (c *Car) Move(randomValue int) {
	if randomValue >= Throttle {
		c.position++
	}
}

func (c *Car) Name() string {
	return c.name
}

func (c *Car) Position() int {
	return c.position
}
