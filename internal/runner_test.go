package internal

import (
	"fmt"
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
