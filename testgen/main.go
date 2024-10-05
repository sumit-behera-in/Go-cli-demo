package main

import (
	"io/ioutil"
	"log"

	"github.com/cweill/gotests"
)

func main() {
	// Specify the Go file or directory where you want to generate tests
	filename := "C:/Users/Sumit/Desktop/Go-cli-demo/testgen/main.go"

	// Create options for generating the tests
	options := &gotests.Options{
		Exported:    false, // Include both exported and unexported functions
		PrintInputs: true,  // Print function parameters in error messages
		Subtests:    false, // Generate normal tests, not subtests
	}

	// Generate test cases for all available functions in the file
	tests, err := gotests.GenerateTests(filename, options)
	if err != nil {
		log.Fatalf("Error generating tests: %v", err)
	}

	// Write the generated test cases to the corresponding _test.go file
	for _, test := range tests {
		// Use the Path field to get the test file path (e.g., yourfile_test.go)
		testFilename := test.Path

		// Write the generated test case to the file
		err = ioutil.WriteFile(testFilename, test.Output, 0644)
		if err != nil {
			log.Fatalf("Error writing test file: %v", err)
		}
		log.Printf("Test case written to: %s\n", testFilename)
	}
}

// filename := "C:/Users/Sumit/Desktop/Go-cli-demo/checker/main.go"
