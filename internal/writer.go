package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReplaceInFile(inputFileName string, outputFile *os.File) error {
	in, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("could not open temp file")
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(outputFile)

	// add md checklist
	// writeLine(writer, `#import "@preview/cheq:0.3.0": checklist`)
	// writeLine(writer, "#show: checklist")
	// writeLine(writer, "")

	// add md alerts
	// writeLine(writer, `#import "@preview/note-me:0.5.0": *`)
	// writeLine(writer, "")

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "- ☐ ", "- [ ] ")
		line = strings.ReplaceAll(line, "- ☒ ", "- [x] ")
		// line = replaceAlert(line, `\[!NOTE\]`, "note")
		// line = replaceAlert(line, `\[!TIP\]`, "tip")
		// line = replaceAlert(line, `\[!IMPORTANT\]`, "important")
		// line = replaceAlert(line, `\[!WARNING\]`, "warning")
		// line = replaceAlert(line, `\[!CAUTION\]`, "caution")
		writeLine(writer, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error while reading input file: %w", err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal("failed to flush output file: %w", err)
	}

	return nil
}

func writeLine(writer *bufio.Writer, line string) {
	_, err := writer.WriteString(line + "\n")
	if err != nil {
		log.Fatal("failed to write to output file: %w", err)
	}
}

func replaceAlert(line string, markdown string, typst string) string {
	if strings.Contains(line, markdown) {
		noalert := strings.ReplaceAll(line, markdown, "")
		return `#` + typst + `[` + noalert + `]`
	}
	return line
}
