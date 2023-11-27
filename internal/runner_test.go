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
	"fmt"
	"os"
	"testing"

	"github.com/AleksanderWWW/go-datareader-cli/config"
	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func NewMockReader() (*reader.MockReader, error) {
	return &reader.MockReader{}, nil
}

func getReader(cmd *cobra.Command, parsedArgs utils.ParsedRootArgs, configParser config.Parser) (reader.DataReader, error) {
	return NewMockReader()
}

func TestRunnerFailure(t *testing.T) {
	errorMsg := "Error in GetReader"
	cmd := &cobra.Command{}
	cmd.Flags().String("out", "stdout", "")

	runner := Runner{
		Cmd: cmd,
		GetReader: func(cmd *cobra.Command, parsedArgs utils.ParsedRootArgs, configParser config.Parser) (reader.DataReader, error) {
			return nil, fmt.Errorf(errorMsg)
		},
	}

	err := runner.Run()

	assert.EqualError(t, err, errorMsg)
}

func TestRunWrongOut(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("out", "unsupported-out", "")
	runner := Runner{
		Cmd:       cmd,
		GetReader: getReader,
	}

	err := runner.Run()
	assert.EqualError(t, err, "Unsupported writer option: 'unsupported-out'")
}

func TestRunWrongDateFormat(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("out", "stdout", "")
	cmd.Flags().String("start-date", "0-0-0", "")
	runner := Runner{
		Cmd:       cmd,
		GetReader: getReader,
	}
	err := runner.Run()
	assert.EqualError(t, err, "parsing time \"0-0-0\" as \"2006-01-02\": cannot parse \"0\" as \"2006\"")
}

func TestRunWrongConfig(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("out", "stdout", "")
	cmd.Flags().String("config", "invalid-config", "")

	runner := Runner{
		Cmd:       cmd,
		GetReader: getReader,
	}
	err := runner.Run()
	assert.EqualError(t, err, "Cannot find an appropriate config for path 'invalid-config'")
}

func TestRunnerSuccess(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("out", "stdout", "")

	runner := Runner{
		Cmd: cmd,
		GetReader: func(cmd *cobra.Command, parsedArgs utils.ParsedRootArgs, configParser config.Parser) (reader.DataReader, error) {
			return NewMockReader()
		},
	}

	err := runner.Run()

	assert.Equal(t, nil, err)
}

func TestRunnerWithConfig(t *testing.T) {
	type test struct {
		readerFunc GetReaderFuncType
		configPath string
	}

	testCases := []test{
		{
			readerFunc: GetFredReader,
			configPath: "../example_configs/fred_config.toml",
		},
		{
			readerFunc: GetStooqReader,
			configPath: "../example_configs/stooq_config.toml",
		},
		{
			readerFunc: GetTiingoReader,
			configPath: "../example_configs/tiingo_config.toml",
		},
		{
			readerFunc: GetBankOfCanadaReader,
			configPath: "../example_configs/bank_of_canada_config.toml",
		},
	}

	var cmd *cobra.Command
	var err error

	for _, testCase := range testCases {

		// Success
		cmd = &cobra.Command{}
		cmd.Flags().String("out", "test.csv", "")
		cmd.Flags().String("config", testCase.configPath, "")

		runner := Runner{
			Cmd:       cmd,
			GetReader: testCase.readerFunc,
		}

		err = runner.Run()
		assert.Equal(t, nil, err)

		// Failure
		cmd = &cobra.Command{}
		cmd.Flags().String("out", "test.csv", "")
		cmd.Flags().String("config", "invalid_dir/config.toml", "")
		runner = Runner{
			Cmd:       cmd,
			GetReader: testCase.readerFunc,
		}

		err = runner.Run()
		assert.EqualError(t, err, "open invalid_dir/config.toml: no such file or directory")
	}

	// Clean up: Remove the created CSV file
	err = os.Remove("test.csv")
	assert.NoError(t, err)
}
