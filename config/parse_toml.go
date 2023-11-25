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
	"time"

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

	startDate, err := time.Parse("2006-01-02", conf.Config.StartDate)
	if err != nil {
		return reader.StooqReaderConfig{}, err
	}
	endDate, err := time.Parse("2006-01-02", conf.Config.EndDate)
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
	return reader.TiingoReaderConfig{}, nil
}
