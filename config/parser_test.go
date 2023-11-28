package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhichParser(t *testing.T) {
	type test struct {
		path             string
		expectedParser   string
		expectedErrorMsg string
	}

	testCases := []test{
		{
			path:             "",
			expectedParser:   "nil",
			expectedErrorMsg: "",
		},
		{
			path:             "config.toml",
			expectedParser:   "*config.TomlConfigParser",
			expectedErrorMsg: "",
		},
		{
			path:             "config.invalid_ext",
			expectedParser:   "nil",
			expectedErrorMsg: "Cannot find an appropriate config for path 'config.invalid_ext'",
		},
	}

	for _, testCase := range testCases {
		parser, err := WhichParser(testCase.path)

		if testCase.expectedErrorMsg == "" {
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, testCase.expectedErrorMsg)
		}

		if testCase.expectedParser == "nil" {
			assert.Nil(t, parser)
		} else {
			assert.Equal(t, testCase.expectedParser, reflect.TypeOf(parser).String())
		}
	}
}
