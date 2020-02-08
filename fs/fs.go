package fs

import (
	"log"
	"os"
	"strconv"
)

// BuildProblemFilename build the filename of the input/output txt file
func BuildProblemFilename(ind int, isInput bool) string {
	var filename string

	if isInput {
		filename = "input-" + strconv.Itoa(ind) + ".txt"
	} else {
		filename = "output-" + strconv.Itoa(ind) + ".txt"
	}

	return filename
}

// WriteTestFile writes the text to the given file (path)
func WriteTestFile(path string, text string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(text)
}
