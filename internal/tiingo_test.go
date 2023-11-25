package internal

import (
	"testing"

	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/spf13/cobra"
)

func TestTiingo(t *testing.T) {

	_, err := GetTiingoReader(&cobra.Command{}, utils.ParsedRootArgs{}, nil)

	if err != nil {
		t.Error(err)
	}
}
