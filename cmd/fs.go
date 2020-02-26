package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "List filesystems",
	RunE: runIndexRO(func(idx *archivist.Index, _ []string) error {
		return FilesystemList(idx, os.Stdout)
	}),
}

func FilesystemList(idx *archivist.Index, out io.Writer) error {
	names := make([]string, 0, len(idx.GetFilesystems()))
	for n := range idx.GetFilesystems() {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, name := range names {
		summary := archivist.Summarize(idx, name, "")
		if _, err := fmt.Fprintf(out, "%-20s %s\n", name, summary); err != nil {
			return err
		}
	}
	summary := archivist.Summarize(idx, "", "")
	if _, err := fmt.Fprintf(out, "%-20s %s\n", "TOTAL", summary); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(fsCmd)
}
