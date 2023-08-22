package cmdline

import (
	"fmt"
	"io"
	"os"
)

type ioValue struct {
	name string
}

// Help prints file or -
func (v ioValue) Help(h string) string {
	return fmt.Sprintf("%s (file or \"-\")", h)
}

// String satisfies part of the flag.Value interface.
func (v ioValue) String() string { return v.name }

type ioInputValue struct {
	ioValue
	file io.Reader
}

// Set satisfies part of the flag.Value interface.
func (v *ioInputValue) Set(s string) error {
	if s != "-" {
		// nolint:gosec
		f, err := os.Open(s)
		if err != nil {
			return err
		}
		v.file = f
		v.name = s
	}
	return nil
}

type ioOutputValue struct {
	ioValue
	file io.Writer
}

// Set satisfies part of the flag.Value interface.
func (v *ioOutputValue) Set(s string) error {
	if s != "-" {
		// nolint:gosec
		f, err := os.Create(s)
		if err != nil {
			return err
		}
		v.file = f
		v.name = s
	}
	return nil
}

func (v *ioOutputValue) WriteString(s string) (int, error) {
	return io.WriteString(v.file, s)
}

func newIoInputVar() *ioInputValue {
	return &ioInputValue{
		ioValue: ioValue{
			name: "-",
		},
		file: os.Stdin,
	}
}

func newIoOutputVar() *ioOutputValue {
	return &ioOutputValue{
		ioValue: ioValue{
			name: "-",
		},
		file: os.Stdout,
	}
}
