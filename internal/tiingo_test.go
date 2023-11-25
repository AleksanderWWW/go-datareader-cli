package internal

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestTiingo(t *testing.T) {

	_, err := GetTiingoReader(&cobra.Command{}, parsedRootArgs{}, nil)

	if err != nil {
		t.Error(err)
	}
}
