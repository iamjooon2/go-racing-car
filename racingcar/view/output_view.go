package view

import (
	"fmt"
	"strings"

	"github.com/poi1649/go-racing-car/racingcar/domain"
)

const (
	MOVE_UNIT               = "-"
	WINNER_NAME_DELIMITER   = ","
	NAME_POSITION_DELIMITER = ":"
)

func PrintRace(cars []*domain.Car) {
	for _, car := range cars {
		position := car.Position()
		name := car.Name()

		fmt.Println(name + NAME_POSITION_DELIMITER + strings.Repeat(MOVE_UNIT, position))
	}
}

func PrintWinners(winningCars []*domain.Car) {
	fmt.Print("Find winners: ")

	var winners []string
	for _, winCar := range winningCars {
		name := winCar.Name()
		winners = append(winners, name)
	}

	fmt.Println(strings.Join(winners, WINNER_NAME_DELIMITER))
}
