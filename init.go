package leialib

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func Init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	os.Chdir(dir)
	rand.Seed(time.Now().UTC().UnixNano())
}
