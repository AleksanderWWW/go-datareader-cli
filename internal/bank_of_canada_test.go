package internal

import (
	"testing"

	"github.com/AleksanderWWW/go-datareader-cli/utils"
	"github.com/spf13/cobra"
)

func TestBankOfCanada(t *testing.T) {
	cmd := &cobra.Command{}
	_, err := GetBankOfCanadaReader(cmd, utils.ParsedRootArgs{}, nil)

	if err != nil {
		t.Error(err)
	}
}
