package controller

import (
	"fmt"

	"github.com/poi1649/go-racing-car/racingcar/domain"
	"github.com/poi1649/go-racing-car/racingcar/view"
)

func Run() {
	cars, err := generateCars(view.ReadNames())
	if err != nil {
		fmt.Println(err)
		return
	}
	trial, err := generateTrial(view.ReadAttempts())
	if err != nil {
		fmt.Println(err)
		return
	}

	game, err := domain.NewRacingGame(cars, trial, &domain.RandomNumberGenerator{})
	if err != nil {
		fmt.Println(err)
		return
	}

	view.PrintRaceResult()
	for trial.IsRemain() {
		game.PlayOneTime()

		view.PrintRace(game.GetCars())
		view.PrintEnter()
	}

	view.PrintWinners(game.FindWinners())
}

func generateTrial(inputAttempts int) (*domain.Trial, error) {
	trial, err := domain.NewTrial(inputAttempts)
	if err != nil {
		return nil, err
	}
	return trial, nil
}

func generateCars(inputNames []string) ([]*domain.Car, error) {
	var cars []*domain.Car
	for _, inputName := range inputNames {
		car, err := domain.NewCar(inputName)
		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}
	return cars, nil
}
