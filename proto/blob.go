package archivist

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/blake2b"
)

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
	blob := &Blob{
		Id:        uuid.NewV4().String(),
		Size:      size,
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
