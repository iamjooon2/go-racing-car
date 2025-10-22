package domain

import (
	"errors"

	"github.com/poi1649/go-racing-car/racingcar/util"
)

type RacingGame struct {
	Trial *Trial
	Cars  []*Car
}

func NewRacingGame(cars []*Car, trial *Trial) (*RacingGame, error) {
	err := validateDuplicatedName(cars)
	if err != nil {
		return nil, err
	}
	return &RacingGame{
		Trial: trial,
		Cars:  cars,
	}, nil
}

func validateDuplicatedName(cars []*Car) error {
	// why here are no set data structure in golang?!
	var set = make(map[string]struct{})
	for _, car := range cars {
		set[car.Name()] = struct{}{} // struct use less memory than bool
		// value
	}
	if len(set) != len(cars) {
		return errors.New("duplicated car names")
	}
	return nil
}

func (r RacingGame) PlayOneTime() {
	for _, car := range r.Cars {
		randomInput := util.GenerateRandomNumber()
		car.Move(randomInput)
	}

	r.Trial.UseOneTime()
}

func (r RacingGame) FindWinners() []*Car {
	maxPosition := r.getMaxPosition()
	return r.getCarsPositionsAt(maxPosition)
}

func (r RacingGame) getCarsPositionsAt(maxPosition int) []*Car {
	winners := []*Car{}
	for _, car := range r.Cars {
		if car.position == maxPosition {
			winners = append(winners, car)
		}
	}
	return winners
}

func (r RacingGame) getMaxPosition() int {
	maxPosition := -1
	for _, car := range r.Cars {
		position := car.position
		if position > maxPosition {
			maxPosition = position
		}
	}
	return maxPosition
}

func (r RacingGame) GetCars() []*Car {
	return r.Cars
}
