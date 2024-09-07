package statemachine

import (
	"slices"
	"strings"
)

// AppendMultipleSlices appends multiple slices into one.
func appendMultipleSlices(slices ...[]byte) []byte {
	var result []byte
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// CreateByteRange creates a byte array with bytes between start and end, inclusive.
func createByteRange(start, end byte) []byte {
	var result []byte
	for b := start; b <= end; b++ {
		result = append(result, b)
	}
	return result
}

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
