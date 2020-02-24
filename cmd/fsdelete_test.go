package cmd_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/archivist/cmd"
	archivist "github.com/thepwagner/archivist/proto"
)

func TestFilesystemDelete_NilSafe(t *testing.T) {
	err := cmd.FilesystemDelete(nil, "", nil)
	require.Error(t, err)
}

func TestFilesystemDelete_NotFound(t *testing.T) {
	err := cmd.FilesystemDelete(&archivist.Index{}, "foo", nil)
	require.Error(t, err)
}

func TestFilesystemDelete(t *testing.T) {
	idx := &archivist.Index{
		Filesystems: map[string]*archivist.Filesystem{
			"foo": {},
			"bar": {},
		},
	}

	var buf bytes.Buffer
	err := cmd.FilesystemDelete(idx, "foo", &buf)
	require.NoError(t, err)
	assert.NotContains(t, idx.Filesystems, "foo")
	assert.Contains(t, idx.Filesystems, "bar")
	assert.Equal(t, "bar                  0 files, 0 B\n", buf.String())
}
