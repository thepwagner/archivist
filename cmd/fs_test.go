package cmd_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/archivist/cmd"
	archivist "github.com/thepwagner/archivist/proto"
)

func TestFilesystemList_NilSafe(t *testing.T) {
	err := cmd.FilesystemList(nil, nil)
	require.NoError(t, err)
}

func TestFilesystemList_Empty(t *testing.T) {
	idx := &archivist.Index{
		Filesystems: map[string]*archivist.Filesystem{
			"/foo": {},
		},
	}

	var buf bytes.Buffer
	err := cmd.FilesystemList(idx, &buf)
	require.NoError(t, err)
	assert.Equal(t, "/foo                 0 files, 0 B\n", buf.String())
}

func TestFilesystemList_Summary(t *testing.T) {
	idx := &archivist.Index{
		Blobs: []*archivist.Blob{
			{Id: "blob1", Size: 1},
			{Id: "blob2", Size: 2},
			{Id: "blob3", Size: 4},
		},
		Filesystems: map[string]*archivist.Filesystem{
			"/blob5": {
				Paths: map[string]*archivist.File{
					"blob1": {BlobId: "blob1"},
					"blob3": {BlobId: "blob3"},
				},
			},
			"/blob2": {
				Paths: map[string]*archivist.File{
					"blob2": {BlobId: "blob2"},
				},
			},
		},
	}

	var buf bytes.Buffer
	err := cmd.FilesystemList(idx, &buf)
	require.NoError(t, err)
	assert.Equal(t,
		"/blob2               1 files, 2 B\n"+
			"/blob5               2 files, 5 B\n", buf.String())
}
