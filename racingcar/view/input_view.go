package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNames() []string {
	fmt.Println("Enter the names of the cars to race. (Names are separated by commas)")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text() // eliminate entering new lines
		names := strings.Split(input, ",")
		for i := range names {
			names[i] = strings.TrimSpace(names[i])
		}
		return names
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	panic("Failed to read names")
}

func ReadAttempts() int {
	fmt.Println("How many attempts?")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text() // eliminate entering new lines
		attempts, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}
		return attempts
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	panic("Failed to read attempts")
}
