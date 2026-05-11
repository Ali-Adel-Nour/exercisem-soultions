package chance

import (
    "math/rand"
    "time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    return rand.Intn(20) + 1 // Generate a random number between 1 and 20
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    return rand.Float64() * 12.0 // Generate a random float between 0.0 and 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    animals := []string{
        "ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog",
    }
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    rand.Shuffle(len(animals), func(i, j int) {
        animals[i], animals[j] = animals[j], animals[i] // Shuffle the animals
    })
    return animals
}