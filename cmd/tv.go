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
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		re, err := buildQueryRE(args)
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
					fsSummary := archivist.Summarize(idx, fs, fmt.Sprintf("video/tv/%s", show))
					_, _ = fmt.Fprintf(out, "  %s - %s\n", fs, fsSummary)
				}
			}
		}

		return nil
	}),
}

func buildQueryRE(args []string) (*regexp.Regexp, error) {
	if len(args) > 0 {
		return regexp.Compile(fmt.Sprintf("(?i)%s", args[0]))
	}
	return regexp.MustCompile("."), nil
}

func init() {
	rootCmd.AddCommand(tvCmd)
}
