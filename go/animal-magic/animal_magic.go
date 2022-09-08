package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	min := 1
	max := 20
	return rand.Intn(max-min+1) + min
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	min := float64(0)
	max := float64(12)
	return min + rand.Float64()*(max-min)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	x := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})

	return x
}
