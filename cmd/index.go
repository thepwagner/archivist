package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
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

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			if !viper.GetBool("readonly") {
				archivist.WriteProtoIndex(&idx, indexFn)
			}
			os.Exit(1)
		}()

		start := time.Now()
		if err := run(&idx, args); err != nil {
			return err
		}
		logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Ran command")

		if viper.GetBool("readonly") {
			return nil
		}
		return archivist.WriteProtoIndex(&idx, indexFn)
	}
}

func runIndexRO(run func(idx *archivist.Index, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		indexFn := viper.GetString("index")

		var idx archivist.Index
		if err := archivist.ReadProtoIndex(indexFn, &idx); err != nil {
			return err
		}

		start := time.Now()
		if err := run(&idx, args); err != nil {
			return err
		}
		logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Ran command")

		return nil
	}
}
