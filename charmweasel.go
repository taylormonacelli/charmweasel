package charmweasel

import (
	"bytes"
	"log/slog"
)

func Main() int {
	return 0
}

type CustomType int

const (
	Foo CustomType = iota
	Bar
	Baz
)

func VariadicFunction(level slog.Level, args ...CustomType) *bytes.Buffer {
	var buf bytes.Buffer

	for _, arg := range args {
		switch arg {
		case Foo:
			buf.WriteString("Foo\n")
		case Bar:
			buf.WriteString("Bar\n")
		case Baz:
			buf.WriteString("Baz\n")
		default:
			buf.WriteString("Unknown\n")
		}
	}

	buf.WriteString("Log level: " + level.String() + "\n")

	return &buf
}
