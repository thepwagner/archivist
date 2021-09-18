package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var tvCmd = &cobra.Command{
	Use:   "tv [query]",
	Short: "Search TV shows",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		return runSearch(idx, args, archivist.FindTV, "video/tv")
	}),
}

func runSearch(
	idx *archivist.Index,
	args []string,
	searchFunc func(*archivist.Index, *regexp.Regexp) map[string][]string,
	prefix string,
) error {
	re, err := buildQueryRE(args)
	if err != nil {
		return fmt.Errorf("compiling regexp: %w", err)
	}

	results := searchFunc(idx, re)
	out := os.Stdout
	_, _ = fmt.Fprint(out, "Results:\n")
	if len(results) == 0 {
		_, _ = fmt.Fprint(out, "(none)\n")
	} else {
		keys := make([]string, 0, len(results))
		for result := range results {
			keys = append(keys, result)
		}
		sort.Strings(keys)

		for _, k := range keys {
			_, _ = fmt.Fprintf(out, "%s\n", k)
			for _, fs := range results[k] {
				fsSummary := archivist.Summarize(idx, fs, filepath.Join(prefix, k))
				_, _ = fmt.Fprintf(out, "  %s - %s\n", fs, fsSummary)
			}
		}
	}
	return nil
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
