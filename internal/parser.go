package internal

import (
	"log"
	"strings"
)

func Parse(args []string) (string, string) {
	if len(args) != 1 && len(args) != 2 {
		log.Fatal("argument number is wrong, must be 1 or 2")
	}

	input := args[0]

	if !checkFiletype(input, ".md") {
		log.Fatal("input filetype is not .md")
	}

	var output string
	if len(args) == 2 {
		output = args[1]
		if !checkFiletype(output, ".pdf") {
			log.Fatal("output filetype is not .pdf")
		}
	} else {
		output = strings.TrimSuffix(input, ".md") + ".pdf"
	}

	return input, output
}

func checkFiletype(name string, filetype string) bool {
	if strings.HasSuffix(name, filetype) {
		return true
	}
	return false
}
