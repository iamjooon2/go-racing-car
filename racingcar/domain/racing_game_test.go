package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRacingGame_WhenCarNameIsDuplicated_Failure(t *testing.T) {
	cars := []*Car{}
	car1, _ := NewCar("junhee")
	car2, _ := NewCar("junhee")
	car3, _ := NewCar("junhee")
	cars = append(cars, car1)
	cars = append(cars, car2)
	cars = append(cars, car3)

	trial, _ := NewTrial(3)

	generator := &FixedNumberGenerator{}

	game, err := NewRacingGame(cars, trial, generator)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}

func TestFindWinners(t *testing.T) {
	testCases := []struct {
		description string
		game        RacingGame
		expected    []*Car
	}{
		{
			description: "find multiple winners",
			game: RacingGame{
				Cars: []*Car{
					{name: "junhee", position: 1},
					{name: "moritz", position: 2},
					{name: "shaheen", position: 2},
				},
				Trial: &Trial{value: 1},
			},
			expected: []*Car{
				{name: "moritz", position: 2},
				{name: "shaheen", position: 2},
			},
		},
		{
			description: "find single winner",
			game: RacingGame{
				Cars: []*Car{
					{name: "junhee", position: 3},
					{name: "moritz", position: 2},
					{name: "shaheen", position: 1},
				},
				Trial: &Trial{value: 1},
			},
			expected: []*Car{
				{name: "junhee", position: 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := testCase.game.FindWinners()
			assert.EqualValues(t, testCase.expected, actual)
		})
	}
}
