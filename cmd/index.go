package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	archivist "github.com/thepwagner/archivist/proto"
)

func runIndex(run func(idx *archivist.Index, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		indexFn := viper.GetString("index")
		var idx archivist.Index
		if err := archivist.ReadProtoIndex(indexFn, &idx); err != nil {
			return err
		}

		if err := run(&idx, args); err != nil {
			return err
		}

		if viper.GetBool("readonly") {
			return nil
		}
		return archivist.WriteProtoIndex(&idx, indexFn)
	}
}
