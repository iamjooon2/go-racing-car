package domain

import (
	"errors"
	"strconv"
)

const (
	MINIMUM_TRIAL_VALUE = 0
)

type Trial struct {
	value int
}

func NewTrial(value int) (*Trial, error) {
	if value <= MINIMUM_TRIAL_VALUE {
		return nil, errors.New("trial value must be greater than " + strconv.Itoa(MINIMUM_TRIAL_VALUE))
	}
	return &Trial{value: value}, nil
}

func (t *Trial) UseOneTime() {
	t.value--
}

func (t *Trial) IsRemain() bool {
	return t.value != 0
}

func (t *Trial) GetValue() int {
	return t.value
}
