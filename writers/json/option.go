package json

import "io"

// Option configures the writer
type Option func(*Writer) error

// SetTimeFormat configures the format used for timestamps.
func SetTimeFormat(f string) Option {
	return func(w *Writer) error {
		w.timeFormat = f
		return nil
	}
}

// SetOutput configures the output.
func SetOutput(o io.Writer) Option {
	return func(w *Writer) error {
		w.writer = o
		return nil
	}
}
