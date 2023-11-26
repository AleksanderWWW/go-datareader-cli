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

package config

import (
	"fmt"

	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/AleksanderWWW/go-datareader/reader"
	"github.com/BurntSushi/toml"
)

type TomlConfigParser struct {
	ConfigPath string
}

func NewTomlConfigParser(path string) *TomlConfigParser {
	return &TomlConfigParser{ConfigPath: path}
}

func (t *TomlConfigParser) ParseStooqConfig() (reader.StooqReaderConfig, error) {
	type config struct {
		Symbols   []string
		StartDate string
		EndDate   string
		Freq      string
	}

	type stooqConfig struct {
		Reader string
		Config config
	}

	var conf stooqConfig

	_, err := toml.DecodeFile(t.ConfigPath, &conf)
	if err != nil {
		return reader.StooqReaderConfig{}, err
	}

	if conf.Reader != "stooq" {
		return reader.StooqReaderConfig{}, fmt.Errorf("Invalid `Reader` field in the config file. Expected: `stooq`. Got: %s", conf.Reader)
	}

	startDate, err := utils.ParseDate(conf.Config.StartDate)
	if err != nil {
		return reader.StooqReaderConfig{}, err
	}
	endDate, err := utils.ParseDate(conf.Config.EndDate)
	if err != nil {
		return reader.StooqReaderConfig{}, err
	}

	return reader.StooqReaderConfig{
		Symbols:   conf.Config.Symbols,
		StartDate: startDate,
		EndDate:   endDate,
		Freq:      conf.Config.Freq,
	}, nil
}

func (t *TomlConfigParser) ParseTiingoConfig() (reader.TiingoReaderConfig, error) {
	type config struct {
		Symbols   []string
		StartDate string
		EndDate   string
		ApiKey    string
	}

	type tiingoConfig struct {
		Reader string
		Config config
	}

	var conf tiingoConfig

	_, err := toml.DecodeFile(t.ConfigPath, &conf)
	if err != nil {
		return reader.TiingoReaderConfig{}, err
	}

	if conf.Reader != "tiingo" {
		return reader.TiingoReaderConfig{}, fmt.Errorf("Invalid `Reader` field in the config file. Expected: `tiingo`. Got: %s", conf.Reader)
	}

	startDate, err := utils.ParseDate(conf.Config.StartDate)
	if err != nil {
		return reader.TiingoReaderConfig{}, err
	}
	endDate, err := utils.ParseDate(conf.Config.EndDate)
	if err != nil {
		return reader.TiingoReaderConfig{}, err
	}

	return reader.TiingoReaderConfig{
		Symbols:   conf.Config.Symbols,
		StartDate: startDate,
		EndDate:   endDate,
		ApiKey:    conf.Config.ApiKey,
	}, nil
}

func (t *TomlConfigParser) ParseBankOfCanadaConfig() (reader.BOCReaderConfig, error) {
	type config struct {
		Symbols   []string
		StartDate string
		EndDate   string
	}

	type tiingoConfig struct {
		Reader string
		Config config
	}

	var conf tiingoConfig

	_, err := toml.DecodeFile(t.ConfigPath, &conf)
	if err != nil {
		return reader.BOCReaderConfig{}, err
	}

	if conf.Reader != "bank of canada" {
		return reader.BOCReaderConfig{}, fmt.Errorf("Invalid `Reader` field in the config file. Expected: `bank of canada`. Got: %s", conf.Reader)
	}

	startDate, err := utils.ParseDate(conf.Config.StartDate)
	if err != nil {
		return reader.BOCReaderConfig{}, err
	}
	endDate, err := utils.ParseDate(conf.Config.EndDate)
	if err != nil {
		return reader.BOCReaderConfig{}, err
	}

	return reader.BOCReaderConfig{
		Symbols:   conf.Config.Symbols,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}
