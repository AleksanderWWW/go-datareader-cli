/*
Copyright © 2023 Aleksander Wojnarowicz <alwojnarowicz@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/spf13/cobra"
)

// tiingoCmd represents the tiingo command
var tiingoCmd = &cobra.Command{
	Use:   "tiingo",
	Short: "Get financial data from Tiingo",
	Run: func(cmd *cobra.Command, args []string) {
		parsedArgs, err := parseDefaultArgs(cmd)

		if err != nil {
			log.Fatal(err)
		}

		apiKey, _ := cmd.Flags().GetString("api-key")

		config := reader.TiingoReaderConfig{
			StartDate: parsedArgs.StartDate,
			EndDate:   parsedArgs.EndDate,
			ApiKey:    apiKey,
		}

		writerFunc, err := getWriterFunc(parsedArgs.Out)
		if err != nil {
			log.Fatal(err)
		}

		tiingoReader, err := reader.NewTiingoDailyReader(parsedArgs.Symbols, config)
		if err != nil {
			log.Fatal(err)
		}

		data := reader.GetData(tiingoReader)
		writerFunc(data)
	},
}

func init() {
	rootCmd.AddCommand(tiingoCmd)
	tiingoCmd.LocalFlags().String("api-key", "", "sample usage")
}
