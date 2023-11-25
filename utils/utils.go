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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/spf13/cobra"
)

func IsCSV(name string) bool {
	splited := strings.Split(name, ".")
	return len(splited) >= 2 && strings.ToLower(splited[len(splited)-1]) == "csv"
}

func GetWriterFunc(out string) (func(dataframe.DataFrame) error, error) {
	if out == "stdout" {
		return func(data dataframe.DataFrame) error {
			fmt.Println(data)
			return nil
		}, nil

	} else if IsCSV(out) {
		return func(data dataframe.DataFrame) error {
			f, err := os.Create(out)
			if err != nil {
				f.Close()
				log.Fatal(err)
			}
			return data.WriteCSV(f)
		}, nil

	} else {
		return nil, fmt.Errorf("Unsupported writer option: '%s'", out)
	}
}

type ParsedRootArgs struct {
	Symbols   []string
	StartDate time.Time
	EndDate   time.Time
	Out       string
	Config    string
}

func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", dateStr)
}

func ParseArgs(
	symbols []string,
	startDateStr string,
	endDateStr string,
	out string,
	config string,
) (ParsedRootArgs, error) {
	startDate, err := ParseDate(startDateStr)
	endDate, err := ParseDate(endDateStr)

	return ParsedRootArgs{
		Symbols:   symbols,
		StartDate: startDate,
		EndDate:   endDate,
		Out:       out,
		Config:    config,
	}, err
}

func ParseRootArgs(cmd *cobra.Command) (ParsedRootArgs, error) {
	symbols, _ := cmd.Flags().GetStringSlice("symbols")

	startDateStr, _ := cmd.Flags().GetString("start-date")

	endDateStr, _ := cmd.Flags().GetString("end-date")

	out, _ := cmd.Flags().GetString("out")

	config, _ := cmd.Flags().GetString("config")

	return ParseArgs(symbols, startDateStr, endDateStr, out, config)
}
