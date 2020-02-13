package cmd

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	archivist "github.com/thepwagner/archivist/proto"
	"golang.org/x/crypto/blake2b"
)

func AddFilesCmd(idx *archivist.Index, args []string) error {
	blob, err := NewBlob(args[0])
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

func NewBlob(path string) (*archivist.Blob, error) {
	log := logrus.WithField("path", path)
	log.Debug("Reading file as blob")
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat path: %w", err)
	}
	if stat.IsDir() {
		return nil, errors.New("path must be file")
	}
	size := uint64(stat.Size())

	integrity, err := NewFileIntegrity(path)
	if err != nil {
		return nil, err
	}
	// TODO: hash

	log.WithFields(logrus.Fields{
		"size":    size,
		"sha512":  base64.StdEncoding.EncodeToString(integrity.Sha512),
		"blake2b": base64.StdEncoding.EncodeToString(integrity.Blake2B512),
	}).Debug("Read file blob")
	blob := &archivist.Blob{
		Size:      size,
		Integrity: integrity,
	}

	return blob, nil
}

func NewFileIntegrity(path string) (*archivist.Integrity, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewIntegrity(f)
}

func NewIntegrity(r io.Reader) (*archivist.Integrity, error) {
	shaHasher := sha512.New()
	blakeHasher, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(io.MultiWriter(shaHasher, blakeHasher), r); err != nil {
		return nil, err
	}
	return &archivist.Integrity{
		Sha512:     shaHasher.Sum(nil),
		Blake2B512: blakeHasher.Sum(nil),
	}, nil
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
