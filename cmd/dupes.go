package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var dupesCmd = &cobra.Command{
	Use:   "dupes",
	Short: "Find duplicate data",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		out := os.Stdout
		reportTvDupes(idx, out)
		return nil
	}),
}

func reportTvDupes(idx *archivist.Index, out io.Writer) {
	tvShows := archivist.FindTV(idx, regexp.MustCompile("."))
	var tvDupes []string
	for show, fses := range tvShows {
		if len(fses) > 1 {
			tvDupes = append(tvDupes, show)
		}
	}
	if len(tvDupes) > 0 {
		_, _ = fmt.Fprint(out, "TV Results:\n")
		sort.Strings(tvDupes)
		for _, show := range tvDupes {
			_, _ = fmt.Fprintf(out, "%s\n", show)
			for _, fs := range tvShows[show] {
				fsSummary := archivist.Summarize(idx, []string{fs}, fmt.Sprintf("video/tv/%s", show))
				_, _ = fmt.Fprintf(out, "  %s - %d files, %s\n", fs, fsSummary.FileCount, ByteCountSI(fsSummary.FileSizeSum))
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(dupesCmd)
}
