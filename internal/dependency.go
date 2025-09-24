package internal

import (
	"fmt"
	"log"
	"os/exec"
)

func CheckCommandExists(name string) error {
	_, err := exec.LookPath(name)
	if err != nil {
		log.Fatal(fmt.Errorf("%s not found in PATH. Please install it", name))
	}
	return nil
}
