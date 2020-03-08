package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	archivist "github.com/thepwagner/archivist/proto"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const (
	indexProtoFilename = "index.pb"
	indexFilename      = "index.json"
)

func runIndex(run func(idx *archivist.Index, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		indexPath := viper.GetString(flagIndex)
		repo, err := git.PlainOpen(indexPath)
		if err != nil {
			return fmt.Errorf("opening repo: %w", err)
		}
		wt, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("getting repo worktree: %w", err)
		}

		var idx archivist.Index
		indexFile := filepath.Join(indexPath, indexFilename)
		if err := archivist.ReadProtoIndex(indexFile, &idx); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			indexProtoFile := filepath.Join(indexPath, indexProtoFilename)
			if err := archivist.ReadProtoIndex(indexProtoFile, &idx); err != nil {
				return err
			}
		}

		message := strings.Join(os.Args[1:], " ")

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			if err := writeIndex(idx, indexFile, wt, message); err != nil {
				logrus.WithError(err).Warn("Error writing index")
			}
			os.Exit(1)
		}()

		defer func() {
			if err := recover(); err != nil {
				logrus.WithField("err", err).Error("caught panic")
				if err := writeIndex(idx, indexFile, wt, message); err != nil {
					logrus.WithError(err).Warn("Error writing index")
				}
				os.Exit(1)
			}
		}()

		start := time.Now()
		if err := run(&idx, args); err != nil {
			return err
		}
		logrus.WithField("dur", logDur(start)).Debug("Ran command")

		return writeIndex(idx, indexFile, wt, message)
	}
}

func writeIndex(idx archivist.Index, indexFn string, wt *git.Worktree, message string) error {
	if viper.GetBool(flagReadOnly) {
		logrus.Warn("Read-only mode, skipping index write")
		return nil
	}
	if err := archivist.WriteProtoIndex(&idx, indexFn); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}
	return commitIndex(wt, message)
}

func commitIndex(wt *git.Worktree, message string) error {
	start := time.Now()
	if _, err := wt.Add(indexFilename); err != nil {
		return fmt.Errorf("adding file: %w", err)
	}
	commit, err := wt.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Archivist",
			Email: "archivist@mycloudand.me",
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("commiting index: %w", err)
	}
	logrus.WithFields(logrus.Fields{
		"commit_id": commit.String(),
		"dur":       logDur(start),
	}).Info("Index committed")
	return nil
}

func runIndexRO(run func(idx *archivist.Index, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, args []string) error {
		indexPath := viper.GetString(flagIndex)
		if _, err := git.PlainOpen(indexPath); err != nil {
			return fmt.Errorf("opening repo: %w", err)
		}

		indexFile := filepath.Join(indexPath, indexFilename)
		var idx archivist.Index
		if err := archivist.ReadProtoIndex(indexFile, &idx); err != nil {
			return err
		}

		start := time.Now()
		if err := run(&idx, args); err != nil {
			return err
		}
		logrus.WithField("dur", logDur(start)).Debug("Ran command")

		return nil
	}
}

func logDur(start time.Time) int64 {
	return time.Since(start).Truncate(time.Millisecond).Milliseconds()
}
