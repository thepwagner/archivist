package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var dupesCmd = &cobra.Command{
	Use:   "dupes",
	Short: "Find duplicate data",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		out := os.Stdout
		reportDupes(idx, out, archivist.FindTV, "tv")
		reportDupes(idx, out, archivist.FindMovies, "movies")
		return nil
	}),
}

func reportDupes(idx *archivist.Index, out io.Writer, lookup func(*archivist.Index, *regexp.Regexp) map[string][]string, mediaType string) {
	dupes := lookup(idx, regexp.MustCompile("."))
	var names []string
	for show, fses := range dupes {
		if len(fses) > 1 {
			names = append(names, show)
		}
	}
	if len(names) > 0 {
		_, _ = fmt.Fprintf(out, "%s Results\n", strings.Title(mediaType))
		_, _ = fmt.Fprintln(out, strings.Repeat("-", len(mediaType)+8))
		sort.Strings(names)
		for _, show := range names {
			_, _ = fmt.Fprintf(out, "%s\n", show)
			for _, fs := range dupes[show] {
				fsSummary := archivist.Summarize(idx, fs, fmt.Sprintf("video/%s/%s", mediaType, show))
				_, _ = fmt.Fprintf(out, "  %s - %s\n", fs, fsSummary)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(dupesCmd)
}
