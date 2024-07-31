package common

import "strings"

func EscapeJS(val string) string {
	builder := strings.Builder{}
	for _, letter := range val {
		if letter == '\'' {
			builder.WriteString("\\")
		}
		builder.WriteByte(byte(letter))
	}

	return builder.String()
}
