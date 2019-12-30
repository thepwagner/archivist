package cmd

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thepwagner/archivist/index"
)

func runIndex(run func(cmd *cobra.Command, args []string, idx *index.Index) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		indexFn := viper.GetString("index")
		idx, err := loadIndex(indexFn)
		if err != nil {
			return err
		}

		if err := run(cmd, args, idx); err != nil {
			return err
		}

		if viper.GetBool("readonly") {
			return nil
		}
		return saveIndex(idx, indexFn)
	}
}

func loadIndex(filename string) (*index.Index, error) {
	start := time.Now()
	idx, err := index.LoadIndex(filename)
	if err != nil {
		return nil, fmt.Errorf("loading index: %w", err)
	}
	logrus.WithFields(logrus.Fields{
		"path": filename,
		"dur":  time.Since(start).Truncate(time.Millisecond).Seconds(),
	}).Debug("Loaded index")
	return idx, nil
}

func saveIndex(idx *index.Index, filename string) error {
	start := time.Now()
	err := index.SaveIndex(idx, filename)
	if err != nil {
		return fmt.Errorf("saving index: %w", err)
	}
	logrus.WithFields(logrus.Fields{
		"path": filename,
		"dur":  time.Since(start).Truncate(time.Millisecond).Seconds(),
	}).Debug("Saved index")
	return nil
}
