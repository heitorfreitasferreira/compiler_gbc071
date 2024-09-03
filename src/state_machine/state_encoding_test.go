package statemachine

import "testing"

func TestEncodingAndDecodingStateList(t *testing.T) {
	testCases := []struct {
		desc    string
		states  []int
		encoded string
	}{
		{
			desc: "smol list",
			states: []int{
				0, 1, 2, 3, 4,
			},
			encoded: "\x00,\x01,\x02,\x03,\x04,",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			encoded := encodeStateSet(tC.states)
			if encoded != tC.encoded {
				t.Errorf("Expected %s, got %s", tC.encoded, encoded)
			}

			decoded := decodeStateSet(encoded)
			if len(tC.states) != len(decoded) {
				t.Errorf("Expected encoded string to have length %d, got %d", len(tC.states), len(decoded))
			}
			for i, state := range tC.states {
				if state != decoded[i] {
					t.Errorf("Expected decoded value to be %d, got %d", state, decoded[i])
				}
			}
		})
	}
}
