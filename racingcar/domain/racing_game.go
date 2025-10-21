package domain

import "github.com/poi1649/go-racing-car/racingcar/util"

type RacingGame struct {
	Trial *Trial
	Cars  []*Car
}

func NewRacingGame(cars []*Car, trial *Trial) RacingGame {
	return RacingGame{
		Trial: trial,
		Cars:  cars,
	}
}

func (r RacingGame) PlayOneTime() {
	randomInput := util.GenerateRandomNumber()

	for _, car := range r.Cars {
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
