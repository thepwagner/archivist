package index

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	archivist "github.com/thepwagner/archivist/proto"
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

	idx := fromProtoIndex(data)
	return idx, nil
}

func SaveIndex(idx *Index, filename string) error {
	data := toProtoIndex(idx)
	return writeProtoIndex(data, filename)
}

func toProtoIndex(idx *Index) *archivist.Index {
	var data archivist.Index
	for _, blob := range idx.blobs {
		data.Blobs = append(data.Blobs, blob)
	}

	data.BlobFilenames = make(map[string]*archivist.IDs, len(idx.filenames))
	for fn, blobIDs := range idx.filenames {
		var ids archivist.IDs
		for _, blobID := range blobIDs {
			ids.Ids = append(ids.Ids, string(blobID))
		}
		data.BlobFilenames[fn] = &ids
	}

	return &data
}

func fromProtoIndex(data archivist.Index) *Index {
	idx := NewIndex()
	for _, blob := range data.Blobs {
		blobID := BlobID(blob.GetId())
		idx.blobs[blobID] = blob
	}

	for fn, fnBlobIDs := range data.BlobFilenames {
		ids := fnBlobIDs.GetIds()
		blobIDs := make([]BlobID, 0, len(ids))
		for _, fnBlobID := range ids {
			blobIDs = append(blobIDs, BlobID(fnBlobID))
		}
		idx.filenames[fn] = blobIDs
	}
	return idx
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
