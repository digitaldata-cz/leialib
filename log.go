package leialib

import (
	"os"
	"strings"
)

func PrintStderr(s ...string) {
	os.Stderr.WriteString(strings.Join(s, " "))
}

func PrintlnStderr(s ...string) {
	str := strings.Join(s, " ")
	if str[len(str)-1] != '\n' {
		str += "\n"
	}
	PrintStderr(str)
}
