/*
Copyright © 2023 Aleksander Wojnarowicz <alwojnarowicz@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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
		out := IsCSV(testItem.in)
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
		_, err := GetWriterFunc(goodOut)
		if err != nil {
			t.Errorf("Unexpected error for %s - %s", goodOut, err)
		}
	}

	for _, badOut := range badOuts {
		_, err := GetWriterFunc(badOut)
		if err == nil {
			t.Errorf("Expected error for %s, but did not get it", badOut)
		}
	}
}

func TestWriterFunctions(t *testing.T) {
	// Create sample DataFrame for testing
	data := dataframe.DataFrame{}

	// Test writing to stdout
	stdoutWriter, _ := GetWriterFunc("stdout")
	err := stdoutWriter(data)
	assert.NoError(t, err)

	// Test writing to CSV file
	csvFileName := "test.csv"
	csvWriter, _ := GetWriterFunc(csvFileName)
	err = csvWriter(data)
	assert.NoError(t, err)

	// Clean up: Remove the created CSV file
	err = os.Remove(csvFileName)
	assert.NoError(t, err)
}

func TestWriterFunctionsInvalidCsvFile(t *testing.T) {
	// Create sample DataFrame for testing
	data := dataframe.DataFrame{}

	// Test writing to stdout
	stdoutWriter, _ := GetWriterFunc("stdout")
	err := stdoutWriter(data)
	assert.NoError(t, err)

	// Test that a non-existent  file path will cause an error
	csvFileName := "./non_existent_dir/file.csv"
	csvWriter, _ := GetWriterFunc(csvFileName)
	err = csvWriter(data)
	assert.Error(t, err, "open ./non_existent_dir/file.csv: no such file or directory")
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
		out, err := ParseDate(test.in)

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

func TestParseArgs(t *testing.T) {
	expectedParsedArgs := ParsedRootArgs{
		Symbols:   []string{"s1", "s2", "s3"},
		StartDate: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
		Out:       "data.csv",
		Config:    "config.toml",
	}

	actualParsedArgs, err := ParseArgs(
		[]string{"s1", "s2", "s3"},
		"2022-12-31",
		"2023-12-31",
		"data.csv",
		"config.toml",
	)

	if err != nil {
		t.Errorf("Unexpected error while parsing args: %s", err)
	}

	if !reflect.DeepEqual(expectedParsedArgs, actualParsedArgs) {
		t.Errorf(
			"Parsed args do not match the exepcted structure:\n Expected: %v\n Actual: %v",
			expectedParsedArgs,
			actualParsedArgs,
		)
	}
}

func TestParseArgsFailure(t *testing.T) {
	_, err := ParseArgs(
		[]string{"s1", "s2", "s3"},
		"0-0-0",
		"2023-12-31",
		"data.csv",
		"config.toml",
	)
	assert.EqualError(t, err, "parsing time \"0-0-0\" as \"2006-01-02\": cannot parse \"0\" as \"2006\"")

	_, err = ParseArgs(
		[]string{"s1", "s2", "s3"},
		"2023-12-31",
		"0-0-0",
		"data.csv",
		"config.toml",
	)
	assert.EqualError(t, err, "parsing time \"0-0-0\" as \"2006-01-02\": cannot parse \"0\" as \"2006\"")
}

func TestParseRootArgs(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().StringSlice("symbols", []string{"s1", "s2", "s3"}, "")
	cmd.Flags().String("start-date", "2023-11-11", "")
	cmd.Flags().String("end-date", "2023-11-25", "")
	cmd.Flags().String("out", "test.csv", "")
	cmd.Flags().String("config", "config.toml", "")

	rootArgs, err := ParseRootArgs(cmd)
	if err != nil {
		t.Error(err)
	}

	startDate, _ := ParseDate("2023-11-11")
	endDate, _ := ParseDate("2023-11-25")
	assert.Equal(
		t,
		ParsedRootArgs{
			Symbols:   []string{"s1", "s2", "s3"},
			StartDate: startDate,
			EndDate:   endDate,
			Out:       "test.csv",
			Config:    "config.toml",
		},
		rootArgs,
	)
}
