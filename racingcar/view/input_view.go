package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadNames() []string {
	fmt.Println("Enter the names of the cars to race. (Names are separated by commas)")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(input, ",")
}

func ReadAttempts() int {
	fmt.Println("How many attempts?")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	attempts, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
	}

	return attempts
}
