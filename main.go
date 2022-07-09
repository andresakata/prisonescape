package main

import (
	"fmt"
	"prisonescape/src/prison"
)

func main() {
	distribution := map[bool]int{true: 0, false: 0}
	for i := 0; i < 100; i++ {
		distribution[run()]++
	}
	fmt.Println(distribution)
}

func run() bool {
	prisonInstance := prison.NewPrison()
	for i := 0; i < 100; i++ {
		if !prisonInstance.FindBox(i) {
			return false
		}
	}
	return true
}
