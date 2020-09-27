package archivist_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	archivist "github.com/thepwagner/archivist/proto"
)

func TestWriteProtoIndex(t *testing.T) {
	tmp := writeTempIndex(t)
	defer os.Remove(tmp.Name())

	stat, err := os.Stat(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, int64(48), stat.Size())
}

func TestReadProtoIndex(t *testing.T) {
	tmp := writeTempIndex(t)
	defer os.Remove(tmp.Name())

	var i archivist.Index
	err := archivist.ReadProtoIndex(tmp.Name(), &i)
	require.NoError(t, err)
	assert.Equal(t, "1", i.Blobs[0].GetId())
}

func writeTempIndex(t *testing.T) *os.File {
	tmp, err := ioutil.TempFile("", "test-index-*")
	require.NoError(t, err)
	i := &archivist.Index{}
	i.Blobs = append(i.Blobs, &archivist.Blob{
		Id: "1",
	})
	err = archivist.WriteProtoIndex(i, tmp.Name())
	require.NoError(t, err)
	return tmp
}
