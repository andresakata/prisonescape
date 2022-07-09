package prison

import (
	"math/rand"
	"time"
)

type Prison struct {
	Boxes map[int]int
}

func NewPrison() *Prison {
	rand.Seed(time.Now().UnixNano())
	prison := Prison{Boxes: shuffleBoxes()}
	return &prison
}

func shuffleBoxes() map[int]int {
	var numbers []int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}

	boxes := make(map[int]int)
	for i := 0; i < 100; i++ {
		index := rand.Intn(len(numbers))
		boxes[i] = numbers[index]
		numbers = append(numbers[:index], numbers[index+1:]...)
	}

	return boxes
}
