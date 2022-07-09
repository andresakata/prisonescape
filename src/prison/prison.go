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

func (p *Prison) FindBox(prisoner int) bool {
	nextBox := prisoner
	for attempts := 0; attempts < 50; attempts++ {
		currentBox := p.Boxes[nextBox]
		if currentBox == prisoner {
			return true
		}
		nextBox = currentBox
	}
	return false
}
