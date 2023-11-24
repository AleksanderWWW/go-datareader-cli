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
	"os"

	"github.com/AleksanderWWW/go-datareader-cli/internal"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, getReaderFunc internal.GetReaderFuncType) {
	cmdRunner := internal.Runner{
		Cmd:       cmd,
		GetReader: getReaderFunc,
	}
	err := cmdRunner.Run()
	if err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "go-datareader",
	Short: "Download tabular financial data with command line",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-datareader-cli.yaml)")
	rootCmd.PersistentFlags().StringSlice("symbols", []string{}, "List of symbols to scrape data for in the form --symbols=s1,s2,...sn")
	rootCmd.PersistentFlags().String("start-date", "", "Start date in the format YYYY-mm-dd e.g. --start-date=2023-07-31. Default depends on the provider used.")
	rootCmd.PersistentFlags().String("end-date", "", "End date in the format YYYY-mm-dd e.g. --end-date=2023-07-31. Default depends on the provider used.")
	rootCmd.PersistentFlags().String("out", "stdout", "Where to write the downloaded data. Leaving the default just prints to console")
}
