package null

import (
	"bytes"
	"io"
	"os"
	"testing"

	"src.userspace.com.au/logger/message"
)

func TestWriter(t *testing.T) {
	orig := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	l := New()

	l.Write(message.Message{Content: "this should not print"})

	out := make(chan string)
	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, r)
		out <- buf.String()
	}()

	w.Close()
	os.Stdout = orig
	actual := <-out

	if len(actual) > 0 {
		t.Errorf("expected 0 bytes")
	}
}
