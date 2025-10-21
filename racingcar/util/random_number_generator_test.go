package util

import "testing"

func TestGenerateRandomNumber(t *testing.T) {
	for i := 0; i < 100; i += 1 {

		n := GenerateRandomNumber()
		if n < MinRandomRange || n > MaxRandomRange {
			t.Errorf("Generated number %d out of range", n)
		}
	}
}
