package index_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/archivist/index"
)

func TestSaveIndex_Empty(t *testing.T) {
	tmp := saveIndex(t, index.NewIndex())
	defer removeAll(t, tmp)

	stat, err := os.Stat(tmp)
	require.NoError(t, err)
	assert.Equal(t, int64(0), stat.Size())
}

func TestLoadIndex(t *testing.T) {
	idx := index.NewIndex()
	blobID, err := idx.Add("/etc/hosts")
	require.NoError(t, err)
	tmp := saveIndex(t, idx)
	defer removeAll(t, tmp)

	loaded, err := index.LoadIndex(tmp)
	require.NoError(t, err)

	loadedBlobID, err := loaded.Add("/etc/hosts")
	require.NoError(t, err)
	assert.Equal(t, blobID, loadedBlobID)
}

func saveIndex(t *testing.T, idx *index.Index) string {
	tmp, err := ioutil.TempFile("", "test-index-*")
	require.NoError(t, err)
	err = tmp.Close()
	require.NoError(t, err)

	err = index.SaveIndex(idx, tmp.Name())
	require.NoError(t, err)
	return tmp.Name()
}

func removeAll(t *testing.T, filename string) {
	err := os.RemoveAll(filename)
	require.NoError(t, err)
}
