package cmd_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/archivist/cmd"
	archivist "github.com/thepwagner/archivist/proto"
)

var helloWorld = []byte("hello world")

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestSyncFiles(t *testing.T) {
	tmp := tmpTree(t)
	defer os.RemoveAll(tmp)

	idx := &archivist.Index{}
	err := cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)

	assert.Len(t, idx.GetBlobs(), 2)
	blobs := archivist.NewBlobIndex(idx.GetBlobs())
	assert.Len(t, idx.GetFilesystems(), 1)
	if assert.Contains(t, idx.GetFilesystems(), tmp) {
		fs := idx.Filesystems[tmp]
		assert.Len(t, fs.Paths, 2)

		assert.Contains(t, fs.Paths, "empty")
		emptyBlob := blobs.ByID[fs.Paths["empty"]]
		assert.Equal(t, uint64(0), emptyBlob.Size)

		helloBlob := blobs.ByID[fs.Paths["hello"]]
		assert.Equal(t, uint64(len(helloWorld)), helloBlob.Size)
	}
}

func TestSyncFiles_Add(t *testing.T) {
	tmp := tmpTree(t)
	defer os.RemoveAll(tmp)

	idx := &archivist.Index{}
	err := cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(tmp, "world"), []byte("hello world"), 0600)
	require.NoError(t, err)

	err = cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)
	assert.Len(t, idx.GetBlobs(), 2, "duplicate file blob")
	blobs := archivist.NewBlobIndex(idx.GetBlobs())

	helloBlob := blobs.ByID[idx.Filesystems[tmp].Paths["world"]]
	assert.Equal(t, uint64(len(helloWorld)), helloBlob.Size)
}

func TestSyncFiles_Remove(t *testing.T) {
	tmp := tmpTree(t)
	defer os.RemoveAll(tmp)

	idx := &archivist.Index{}
	err := cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)

	err = os.Remove(filepath.Join(tmp, "hello"))
	require.NoError(t, err)

	err = cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)
	assert.Len(t, idx.GetBlobs(), 2, "removed blob purged")
	assert.NotContains(t, idx.Filesystems[tmp].Paths, "hello")
}

func tmpTree(t *testing.T) string {
	tmp, err := ioutil.TempDir("", "sync-files-test-")
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(tmp, "empty"), nil, 0600)
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(tmp, "hello"), helloWorld, 0600)
	require.NoError(t, err)
	return tmp
}
