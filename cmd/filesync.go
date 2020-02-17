package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

func SyncFiles(idx *archivist.Index, root string) error {
	logrus.WithField("root", root).Debug("Syncing files")
	// Verify provided path is a directory:
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	if err := ensureDir(absRoot); err != nil {
		return err
	}
	if idx.Filesystems == nil {
		idx.Filesystems = make(map[string]*archivist.Filesystem)
	}
	fs, ok := idx.Filesystems[root]
	if !ok {
		fs = &archivist.Filesystem{}
		idx.Filesystems[root] = fs
	}

	// Walk tree and index files
	newPaths := make(map[string]string, len(fs.Paths))
	indexWalk := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pathRel, err := filepath.Rel(absRoot, path)
		if err != nil {
			return err
		}
		file, err := AddFile(idx, path, info)
		if err != nil {
			return err
		}
		newPaths[pathRel] = file.GetId()
		return nil
	}
	if err := filepath.Walk(absRoot, indexWalk); err != nil {
		return err
	}
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		for p := range fs.GetPaths() {
			if _, ok := newPaths[p]; !ok {
				logrus.WithField("path", p).Debug("File no longer in index")
			}
		}
	}
	fs.Paths = newPaths
	return nil
}

func AddFile(idx *archivist.Index, path string, info os.FileInfo) (*archivist.Blob, error) {
	log := logrus.WithField("path", path)
	integrity, err := archivist.NewFileIntegrity(path)
	if err != nil {
		return nil, err
	}

	blobSha := integrity.GetSha512()
	blobBlake2b := integrity.GetBlake2B512()
	for _, b := range idx.Blobs {
		indexedBlob := b.GetIntegrity()
		if bytes.Compare(indexedBlob.GetSha512(), blobSha) != 0 {
			continue
		}
		if bytes.Compare(indexedBlob.GetBlake2B512(), blobBlake2b) != 0 {
			continue
		}
		log.Debug("File exists in index")
		return b, nil
	}

	blob := &archivist.Blob{
		Id:        uuid.NewV4().String(),
		Size:      uint64(info.Size()),
		Integrity: integrity,
	}
	log.Debug("Adding file to index")
	idx.Blobs = append(idx.Blobs, blob)
	return blob, nil
}

func ensureDir(path string) error {
	pathStat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stating path: %w", err)
	}
	if !pathStat.IsDir() {
		return fmt.Errorf("not directory: %q", path)
	}
	return nil
}

var filesystemAddCmd = &cobra.Command{
	Use:   "sync [path]",
	Short: "Sync filesystem",
	Args:  cobra.MinimumNArgs(1),
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		return SyncFiles(idx, args[0])
	}),
}

func init() {
	filesCmd.AddCommand(filesystemAddCmd)
}
