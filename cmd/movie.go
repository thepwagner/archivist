package cmd

import (
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var movieCmd = &cobra.Command{
	Use:   "movie [query]",
	Short: "Search movies",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		return runSearch(idx, args, archivist.FindMovies, "video/movies")
	}),
}

func init() {
	rootCmd.AddCommand(movieCmd)
}
