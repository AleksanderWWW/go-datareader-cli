package internal

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestStooq(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("freq", "d", "")
	_, err := GetStooqReader(cmd, parsedRootArgs{})

	if err != nil {
		t.Error(err)
	}

	_, err = GetStooqReader(&cobra.Command{}, parsedRootArgs{})
	assert.Error(t, err)
}
