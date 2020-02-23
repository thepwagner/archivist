package archivist

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

func WriteProtoIndex(data proto.Message, path string) error {
	start := time.Now()
	b, err := proto.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshaling index: %w", err)
	}
	if err := ioutil.WriteFile(path, b, 0600); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}
	logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Wrote index")
	return nil
}

func ReadProtoIndex(path string, data proto.Message) error {
	start := time.Now()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("reading index file: %w", err)
	}
	if len(b) == 0 {
		return nil
	}

	if err := proto.Unmarshal(b, data); err != nil {
		return fmt.Errorf("unmarshaling index: %w", err)
	}
	logrus.WithField("dur", time.Since(start).Truncate(time.Millisecond).Milliseconds()).Debug("Read index")
	return nil
}
