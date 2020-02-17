package cmd

import (
	"bytes"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

func AddFilesCmd(idx *archivist.Index, args []string) error {
	blob, err := archivist.NewBlob(args[0])
	if err != nil {
		return err
	}

	blobSha := blob.GetIntegrity().GetSha512()
	blobBlake2b := blob.GetIntegrity().GetBlake2B512()
	for _, b := range idx.Blobs {
		if b.Size != blob.Size {
			continue
		}
		if bytes.Compare(b.GetIntegrity().GetSha512(), blobSha) != 0 {
			continue
		}
		if bytes.Compare(b.GetIntegrity().GetBlake2B512(), blobBlake2b) != 0 {
			continue
		}
		logrus.Debug("File exists in index")
		return nil
	}
	logrus.Debug("Adding file to index")
	idx.Blobs = append(idx.Blobs, blob)
	return nil
}

var filesystemAddCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "Add filesystem",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runIndex(AddFilesCmd),
}

func
init() {
	filesCmd.AddCommand(filesystemAddCmd)
}
