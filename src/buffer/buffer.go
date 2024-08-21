package buffer

import (
	"errors"
	"io"
)

const (
	// DefaultBufferSize is the default size of the buffer.
	DefaultBufferSize = 8
	numberOfBuffers   = 2
)

var (
	EOFError = errors.New("EOF")
)

type DualBufferReader struct {
	buffers       [numberOfBuffers][]byte
	currBuffer    int
	currBufferPos int
	reader        io.Reader
	globalIni     int // Inicio do lexema atual
	globalProx    int // Próximo caractere a ser lido
}

func NewDualBufferReader(reader io.Reader) *DualBufferReader {

	buff := &DualBufferReader{
		buffers:    [numberOfBuffers][]byte{make([]byte, DefaultBufferSize), make([]byte, DefaultBufferSize)},
		reader:     reader,
		currBuffer: 1, // Vai voltar a ser 0 na primeira chamada de loadBuffer
		globalIni:  0,
		globalProx: 0,
	}
	buff.loadBuffer()
	return buff
}

func (dbuff *DualBufferReader) loadBuffer() error {
	dbuff.currBuffer = (dbuff.currBuffer + 1) % numberOfBuffers
	n, err := dbuff.reader.Read(dbuff.buffers[dbuff.currBuffer])
	if err != nil {
		return err
	}

	dbuff.currBufferPos = 0
	dbuff.buffers[dbuff.currBuffer] = dbuff.buffers[dbuff.currBuffer][:n]
	return nil
}

func (dbuff *DualBufferReader) LookAhead() (byte, error) {
	if dbuff.currBufferPos >= len(dbuff.buffers[dbuff.currBuffer]) {
		if err := dbuff.loadBuffer(); err != nil {
			return 0, err
		}
	}
	return dbuff.buffers[dbuff.currBuffer][dbuff.globalProx], nil
}

func (dbuff *DualBufferReader) Read() (byte, error) {
	if dbuff.globalProx >= len(dbuff.buffers[dbuff.currBuffer]) {
		if err := dbuff.loadBuffer(); err != nil {
			return 0, err
		}
	}
	c := dbuff.buffers[dbuff.currBuffer][dbuff.currBufferPos]
	dbuff.globalProx++
	dbuff.currBufferPos++
	return c, nil
}

func (dbuff *DualBufferReader) FoundToken() {
	dbuff.globalIni = dbuff.globalProx
}
