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

package internal

import (
	"reflect"
	"testing"
	"time"
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

func TestParseArgs(t *testing.T) {
	expectedParsedArgs := parsedRootArgs{
		Symbols:   []string{"s1", "s2", "s3"},
		StartDate: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
		Out:       "data.csv",
	}

	actualParsedArgs, err := parseArgs(
		[]string{"s1", "s2", "s3"},
		"2022-12-31",
		"2023-12-31",
		"data.csv",
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
