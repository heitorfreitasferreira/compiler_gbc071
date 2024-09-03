package statemachine

import (
	"slices"
	"strings"
)

// Helper function para codificar o conjunto de estados em uma string única
func encodeStateSet(states []int) string {
	slices.Sort(states)
	result := ""
	for _, state := range states {
		result += string(rune(state)) + ","
	}
	return result
}

// Helper function para decodificar uma string de conjunto de estados
func decodeStateSet(encoded string) []int {
	stateStrings := strings.Split(encoded, ",")
	states := []int{}
	for _, stateStr := range stateStrings {
		if stateStr == "" {
			continue
		}
		states = append(states, int([]rune(stateStr)[0]))
	}
	return states
}
