package archivist

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/golang/protobuf/ptypes"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/blake2b"
)

type BlobIndex struct {
	ByID map[string]*Blob
}

func NewBlobIndex(blobs []*Blob) *BlobIndex {
	bi := &BlobIndex{
		ByID: make(map[string]*Blob, len(blobs)),
	}
	for _, b := range blobs {
		bi.ByID[b.Id] = b
	}
	return bi
}

func NewBlob(path string) (*Blob, error) {
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

	log.WithFields(logrus.Fields{
		"size":    size,
		"sha512":  base64.StdEncoding.EncodeToString(integrity.Sha512),
		"blake2b": base64.StdEncoding.EncodeToString(integrity.Blake2B512),
	}).Debug("Read file blob")
	modTime, err := ptypes.TimestampProto(stat.ModTime())
	if err != nil {
		return nil, err
	}
	blob := &Blob{
		Id:        uuid.NewV4().String(),
		Size:      size,
		ModTime:   modTime,
		Integrity: integrity,
	}

	return blob, nil
}

func NewFileIntegrity(path string) (*Integrity, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewIntegrity(f)
}

func NewIntegrity(r io.Reader) (*Integrity, error) {
	shaHasher := sha512.New()
	blakeHasher, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(io.MultiWriter(shaHasher, blakeHasher), r); err != nil {
		return nil, err
	}
	return &Integrity{
		Sha512:     shaHasher.Sum(nil),
		Blake2B512: blakeHasher.Sum(nil),
	}, nil
}
