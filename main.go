package main

import (
	"fmt"
	"prisonescape/src/prison"
)

func main() {
	fmt.Println("Hello, World!")
	prisonInstance := prison.NewPrison()
	fmt.Println(prisonInstance.Boxes)
}
