package cmd

import (
	"fmt"
	"io"
	"os"
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

		// Parse arguments to filter output:
		var tv, movies, details bool
		var query []string
		for _, a := range args {
			if strings.Contains(a, "tv") {
				tv = true
			} else if strings.Contains(a, "movie") {
				movies = true
			} else if strings.Contains(a, "details") {
				details = true
			} else {
				query = append(query, a)
				details = true
			}
		}

		queryRe, err := buildQueryRE(query)
		if err != nil {
			return err
		}

		both := !tv && !movies
		if tv || both {
			dupes := archivist.FindTV(idx, queryRe)
			reportDupes(idx, out, details, dupes, "tv")
		}
		if movies || both {
			dupes := archivist.FindMovies(idx, queryRe)
			reportDupes(idx, out, details, dupes, "movies")
		}
		return nil
	}),
}

func reportDupes(idx *archivist.Index, out io.Writer, details bool, dupes map[string][]string, mediaType string) {
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
			if details {
				_, _ = fmt.Fprintf(out, "%s\n", show)
				for _, fs := range dupes[show] {
					fsSummary := archivist.Summarize(idx, fs, fmt.Sprintf("video/%s/%s", mediaType, show))
					_, _ = fmt.Fprintf(out, "  %s - %s\n", fs, fsSummary)
				}
			} else {
				_, _ = fmt.Fprintf(out, "  %s || %d copies \n", show, len(dupes[show]))
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(dupesCmd)
}
