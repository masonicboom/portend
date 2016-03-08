// Package portend provides a simple io.Reader that normalizes line endings.
package portend

import (
	"bytes"
	"io"
	"log"
)

const (
	cr byte = 13
	lf byte = 10
)

// New creates an io.Reader that normalizes line endings.
// Strategy:
// 1) read the whole input into a bytes.Buffer,
// 2) CRLF -> LF,
// 3) CR -> LF.
func New(r io.Reader) io.Reader {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		log.Printf("wrapping reader with portend: %v", err)
		return buf
	}

	bs := buf.Bytes()
	bs = bytes.Replace(bs, []byte{cr, lf}, []byte{lf}, -1)
	bs = bytes.Replace(bs, []byte{cr}, []byte{lf}, -1)
	buf = bytes.NewBuffer(bs)

	return buf
}
