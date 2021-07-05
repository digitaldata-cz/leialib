package leialib

import (
	"os"
	"strings"
)

func GetEnvString(env, def string) (s string) {
	s = strings.Trim(os.Getenv(env), " \n\r")
	if s == "" {
		s = def
	}
	return s
}
