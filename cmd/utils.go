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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/spf13/cobra"
)

func isCSV(name string) bool {
	splited := strings.Split(name, ".")
	return strings.ToLower(splited[len(splited)-1]) == "csv"
}

func getWriterFunc(out string) (func(dataframe.DataFrame), error) {
	if out == "stdout" {
		return func(data dataframe.DataFrame) {
			fmt.Println(data)
		}, nil

	} else if isCSV(out) {
		return func(data dataframe.DataFrame) {
			f, err := os.Create(out)
			if err != nil {
				f.Close()
				log.Fatal(err)
			}
			data.WriteCSV(f)
		}, nil
	} else {
		return nil, fmt.Errorf("Unsupported writer option: '%s'", out)
	}
}

type parsedDefaultArgs struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	Out       string
}

func parseDefaultArgs(cmd *cobra.Command) (parsedDefaultArgs, error) {
	symbols, _ := cmd.Flags().GetStringSlice("symbols")

	startDateStr, _ := cmd.Flags().GetString("start-date")
	startDate, err := time.Parse("2006-01-02", startDateStr)

	endDateStr, _ := cmd.Flags().GetString("end-date")
	endDate, err := time.Parse("2006-01-02", endDateStr)

	out, _ := cmd.Flags().GetString("out")

	return parsedDefaultArgs{
		Symbols:   symbols,
		StartDate: startDate,
		EndDate:   endDate,
		Out:       out,
	}, err
}
