package cmd

import (
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

func AddFilesCmd(id *archivist.Index, args []string) error {
	return nil
}

var filesystemAddCmd = &cobra.Command{
	Use:   "fsadd [path]",
	Short: "Add filesystem",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runIndex(AddFilesCmd),
}

func init() {
	filesCmd.AddCommand(filesystemAddCmd)
}
