package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/ptypes"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
)

func SyncFilesystem(idx *archivist.Index, root string) error {
	logrus.WithField("root", root).Debug("Syncing filesystem")
	// Verify provided path is a directory:
	rootDir, err := ensureDir(root)
	if err != nil {
		return err
	}

	// Walk tree and index files
	fs := idx.GetFilesystem(root)
	blobs := archivist.NewBlobIndex(idx.GetBlobs())
	newPaths := make(map[string]string, len(fs.Paths))
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pathRel, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}
		log := logrus.WithField("path", path)

		// If file exists with the same size+mtime, skip integrity calculation:
		if oldBlobID, ok := fs.Paths[pathRel]; ok {
			if oldBlob, ok := blobs.ByID[oldBlobID]; ok {
				if oldBlob.Size == uint64(info.Size()) {
					oldBlobModTime, err := ptypes.Timestamp(oldBlob.GetModTime())
					if err == nil && oldBlobModTime == info.ModTime().UTC() {
						log.WithField("blob_id", oldBlobID).Debug("Path matched existing blob")
						newPaths[pathRel] = oldBlobID
						return nil
					}
				}
			}
		}

		// File does not exist, add to blob index:
		// FIXME: if a file is "touched", we passed through the above (mtime mismatch) and will calculate integrity
		// AddBlob() will match only on hash, so it won't update mtime of the blob - we'll re-hash every time.
		blob, err := AddBlob(idx, path, info)
		if err != nil {
			return err
		}
		log.WithField("blob_id", blob.Id).Debug("Indexed new path")
		newPaths[pathRel] = blob.GetId()
		return nil
	}
	if err := filepath.Walk(rootDir, walkFunc); err != nil {
		return err
	}

	// Log paths no longer in this filesystem:
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

func AddBlob(idx *archivist.Index, path string, info os.FileInfo) (*archivist.Blob, error) {
	integrity, err := archivist.NewFileIntegrity(path)
	if err != nil {
		return nil, err
	}

	blobSha := integrity.GetSha512()
	log := logrus.WithField("sha512", base64.StdEncoding.EncodeToString(blobSha))
	blobBlake2b := integrity.GetBlake2B512()
	for _, b := range idx.Blobs {
		indexedBlob := b.GetIntegrity()
		if bytes.Compare(indexedBlob.GetSha512(), blobSha) != 0 {
			continue
		}
		if bytes.Compare(indexedBlob.GetBlake2B512(), blobBlake2b) != 0 {
			continue
		}
		log.WithField("blob_id", b.Id).Debug("Blob exists in index")
		return b, nil
	}

	modTime, err := ptypes.TimestampProto(info.ModTime())
	if err != nil {
		return nil, err
	}
	blob := &archivist.Blob{
		Id:        uuid.NewV4().String(),
		ModTime:   modTime,
		Size:      uint64(info.Size()),
		Integrity: integrity,
	}
	log.WithField("blob_id", blob.Id).Debug("Adding blob to index")
	idx.Blobs = append(idx.Blobs, blob)
	return blob, nil
}

func ensureDir(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("abs path: %w", err)
	}
	pathStat, err := os.Stat(abs)
	if err != nil {
		return "", fmt.Errorf("stat path: %w", err)
	}
	if !pathStat.IsDir() {
		return "", fmt.Errorf("not directory: %q", path)
	}
	return abs, nil
}

var filesystemAddCmd = &cobra.Command{
	Use:   "sync [path]",
	Short: "Sync filesystem",
	Args:  cobra.MinimumNArgs(1),
	RunE: runIndex(func(idx *archivist.Index, args []string) error {
		return SyncFilesystem(idx, args[0])
	}),
}

func init() {
	filesCmd.AddCommand(filesystemAddCmd)
}
