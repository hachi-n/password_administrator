package util

import (
	"math/rand"
	"time"
)

func Shuffle(slice []int) {
	n := len(slice)
	rand.Seed(time.Now().UnixNano())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
