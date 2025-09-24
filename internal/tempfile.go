package internal

import (
	"log"
	"os"
)

func CreateTempDir() string {
	d, err := os.MkdirTemp("", "tempdirmdpdf")
	if err != nil {
		log.Fatal("failed to create temporary .typ file")
	}
	return d
}

func CreateTempFile(dir string) *os.File {
	file, err := os.CreateTemp(dir, "*.typ")
	if err != nil {
		log.Fatal("failed to create temporary file for parsing")
	}
	return file
}

func CheckFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
