package myBufferedByteReader_test

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"testing"

	"github.com/heitorfreitasferreira/compiler/myBufferedByteReader"
)

func TestMyByteReader(t *testing.T) {
	testCases := []struct {
		desc     string
		src      string
		expected []byte
	}{
		{
			desc: "reading relop tokens",
			src:  ">=<!===",
			expected: []byte{
				'>', '=', '<', '!', '=', '=', '=',
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := &myBufferedByteReader.BufferedByteReader{}
			myBufferedByteReader.InitBufferedByteReader(r, strings.NewReader(tC.src))
			i := 0
			for {
				b, err := r.ReadByte()
				if err == io.EOF {
					break
				} else if err != nil {
					t.Errorf("unexpected error reading file: %v", err)

				}
				if i >= len(tC.expected) {
					t.Errorf("unexpected byte: %c", b)
				}
				if b != tC.expected[i] {
					t.Errorf("error trying to %s: expected %c, got %c", tC.desc, tC.expected[0], b)
				}
				i++
			}
		})
	}
}

func TestFuzzyMyByteReaderReadsLikeBufioByteReader(t *testing.T) {
	for i := range 100_000 {
		_ = i
		src := createRandomString(1000)
		src2 := strings.Clone(src)
		myImpl := &myBufferedByteReader.BufferedByteReader{}
		myBufferedByteReader.InitBufferedByteReader(myImpl, strings.NewReader(src))
		buffioImpl := bufio.NewReader(strings.NewReader(src2))

		for {
			b1, err1 := myImpl.ReadByte()
			b2, err2 := buffioImpl.ReadByte()
			if b1 != b2 {
				t.Errorf("expected byte %c, got %c", b2, b1)
			}
			if err1 != err2 {
				t.Errorf("expected error %v, got %v", err2, err1)
			}
			if err1 == io.EOF {
				break
			}
		}
	}
}

func createRandomString(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteByte(byte(rand.Intn(256)))
	}
	return b.String()
}
