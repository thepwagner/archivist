package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var movieCmd = &cobra.Command{
	Use:   "movie [query]",
	Short: "Search movies",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		re, err := buildQueryRE(args)
		if err != nil {
			return fmt.Errorf("compiling regexp: %w", err)
		}

		results := archivist.FindMovies(idx, re)
		out := os.Stdout
		_, _ = fmt.Fprint(out, "Movie Results:\n")
		if len(results) == 0 {
			_, _ = fmt.Fprint(out, "(none)\n")
		} else {
			movies := make([]string, 0, len(results))
			for movie := range results {
				movies = append(movies, movie)
			}
			sort.Strings(movies)

			for _, movie := range movies {
				_, _ = fmt.Fprintf(out, "%s\n", movie)
				for _, fs := range results[movie] {
					fsSummary := archivist.Summarize(idx, []string{fs}, fmt.Sprintf("video/movies/%s", movie))
					_, _ = fmt.Fprintf(out, "  %s - %d files, %s\n", fs, fsSummary.FileCount, ByteCountSI(fsSummary.FileSizeSum))
				}
			}
		}

		return nil
	}),
}

func init() {
	rootCmd.AddCommand(movieCmd)
}
