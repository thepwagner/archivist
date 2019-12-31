package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thepwagner/archivist/index"
	archivist "github.com/thepwagner/archivist/proto"
)

func runIndex(run func(idx *archivist.Index, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		indexFn := viper.GetString("index")
		idx, err := index.ReadProtoIndex(indexFn)
		if err != nil {
			return err
		}

		if err := run(idx, args); err != nil {
			return err
		}

		if viper.GetBool("readonly") {
			return nil
		}
		return index.WriteProtoIndex(idx, indexFn)
	}
}
