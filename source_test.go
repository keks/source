package source

import (
	"io"
	"os"
	"testing"
)

type test struct {
	path    string
	content []byte
}

var (
	table = []test{
		test{"testresource", []byte("test\n")},
	}
)

func TestSource(t *testing.T) {
	for _, tt := range table {
		path := Path(tt.path)

		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		buf := make([]byte, len(tt.content))

		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}
		if n != len(tt.content) {
			t.Fatal("unexpected read len:", n, "!=", len(tt.content))
		}

		for i := range buf {
			if buf[i] != tt.content[i] {
				t.Fatal("content not equal: ", tt.path)
			}
		}

	}
}
