package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var lsCmd = &cobra.Command{
	Use:   "ls [path]",
	Short: "List merged filesystem",
	RunE: runIndexRO(func(idx *archivist.Index, args []string) error {
		var prefix string
		if len(args) > 0 {
			prefix = args[0]
		} else {
			prefix = ""
		}
		prefixSplit := strings.Split(prefix, "/")

		summaries := map[string]archivist.PathSummary{}
		var entries []string
		for _, fs := range idx.GetFilesystems() {
			for path := range fs.GetPaths() {
				if !strings.HasPrefix(path, prefix) {
					continue
				}

				pathSplit := strings.Split(path, "/")
				next := pathSplit[len(prefixSplit)-1]
				if _, ok := summaries[next]; ok {
					continue
				}

				summaryPrefix := strings.Join(prefixSplit[:len(prefixSplit)-1], "/")
				if len(prefixSplit) > 1 {
					summaryPrefix += "/"
				}
				summaryPrefix += next
				if len(pathSplit)-len(prefixSplit) > 0 {
					summaryPrefix += "/"
				}

				summaries[next] = archivist.Summarize(idx, "", summaryPrefix)
				entries = append(entries, next)
			}
		}

		sort.Strings(entries)
		for _, p := range entries {
			fmt.Printf("%s - %s\n", p, summaries[p])
		}
		return nil
	}),
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
