package app

import (
	"os"
)

func HandleFatal(err error) {
	buf := []byte("FATAL: ")
	buf = append(buf, err.Error()...)
	buf = append(buf, '\n')
	os.Stderr.Write(buf)
	os.Exit(-1)
}
