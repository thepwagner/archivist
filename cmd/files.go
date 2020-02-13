package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List files",
	RunE: runIndex(func(idx *archivist.Index, _ []string) error {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(map[string]interface{}{
		})
	}),
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
