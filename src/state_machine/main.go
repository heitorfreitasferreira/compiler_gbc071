package statemachine

import "github.com/heitorfreitasferreira/compiler/types"

// TODO: Quando o diagrama final estiver pronto, implementa-lo aqui
var DefaultDFA *DFA = NewDFA(
	[][]int{
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: []byte{'<'}, Second: 1},
			{First: []byte{'>'}, Second: 2},
			{First: []byte{'='}, Second: 3},
			{First: []byte{'!'}, Second: 4},
			{First: []byte{' '}, Second: 5},
			{First: []byte{'\n', '\t'}, Second: 6},
			{First: []byte{':'}, Second: 7},
			{First: []byte{')'}, Second: 8},
			{First: []byte{'('}, Second: 9},
			{First: []byte{','}, Second: 10},
			{First: []byte{';'}, Second: 11},
			{First: []byte{'*'}, Second: 12},
			{First: []byte{'/'}, Second: 13},
			{First: []byte{'+', '-'}, Second: 14},
			{First: []byte{'\''}, Second: 15},
			{First: Digit, Second: 16},
			{
				First: append(
					[]byte{'a', 'g', 'h', 'j', 'k', 'l', 'n', 'o', 'p', 'q', 's', 'v', 'w', 'x', 'y', 'z', '_'},
					UppercaseLetter...,
				),
				Second: 17,
			},
			{First: []byte{'i'}, Second: 18},
			{First: []byte{'f'}, Second: 19},
			{First: []byte{'c'}, Second: 20},
			{First: []byte{'b'}, Second: 21},
			{First: []byte{'e'}, Second: 23},
			{First: []byte{'m'}, Second: 24},
			{First: []byte{'d'}, Second: 25},
			{First: []byte{'t'}, Second: 26},
			{First: []byte{'w'}, Second: 27},
			{First: []byte{'r'}, Second: 28},
			{First: []byte{'u'}, Second: 29},
			// TODO: Add other states here when available
		}), // 0
		GetTransition([]types.Tuple[byte, int]{{First: '=', Second: 30}}, 31),             // 1
		GetTransition([]types.Tuple[byte, int]{{First: '=', Second: 32}}, 31),             // 2
		GetTransition([]types.Tuple[byte, int]{{First: '=', Second: 33}}),                 // 3
		GetTransition([]types.Tuple[byte, int]{{First: '=', Second: 34}}),                 // 4
		GetTransition([]types.Tuple[byte, int]{}),                                         // 5
		GetTransition([]types.Tuple[byte, int]{}),                                         // 6
		GetTransition([]types.Tuple[byte, int]{{First: '=', Second: 36}}, 35),             // 7
		GetTransition([]types.Tuple[byte, int]{}),                                         // 8
		GetTransition([]types.Tuple[byte, int]{}),                                         // 9
		GetTransition([]types.Tuple[byte, int]{}),                                         // 10
		GetTransition([]types.Tuple[byte, int]{}),                                         // 11
		GetTransition([]types.Tuple[byte, int]{{First: '*', Second: 37}}, 38),             // 12
		GetTransition([]types.Tuple[byte, int]{}),                                         // 13
		GetTransition([]types.Tuple[byte, int]{}),                                         // 14
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{{First: Letter, Second: 39}}), // 15
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: Digit, Second: 16},
			{First: []byte{'.'}, Second: 40},
			{First: []byte{'E'}, Second: 56},
		}, 41), // 16
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{{First: DigitOrLetterOrUnderscore, Second: 17}}, 99), // 17
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'e'),
				createByteRange('g', 'm'),
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			), Second: 17},
			{First: []byte{'n'}, Second: 42},
			{First: []byte{'f'}, Second: 43},
		}, 41), // 18
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'k'),
				createByteRange('m', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			), Second: 17},
			{First: []byte{'l'}, Second: 44},
		}, 99), // 19
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'g'),
				createByteRange('i', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'h'}, Second: 45},
		}, 99), // 20
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'b'),
				createByteRange('d', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 46},
		}, 99), // 21
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'r'),
				createByteRange('t', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'s'}, Second: 47},
		}, 99), // 22
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'k'),
				[]byte{'m'},
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'l'}, Second: 22},
			{First: []byte{'n'}, Second: 48},
		}, 99), // 23
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('b', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'a'}, Second: 49},
		}, 99), // 24
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('p', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'o'}, Second: 50},
		}, 99), // 25
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'g'),
				createByteRange('i', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'h'}, Second: 51},
		}, 99), // 26
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('p', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'h'}, Second: 52},
		}, 99), // 27
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'd'),
				createByteRange('f', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 53},
		}, 99), // 28
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'n'}, Second: 54},
		}, 99), // 29
		GetTransition([]types.Tuple[byte, int]{}),                          // 30
		GetTransition([]types.Tuple[byte, int]{}),                          // 31
		GetTransition([]types.Tuple[byte, int]{}),                          // 32
		GetTransition([]types.Tuple[byte, int]{}),                          // 33
		GetTransition([]types.Tuple[byte, int]{}),                          // 34
		GetTransition([]types.Tuple[byte, int]{}),                          // 35
		GetTransition([]types.Tuple[byte, int]{}),                          // 36
		GetTransition([]types.Tuple[byte, int]{}),                          // 37
		GetTransition([]types.Tuple[byte, int]{}),                          // 38
		GetTransition([]types.Tuple[byte, int]{{First: '\'', Second: 55}}), // 39
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: Digit, Second: 40},
			{First: []byte{'E'}, Second: 56},
		}, 41), // 40
		GetTransition([]types.Tuple[byte, int]{}), // 41
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 's'),
				createByteRange('u', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'t'}, Second: 57},
		}, 99), // 42
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 58), // 43
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'n'),
				createByteRange('p', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'o'}, Second: 59},
		}, 99), // 44
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('b', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'a'}, Second: 60},
		}, 99), // 45
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'f'),
				createByteRange('h', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'g'}, Second: 61},
		}, 99), // 46
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'd'),
				createByteRange('f', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 62},
		}, 99), // 47
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'c'),
				createByteRange('e', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'d'}, Second: 63},
		}, 99), // 48
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'h'),
				createByteRange('j', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'i'}, Second: 64},
		}, 99), // 49
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 65), // 50
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'd'),
				createByteRange('f', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 66},
		}, 99), // 51
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'h'),
				createByteRange('j', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'i'}, Second: 67},
		}, 99), // 52
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'o'),
				createByteRange('q', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'p'}, Second: 68},
		}, 99), // 53
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 's'),
				createByteRange('u', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'t'}, Second: 69},
		}, 99), // 54
		GetTransition([]types.Tuple[byte, int]{}), // 55
		GetTransition([]types.Tuple[byte, int]{
			{First: '+', Second: 70},
			{First: '-', Second: 71},
		}), //56
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 72), // 57
		GetTransition([]types.Tuple[byte, int]{}), // 58
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('b', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'a'}, Second: 73},
		}, 99), // 59
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'q'),
				createByteRange('s', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'r'}, Second: 74},
		}, 99), // 60
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'h'),
				createByteRange('j', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'i'}, Second: 75},
		}, 99), // 61
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 76), // 62
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 77), // 63
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'n'}, Second: 78},
		}, 99), // 64
		GetTransition([]types.Tuple[byte, int]{}), // 65
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'n'}, Second: 79},
		}, 99), // 66
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'k'),
				createByteRange('m', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'l'}, Second: 80},
		}, 99), // 67
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'd'),
				createByteRange('f', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 81},
		}, 99), // 68
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'h'),
				createByteRange('j', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'i'}, Second: 82},
		}, 99), // 69
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: Digit, Second: 83},
		}), // 70
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: Digit, Second: 83},
		}), // 71
		GetTransition([]types.Tuple[byte, int]{}), // 72
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 's'),
				createByteRange('u', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'t'}, Second: 84},
		}, 99), // 73
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 85), // 74
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'm'),
				createByteRange('o', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'n'}, Second: 86},
		}, 99), // 75
		GetTransition([]types.Tuple[byte, int]{}), // 76
		GetTransition([]types.Tuple[byte, int]{}), // 77
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 87), // 78
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 88), // 79
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'd'),
				createByteRange('f', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'e'}, Second: 89},
		}, 99), // 80
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('b', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'a'}, Second: 90},
		}, 99), // 81
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 'k'),
				createByteRange('m', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'l'}, Second: 91},
		}, 99), // 82
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{{First: Digit, Second: 83}}, 92), // 83
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 93), // 84
		GetTransition([]types.Tuple[byte, int]{}), // 85
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 94), // 86
		GetTransition([]types.Tuple[byte, int]{}), // 87
		GetTransition([]types.Tuple[byte, int]{}), // 88
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 95), // 89
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: appendMultipleSlices(
				createByteRange('a', 's'),
				createByteRange('u', 'z'),
				UppercaseLetterOrDigitOrUnderscore,
			),
				Second: 17},
			{First: []byte{'t'}, Second: 96},
		}, 99), // 90
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 97), // 91
		GetTransition([]types.Tuple[byte, int]{}), // 92
		GetTransition([]types.Tuple[byte, int]{}), // 93
		GetTransition([]types.Tuple[byte, int]{}), // 94
		GetTransition([]types.Tuple[byte, int]{}), // 95
		GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
			{First: DigitOrLetterOrUnderscore, Second: 17},
		}, 98), // 96
		GetTransition([]types.Tuple[byte, int]{}), // 97
		GetTransition([]types.Tuple[byte, int]{}), // 98
		GetTransition([]types.Tuple[byte, int]{}), // 99
	},
	map[int]types.Tuple[types.TokenType, bool]{
		5:  {First: types.SEPARATOR, Second: false},
		6:  {First: types.SEPARATOR, Second: false},
		8:  {First: types.END_PAREN, Second: false},
		9:  {First: types.START_PAREN, Second: false},
		10: {First: types.KKOMA, Second: false},
		11: {First: types.SEMICOLON, Second: false},
		13: {First: types.ARIOP_MULT, Second: false},
		14: {First: types.ARIOP_SUM, Second: false},
		30: {First: types.RELOP, Second: false},
		31: {First: types.RELOP, Second: false},
		32: {First: types.RELOP, Second: false},
		33: {First: types.RELOP, Second: false},
		34: {First: types.RELOP, Second: false},
		35: {First: types.TYPE_SEPARATOR, Second: false},
		36: {First: types.ASSIGN, Second: false},
		37: {First: types.ARIOP_POW, Second: false},
		38: {First: types.ARIOP_MULT, Second: true},
		41: {First: types.CONST, Second: false},
		55: {First: types.CONST, Second: false},
		58: {First: types.KW_IF, Second: false},
		65: {First: types.KW_DO, Second: false},
		72: {First: types.KW_TYPE, Second: false},
		76: {First: types.KW_ELSE, Second: false},
		77: {First: types.KW_END, Second: false},
		87: {First: types.KW_MAIN, Second: false},
		88: {First: types.KW_THEN, Second: false},
		95: {First: types.KW_WHILE, Second: false},
		97: {First: types.KW_UNTIL, Second: false},
		98: {First: types.KW_REPEAT, Second: false},
		99: {First: types.IDENTIFIER, Second: false},
	},
)
