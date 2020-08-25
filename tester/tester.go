package tester

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/vitorfhc/goforces/configs"
	"github.com/vitorfhc/goforces/fs"
)

type testRound struct {
	input          string // the input to be used on the test round
	expectedOutput string // the file with expected output when using the given input
	realOutput     string // the output given after running the test
	correct        bool   // the veredict, if realOutput == expectedOutput
}

// Test execute the tests for all input files with a given executable
func Test(executable string) {
	rounds := constructRounds()

	if len(rounds) == 0 {
		log.Fatal("No inputs with corresponding outputs found")
	}

	for _, round := range rounds {
		runTest(executable, &round)
	}
}

func runTest(executable string, round *testRound) {
	input, err := os.Open(round.input)

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(executable)
	cmd.Stdin = input

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	round.realOutput = string(output)

	expected := fs.ReadFile(round.expectedOutput)
	if expected != round.realOutput {
		fmt.Printf("Test failed - %s and %s\n", round.input, round.expectedOutput)
		round.correct = false
		return
	}

	round.correct = true
	fmt.Printf("Test success - %s and %s\n", round.input, round.expectedOutput)
}

func constructRounds() []testRound {
	var rounds []testRound
	matches := fs.FindFiles(configs.InputPrefix + "*" + configs.FilePostfix)

	for _, match := range matches {
		var round testRound

		hasOutput, outputFile := fs.FindCorrespondentOutput(match)

		if !hasOutput {
			continue
		}

		round.input = match
		round.correct = false
		round.expectedOutput = outputFile
		round.realOutput = ""

		rounds = append(rounds, round)
	}

	return rounds
}
