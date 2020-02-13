package index

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	archivist "github.com/thepwagner/archivist/proto"
)

func WriteProtoIndex(data *archivist.Index, filename string) error {
	b, err := proto.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshaling index: %w", err)
	}
	if err := ioutil.WriteFile(filename, b, 0600); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}
	return nil
}

func ReadProtoIndex(path string) (*archivist.Index, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &archivist.Index{}, nil
		}
		return nil, fmt.Errorf("reading index file: %w", err)
	}

	var data archivist.Index
	if err := proto.Unmarshal(b, &data); err != nil {
		return nil, fmt.Errorf("unmarshaling index: %w", err)
	}
	return &data, nil
}
