package myBufferedByteReader

import (
	"io"
)

const (
	DefaultBufferSize uint = 4096
	numberOfBuffers   uint = 2
)

type BufferedByteReader struct {
	file io.Reader
	bufs [numberOfBuffers][DefaultBufferSize]byte

	currBuffIndex, posInCurrBuf, currBuffLen uint
}

func InitBufferedByteReader(bbr *BufferedByteReader, f io.Reader) {
	bbr.file = f
	bbr.currBuffIndex = numberOfBuffers - 1 // Após a primeira chamada de loadBuff, currBuffIndex será 0
	bbr.posInCurrBuf = DefaultBufferSize    // Forçar a chamada de loadBuff na primeira chamada de ReadByte
	bbr.currBuffLen = DefaultBufferSize     // Se não for igual em algum momento, significa que faltam currBuffLen - posInCurrBuf bytes para serem lidos
}

func (r *BufferedByteReader) ReadByte() (byte, error) {
	if r.posInCurrBuf == DefaultBufferSize {
		err := r.loadBuff()
		if err != nil {
			return 0, err
		}
	}
	if r.posInCurrBuf >= r.currBuffLen {
		return 0, io.EOF
	}
	r.posInCurrBuf++
	return r.bufs[r.currBuffIndex][r.posInCurrBuf-1], nil
}

// Read the next DefaultBufferSize bytes from the file to the buffer
func (r *BufferedByteReader) loadBuff() error {
	r.currBuffIndex = (r.currBuffIndex + 1) % numberOfBuffers

	n, err := r.file.Read(r.bufs[r.currBuffIndex][:])
	r.currBuffLen = uint(n)
	if err != nil {
		return err
	}
	r.posInCurrBuf = 0
	return nil
}
