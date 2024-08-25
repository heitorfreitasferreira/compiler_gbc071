package buffer

import (
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	dbuff := NewDualBufferReader(strings.NewReader("a"))
	if dbuff == nil {
		t.Fatalf("Error creating DualBufferReader")
	}
	if dbuff.currBuffer != 0 {
		t.Error("Error initializing current buffer")
	}

	if dbuff.globalIni != 0 {
		t.Error("Error initializing buffer start")
	}
}

func TestChangingIniAndProxPos(t *testing.T) {
	dbuff := NewDualBufferReader(strings.NewReader("Hello World"))

	read := ""
	prox := dbuff.globalProx
	for {
		lkah, err := dbuff.LookAhead()

		if err != nil {
			t.Fatalf("Error reading LookAhead: %v", err)
		}

		if lkah == ' ' {
			dbuff.FoundToken()
			break
		}

		if prox != dbuff.globalProx {
			t.Fatalf("Error changing prox position: %v\nShould be %d", dbuff.globalProx, prox)
		}
		readByte, err := dbuff.Read()
		if err != nil {
			t.Fatalf("Error reading byte: %v", err)
		}
		prox++
		read += string(readByte)
	}

	if read != "Hello" {
		t.Fatalf("Error reading bytes: %v\nShould be \"Hello\"", read)
	}

	if dbuff.globalIni != len(read) {
		t.Fatalf("Error changing ini position: %v\nShould be 0", dbuff.globalIni)
	}

	if dbuff.globalProx != len(read) {
		t.Fatalf("Error changing prox position: %v\nShould be %d", dbuff.globalProx, len(read))
	}
}
func TestReadingOverTheBuffer(t *testing.T) {
	s := nAs(DefaultBufferSize, 'a')
	dbuff := NewDualBufferReader(strings.NewReader(s + "b"))

	var lastByteRead byte
	var err error

	for i := 0; i < DefaultBufferSize+1; i++ {
		lastByteRead, err = dbuff.Read()
		if err != nil {
			t.Errorf("Error reading byte: %v at pos:%d", err, i)
		}
		if i == DefaultBufferSize {
			if lastByteRead != 'b' {
				t.Errorf("Error reading byte: %v\nShould be 'b'", lastByteRead)
			}
			if dbuff.currBuffer != 1 {
				t.Errorf("Error changing buffer: %v\nShould be 1", dbuff.currBuffer)
			}
			if dbuff.currBufferPos != 1 {
				t.Errorf("Error changing buffer position: %v\nShould be 1", dbuff.currBufferPos)
			}
		} else if lastByteRead != 'a' {
			t.Errorf("Error reading byte: %v\nShould be 'a'", lastByteRead)
		}
	}
}

func TestInputNSizeOfTheBuffer(t *testing.T) {
	nTimesBufferSize := 3
	s := nAs(DefaultBufferSize, 'a') + nAs(DefaultBufferSize, 'b') + nAs(DefaultBufferSize, 'c')
	dbuff := NewDualBufferReader(strings.NewReader(s))
	for i := 1; i < DefaultBufferSize*nTimesBufferSize+1; i++ {
		_, err := dbuff.Read()
		if err != nil {
			t.Errorf("Error reading byte: %v at pos:%d", err, i)
		}
		chunk := i / DefaultBufferSize % nTimesBufferSize
		if chunk == 0 {
			if dbuff.currBuffer != 0 {
				t.Errorf("Error changing buffer: %d\nShould be 0\n dbuffState: %v", dbuff.currBuffer, dbuff)
			}
		} else if chunk == 1 {
			if dbuff.currBuffer != 1 {
				t.Errorf("Error changing buffer: %d\nShould be 1", dbuff.currBuffer)
			}
		} else if chunk == 2 {
			if dbuff.currBuffer != 0 {
				t.Errorf("Error changing buffer: %d\nShould be 0", dbuff.currBuffer)
			}
		}
	}
}
func nAs(size int, r rune) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteRune(r)
	}
	return b.String()
}
