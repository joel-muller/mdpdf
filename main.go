package main

import (
	"fmt"
	"log"
	"mdpdf/internal"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	input, output := internal.Parse(args)
	internal.CheckCommandExists("pandoc")
	internal.CheckCommandExists("typst")

	if !internal.CheckFileExists(input) {
		log.Fatalf("%v does not exists", input)
	}

	tempfolder := internal.CreateTempDir()
	defer os.RemoveAll(tempfolder)
	tempfile := filepath.Join(tempfolder, "temp.typ")

	pandocCmd := exec.Command("pandoc", input, "-o", tempfile)
	pandocCmd.Stdout = os.Stdout
	pandocCmd.Stderr = os.Stderr
	if err := pandocCmd.Run(); err != nil {
		log.Fatalf("Pandoc failed: %v", err)
	}

	newtempfile := internal.CreateTempFile(tempfolder)
	internal.ReplaceInFile(tempfile, newtempfile)

	typstCmd := exec.Command("typst", "compile", newtempfile.Name(), output)
	typstCmd.Stdout = os.Stdout
	typstCmd.Stderr = os.Stderr
	if err := typstCmd.Run(); err != nil {
		log.Fatalf("Typst failed: %v", err)
	}

	fmt.Printf("Successfully generated PDF: %s\n", output)
}
