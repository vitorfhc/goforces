package fs

import (
	"log"
	"os"
	"path/filepath"
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

// FindCorrespondentOutput finds the correspondent output file to a given input
func FindCorrespondentOutput(inputFile string) (bool, string) {
	prefixLen := len(configs.InputPrefix)
	outputFile := configs.OutputPrefix + inputFile[prefixLen:]
	matches := FindFiles(outputFile)

	if len(matches) == 0 {
		return false, ""
	}

	return true, outputFile
}

// FindFiles gets a pattern and find all files using them in the current working directory
func FindFiles(pattern string) []string {
	matches, err := filepath.Glob(pattern)

	if err != nil {
		log.Fatal(err)
	}

	return matches
}
