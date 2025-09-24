package internal

import (
	"log"
	"os"
)

func CreateTempFile() string {
	d, err := os.MkdirTemp("", "tempdirmdpdf")
	if err != nil {
		log.Fatal("failed to create temporary .typ file")
	}
	return d
}

func CheckFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
