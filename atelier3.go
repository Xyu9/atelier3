package atelier3

import (
	"math/rand"
)

func ArrayGenerate(nb int) int {
	randomNumber := rand.Intn(10000)
	result := randomNumber * nb
	return result
}
