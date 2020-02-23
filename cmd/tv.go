package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var tvCmd = &cobra.Command{
	Use:   "tv [query]",
	Short: "Search TV shows",
	Args:  cobra.MinimumNArgs(1),
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		re, err := regexp.Compile(fmt.Sprintf("(?i)%s", args[0]))
		if err != nil {
			return fmt.Errorf("compiling regexp: %w", err)
		}
		results := archivist.FindTV(idx, re)

		out := os.Stdout
		_, _ = fmt.Fprint(out, "TV Results:\n")
		if len(results) == 0 {
			_, _ = fmt.Fprint(out, "(none)\n")
		} else {
			shows := make([]string, 0, len(results))
			for show := range results {
				shows = append(shows, show)
			}
			sort.Strings(shows)

			for _, show := range shows {
				_, _ = fmt.Fprintf(out, "%s\n", show)
				for _, fs := range results[show] {
					fsSummary := archivist.Summarize(idx, []string{fs}, fmt.Sprintf("video/tv/%s", show))
					_, _ = fmt.Fprintf(out, "  %s - %d files, %s\n", fs, fsSummary.FileCount, ByteCountSI(fsSummary.FileSizeSum))
				}
			}
		}

		return nil
	}),
}

func ByteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func init() {
	rootCmd.AddCommand(tvCmd)
}
