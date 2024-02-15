package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"

	"testing"
)

// 예시 테스트 함수, 구현 시 제거
func TestHello(t *testing.T) {
	// given
	expected := "Hello, Racing Car!"

	// when
	actual := racingcar.Hello()

	// then
	assert.Equal(
		t,
		expected,
		actual,
		"they should be equal",
	)
}
