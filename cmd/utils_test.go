package cmd

import "testing"

func TestIsCSV(t *testing.T) {
	type test struct {
		in          string
		expectedOut bool
	}

	tests := []test{
		{
			in:          "example.csv",
			expectedOut: true,
		},
		{
			in:          "csv", // that is not a file
			expectedOut: false,
		},
		{
			in:          ".csv",
			expectedOut: true,
		},
		{
			in:          "example.json",
			expectedOut: false,
		},
		{
			in:          "example",
			expectedOut: false,
		},
		{
			in:          "",
			expectedOut: false,
		},
	}

	for _, testItem := range tests {
		out := isCSV(testItem.in)
		if out != testItem.expectedOut {
			t.Errorf("Failed checking if CSV:\n In: %s\n Expected: %t\n Actual: %t", testItem.in, testItem.expectedOut, out)
		}
	}
}
