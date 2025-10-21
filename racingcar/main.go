package main

import (
	"fmt"
	"os"

	"github.com/poi1649/go-racing-car/racingcar/controller"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program terminated due to error:", r)
			os.Exit(1)
		}
	}()

	controller.Run()
}
