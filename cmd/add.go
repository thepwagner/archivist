package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/thepwagner/archivist/index"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [file...]",
	Short: "Add a file to the index",
	Long: `Add file to index`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idx := loadIndex()

		paths := make(map[string]index.BlobID)
		for _, path := range args {
			blobID, err := idx.Add(path)
			if err != nil {
				return fmt.Errorf("indexing path %q: %w", path, err)
			}
			paths[path] = blobID
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(paths)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
