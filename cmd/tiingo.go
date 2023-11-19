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

package cmd

import (
	"log"

	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/spf13/cobra"
)

func NewTiingoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tiingo",
		Short: "Get financial data from Tiingo",
		Run: func(cmd *cobra.Command, args []string) {
			parsedArgs, err := parseRootArgs(cmd)

			if err != nil {
				log.Fatal(err)
			}

			apiKey, _ := cmd.Flags().GetString("api-key")

			config := reader.TiingoReaderConfig{
				Symbols:   parsedArgs.Symbols,
				StartDate: parsedArgs.StartDate,
				EndDate:   parsedArgs.EndDate,
				ApiKey:    apiKey,
			}

			writerFunc, err := getWriterFunc(parsedArgs.Out)
			if err != nil {
				log.Fatal(err)
			}

			tiingoReader, err := reader.NewTiingoDailyReader(config)
			if err != nil {
				log.Fatal(err)
			}

			data := reader.GetData(tiingoReader)
			writerFunc(data)
		},
	}
}

func init() {
	tiingoCmd := NewTiingoCmd()
	rootCmd.AddCommand(tiingoCmd)
	tiingoCmd.LocalFlags().String("api-key", "", "[Optional] Pass your Tiingo API token here")
}
