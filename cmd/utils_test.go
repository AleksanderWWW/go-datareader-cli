package cmd

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

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

func TestGetWriterFunc(t *testing.T) {

	goodOuts := []string{
		"stdout",
		"file.csv",
	}

	badOuts := []string{
		"abc",
		"csv",
		"json",
		"xlsx",
	}

	for _, goodOut := range goodOuts {
		_, err := getWriterFunc(goodOut)
		if err != nil {
			t.Errorf("Unexpected error for %s - %s", goodOut, err)
		}
	}

	for _, badOut := range badOuts {
		_, err := getWriterFunc(badOut)
		if err == nil {
			t.Errorf("Expected error for %s, but did not get it", badOut)
		}
	}
}

func TestParseDate(t *testing.T) {
	type test struct {
		in          string
		expectedOut time.Time
	}

	tests := []test{
		{
			in:          "",
			expectedOut: time.Time{},
		},
		{
			in:          "2023-12-31",
			expectedOut: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		out, err := parseDate(test.in)

		if err != nil {
			t.Errorf("Unexpected error while parsing dates - %s", err)
		}

		if out != test.expectedOut {
			t.Errorf(
				"Wrong output parseDate:\n Expected: %s\n Actual: %s",
				test.expectedOut,
				out,
			)
		}
	}
}
