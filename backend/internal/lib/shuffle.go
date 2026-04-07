package lib

import (
	"math/rand/v2"
	"slices"
)

func shuffle() string {
	keys := "abdefghijklmnopqrstuvwxyz012456789"
	var usedKeysIndex []int
	scrambledKeys := ""
	for range 7 {
		random := rand.IntN(len(keys))
		for {
			if slices.Contains(usedKeysIndex, random) {
				random = rand.IntN(len(keys))
			} else {
				break
			}
		}
		usedKeysIndex = append(usedKeysIndex, random)
		scrambledKeys += string(keys[random])
	}
	return scrambledKeys
}
