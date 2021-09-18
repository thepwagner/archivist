package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
)

const remoteName = "origin"

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull upstream repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// Open repo:
		indexPath := viper.GetString(flagIndex)
		repo, err := git.PlainOpen(indexPath)
		if err != nil {
			return fmt.Errorf("opening repo: %w", err)
		}

		logrus.WithFields(logrus.Fields{
			"repo":   indexPath,
			"remote": remoteName,
		}).Info("fetching remote...")
		if err := repoPull(ctx, repo); err != nil {
			return err
		}

		h, err := repo.Head()
		if err != nil {
			return fmt.Errorf("checking repo HEAD: %w", err)
		}
		commit, err := repo.CommitObject(h.Hash())
		if err != nil {
			return fmt.Errorf("checking repo HEAD commit: %w", err)
		}

		logrus.WithFields(logrus.Fields{
			"commit": h.Hash(),
			"date":   commit.Author.When.Format(time.RFC3339),
		}).Info("updated")

		return nil
	},
}

func repoPull(ctx context.Context, repo *git.Repository) error {
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("getting worktree: %w", err)
	}

	err = wt.PullContext(ctx, &git.PullOptions{
		RemoteName: remoteName,
		Progress:   os.Stdout,
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("pulling remote: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
