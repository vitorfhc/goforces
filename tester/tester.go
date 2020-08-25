package tester

type testRound struct {
	input          string // the input to be used on the test round
	expectedOutput string // the output expected when using the given input
	realOutput     string // the output given after running the test
	correct        bool   // the veredict, if realOutput == expectedOutput
}

// Test execute the tests for all input files with a given executable
func Test(executable string) {
	// 1. find all input and output files
	// 2. create an array of testRound structs
	// 3. execute the program for each testRound and fill it
	// 4. display results
}
