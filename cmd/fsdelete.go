package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

var fsDeleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove filesystem",
	Args:  cobra.MinimumNArgs(1),
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		return FilesystemDelete(idx, args[0], os.Stdout)
	}),
}

func FilesystemDelete(idx *archivist.Index, fs string, out io.Writer) error {
	logrus.WithField("fs", fs).Info("Delete filesystem")
	if _, ok := idx.GetFilesystems()[fs]; !ok {
		return fmt.Errorf("filesystem not found: %s", fs)
	}
	delete(idx.Filesystems, fs)
	return FilesystemList(idx, out)
}

func init() {
	fsCmd.AddCommand(fsDeleteCmd)
}
