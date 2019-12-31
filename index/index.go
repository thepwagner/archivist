package index

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	archivist "github.com/thepwagner/archivist/proto"
)



type BlobID string

type Index struct {
	drives    map[DriveID]*archivist.Drive
	blobs     map[BlobID]*archivist.Blob
	filenames map[string][]BlobID
}

func NewIndex() *Index {
	return &Index{
		drives:    make(map[DriveID]*archivist.Drive),
		blobs:     make(map[BlobID]*archivist.Blob),
		filenames: make(map[string][]BlobID),
	}
}

func (i *Index) Add(path string) (BlobID, error) {
	log := logrus.WithField("path", path)
	log.Debug("Adding file to index")

	stat, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("stat path: %w", err)
	}
	if stat.IsDir() {
		return "", errors.New("path must be file")
	}
	size := uint64(stat.Size())

	// Check for existing Blob:
	base := filepath.Base(path)
	if fnBlobIDs := i.filenames[base]; len(fnBlobIDs) > 0 {
		log.WithField("blobs", len(fnBlobIDs)).Debug("Matched by filename")
		for _, blobID := range fnBlobIDs {
			if size == i.blobs[blobID].Size {
				// XXX: maybe verify hashes too?
				log.WithField("blob_id", blobID).Debug("Matched by size")
				return blobID, nil
			}
		}
	}

	blobUUID := uuid.NewV4().String()
	blobID := BlobID(blobUUID)
	log.WithField("blob_id", blobID).Debug("Adding to index")
	i.blobs[blobID] = &archivist.Blob{
		Id:   blobUUID,
		Size: size,
		// TODO: read hashes
	}
	i.filenames[base] = append(i.filenames[base], blobID)
	return blobID, nil
}
