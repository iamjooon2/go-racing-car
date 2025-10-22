package controller

import (
	"fmt"

	"github.com/poi1649/go-racing-car/racingcar/domain"
	"github.com/poi1649/go-racing-car/racingcar/view"
)

func Run() {
	cars := generateCars(view.ReadNames())
	trial := generateTrial(view.ReadAttempts())

	game, err := domain.NewRacingGame(cars, trial)
	if err != nil {
		fmt.Println(err)
	}

	view.PrintRaceResult()
	for trial.IsRemain() {
		game.PlayOneTime()

		view.PrintRace(game.GetCars())
		view.PrintEnter()
	}

	view.PrintWinners(game.FindWinners())
}

func generateTrial(inputAttempts int) *domain.Trial {
	trial, err := domain.NewTrial(inputAttempts)
	if err != nil {
		fmt.Println(err)
	}
	return trial
}

func generateCars(inputNames []string) []*domain.Car {
	var cars []*domain.Car
	for _, inputName := range inputNames {
		car, err := domain.NewCar(inputName)
		if err != nil {
			fmt.Println(err)
		}

		cars = append(cars, car)
	}
	return cars
}
