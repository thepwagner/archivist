package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var drivesCmd = &cobra.Command{
	Use:   "drives",
	Short: "List drives",
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(map[string]interface{}{
			"drives": idx.Drives,
		})
	}),
}

func init() {
	rootCmd.AddCommand(drivesCmd)
}
