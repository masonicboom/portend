package portend

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func ExamplePipe() {
	io.Copy(os.Stdout, New(os.Stdin))
}

func TestTransformation(t *testing.T) {
	type testCase struct {
		Input    []byte
		Expected []byte
	}
	cases := []testCase{
		{Input: []byte{cr}, Expected: []byte{lf}},
		{Input: []byte{cr, lf}, Expected: []byte{lf}},
		{Input: []byte{lf}, Expected: []byte{lf}},
		{Input: []byte{0, cr, lf}, Expected: []byte{0, lf}},
		{Input: []byte{cr, 0, lf}, Expected: []byte{lf, 0, lf}},
		{Input: []byte{cr, lf, 0}, Expected: []byte{lf, 0}},
	}

	for _, c := range cases {
		pb := New(bytes.NewBuffer(c.Input))
		res := new(bytes.Buffer)
		_, err := io.Copy(res, pb)
		if err != nil {
			t.Errorf("copying through portend Reader: %v", err)
		}
		actual := res.Bytes()
		if !bytes.Equal(actual, c.Expected) {
			t.Errorf("%v -> %v; expected %v", c.Input, actual, c.Expected)
		}
	}
}
