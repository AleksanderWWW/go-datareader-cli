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
	"github.com/AleksanderWWW/go-datareader-cli/config"
	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/spf13/cobra"
)

type GetReaderFuncType func(*cobra.Command, parsedRootArgs, config.Parser) (reader.DataReader, error)

type Runner struct {
	Cmd       *cobra.Command
	GetReader GetReaderFuncType
}

func (r *Runner) Run() error {
	parsedArgs, err := parseRootArgs(r.Cmd)

	if err != nil {
		return err
	}

	writerFunc, err := getWriterFunc(parsedArgs.Out)
	if err != nil {
		return err
	}

	parser, err := config.WhichParser(parsedArgs.Config)
	if err != nil {
		return err
	}

	dataReader, err := r.GetReader(r.Cmd, parsedArgs, parser)
	if err != nil {
		return err
	}

	data := reader.GetData(dataReader)
	return writerFunc(data)
}
