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
