package fs

import (
	"log"
	"os"
	"strconv"

	"github.com/vitorfhc/goforces/configs"
)

// BuildProblemFilename build the filename of the input/output txt file
func BuildProblemFilename(ind int, isInput bool) string {
	var filename string

	if isInput {
		filename = configs.InputPrefix + strconv.Itoa(ind) + configs.FilePostfix
	} else {
		filename = configs.OutputPrefix + strconv.Itoa(ind) + configs.FilePostfix
	}

	return filename
}

// WriteFile writes the text to the given file (path)
func WriteFile(path string, text string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(text)
}
