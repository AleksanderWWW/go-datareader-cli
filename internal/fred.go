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
	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/spf13/cobra"
)

func GetFredReader(cmd *cobra.Command, parsedArgs utils.ParsedRootArgs, configParser config.Parser) (reader.DataReader, error) {
	var config reader.FredReaderConfig
	var err error

	if parsedArgs.Config == "" {
		config = reader.FredReaderConfig{
			Symbols:   parsedArgs.Symbols,
			StartDate: parsedArgs.StartDate,
			EndDate:   parsedArgs.EndDate,
		}
	} else {
		config, err = configParser.ParseFredConfig()
		if err != nil {
			return &reader.BOCDataReader{}, err
		}
	}
	return reader.NewFredDataReader(config)
}
