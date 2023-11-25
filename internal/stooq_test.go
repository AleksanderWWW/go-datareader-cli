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
	"testing"

	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestStooq(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("freq", "d", "")
	_, err := GetStooqReader(cmd, utils.ParsedRootArgs{}, nil)

	if err != nil {
		t.Error(err)
	}

	_, err = GetStooqReader(&cobra.Command{}, utils.ParsedRootArgs{}, nil)
	assert.Error(t, err)
}
