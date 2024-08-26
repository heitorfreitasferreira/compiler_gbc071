package statemachine

var alphabet = map[byte]bool{
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

func addNegationTransitions(oldTransitionmap map[byte]int, negationChars []byte, negationState int) map[byte]int {
	clone := make(map[byte]int)
	for k, v := range oldTransitionmap {
		clone[k] = v
	}

	for a := range alphabet {
		for i := 0; i < len(negationChars); i++ {
			if negationChars[i] == a {
				continue
			}
			clone[a] = negationState
		}
	}
	return clone
}
