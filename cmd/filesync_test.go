package cmd_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

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
		emptyBlob := blobs.ByID[fs.Paths["empty"].BlobId]
		assert.Equal(t, uint64(0), emptyBlob.Size)

		helloBlob := blobs.ByID[fs.Paths["hello"].BlobId]
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

	helloBlob := blobs.ByID[idx.Filesystems[tmp].Paths["world"].BlobId]
	assert.Equal(t, uint64(len(helloWorld)), helloBlob.Size)
}

func TestSyncFiles_Touch(t *testing.T) {
	tmp := tmpTree(t)
	defer os.RemoveAll(tmp)

	idx := &archivist.Index{}
	err := cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)

	// Snapshot:
	beforeFS := idx.GetFilesystem(tmp)
	beforeHello := beforeFS.Paths["hello"]
	helloBlobID := beforeHello.BlobId
	helloMtime := beforeHello.ModTime.Nanos
	beforeEmpty := beforeFS.Paths["empty"]
	emptyBlobID := beforeEmpty.BlobId
	emptyMtime := beforeEmpty.ModTime.Nanos

	// Touch and resync:
	now := time.Now()
	err = os.Chtimes(filepath.Join(tmp, "hello"), now, now)
	require.NoError(t, err)
	err = cmd.SyncFilesystem(idx, tmp)
	require.NoError(t, err)

	afterFS := idx.GetFilesystem(tmp)
	afterHello := afterFS.Paths["hello"]
	assert.Equal(t, helloBlobID, afterHello.BlobId)
	assert.NotEqual(t, helloMtime, afterHello.ModTime.Nanos)
	afterEmpty := afterFS.Paths["empty"]
	assert.Equal(t, emptyBlobID, afterEmpty.BlobId)
	assert.Equal(t, emptyMtime, afterEmpty.ModTime.Nanos)

	assert.Len(t, idx.GetBlobs(), 2)
	blobs := archivist.NewBlobIndex(idx.GetBlobs())
	helloBlob := blobs.ByID[afterHello.BlobId]
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
