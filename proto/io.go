package archivist

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func WriteProtoIndex(data proto.Message, path string) error {
	start := time.Now()

	var b []byte
	if filepath.Ext(path) == ".pb" {
		var err error
		if b, err = proto.Marshal(data); err != nil {
			return fmt.Errorf("marshaling index: %w", err)
		}
	} else {
		var err error
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		if b, err = json.MarshalIndent(data, "", "  "); err != nil {
			return fmt.Errorf("marshaling index: %w", err)
		}
	}

	if err := ioutil.WriteFile(path, b, 0600); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}
	logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Wrote index")
	return nil
}

func ReadProtoIndex(path string, data proto.Message) error {
	logrus.WithField("path", path).Debug("Reading proto index")
	start := time.Now()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading index file: %w", err)
	}
	if len(b) == 0 {
		return nil
	}

	if filepath.Ext(path) == ".pb" {
		if err := proto.Unmarshal(b, data); err != nil {
			return fmt.Errorf("unmarshaling index: %w", err)
		}
	} else {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		if err := json.Unmarshal(b, data); err != nil {
			return fmt.Errorf("unmarshaling index: %w", err)
		}
	}

	logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Read index")
	return nil
}
