package statemachine

import "github.com/heitorfreitasferreira/compiler/types"

var globalAlphabet = map[byte]bool{
	'a': true,
	'A': true,
	'b': true,
	'B': true,
	'c': true,
	'C': true,
	'd': true,
	'D': true,
	'e': true,
	'E': true,
	'f': true,
	'F': true,
	'g': true,
	'G': true,
	'h': true,
	'H': true,
	'i': true,
	'I': true,
	'j': true,
	'J': true,
	'k': true,
	'K': true,
	'l': true,
	'L': true,
	'm': true,
	'M': true,
	'n': true,
	'N': true,
	'o': true,
	'O': true,
	'p': true,
	'P': true,
	'q': true,
	'Q': true,
	'r': true,
	'R': true,
	's': true,
	'S': true,
	't': true,
	'T': true,
	'u': true,
	'U': true,
	'v': true,
	'V': true,
	'w': true,
	'W': true,
	'x': true,
	'X': true,
	'y': true,
	'Y': true,
	'z': true,
	'Z': true,

	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'.': true,

	',': true,
	';': true,

	'(': true,
	')': true,

	'\n': true,
	'\t': true,
	' ':  true,

	'+': true,
	'-': true,
	'*': true,
	'/': true,

	'=': true,
	':': true,

	'{': true,
	'}': true,

	'>': true,
	'<': true,
	'!': true,
}

// Pelo amor de deus não pode ser um caractere do alfabeto
const emptyTransition byte = '$'

const notInAlphabet = -1

var Digit []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var Letter []byte = []byte{'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F', 'g', 'G', 'h', 'H', 'i', 'I', 'j', 'J', 'k', 'K', 'l', 'L', 'm', 'M', 'n', 'N', 'o', 'O', 'p', 'P', 'q', 'Q', 'r', 'R', 's', 'S', 't', 'T', 'u', 'U', 'v', 'V', 'w', 'W', 'x', 'X', 'y', 'Y', 'z', 'Z'}
var DigitOrLetter []byte = append(Digit, Letter...)

func GetTransition(positives []types.Tuple[byte, int], other ...int) []int {
	max := 256

	transition := make([]int, max)
	for i := 0; i < max; i++ {
		transition[i] = notInAlphabet
	}

	for i := 0; i < len(positives); i++ {
		transition[positives[i].First] = positives[i].Second
	}

	if len(other) == 1 {
		for a := range globalAlphabet {
			if transition[a] == notInAlphabet {
				transition[a] = other[0]
			}
		}
	}
	return transition
}
