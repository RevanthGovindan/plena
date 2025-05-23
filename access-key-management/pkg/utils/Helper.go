package utils

import (
	"math/rand"
	"time"
)

// possible collision, try better approach
func GenerateRandom() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(100)
}
