package index

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	archivist "github.com/thepwagner/archivist/proto"
	"io/ioutil"
	"os"
)

func LoadIndex(filename string) (*Index, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return NewIndex(), nil
		}
		return nil, fmt.Errorf("reading index file: %w", err)
	}

	var data archivist.Index
	if err := proto.Unmarshal(b, &data); err != nil {
		return nil, fmt.Errorf("unmarshaling index: %w", err)
	}

	idx := NewIndex()
	for _, blob := range data.Blobs {
		blobUUID := uuid.FromBytesOrNil(blob.GetBlobId())
		blobID := BlobID(blobUUID.String())
		idx.blobs[blobID] = blob
	}
	return idx, nil
}

func SaveIndex(idx *Index, filename string) error {
	data := newProtoIndex(idx)
	return writeProtoIndex(data, filename)
}

func newProtoIndex(idx *Index) *archivist.Index {
	var data archivist.Index
	for _, blob := range idx.blobs {
		data.Blobs = append(data.Blobs, blob)
	}
	return &data
}

func writeProtoIndex(data *archivist.Index, filename string) error {
	b, err := proto.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshaling index: %w", err)
	}
	if err := ioutil.WriteFile(filename, b, 0600); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}
	return nil
}