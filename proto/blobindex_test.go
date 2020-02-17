package archivist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	archivist "github.com/thepwagner/archivist/proto"
)

func TestNewBlobIndex(t *testing.T) {
	blobs := archivist.NewBlobIndex([]*archivist.Blob{
		{Id: "foo"},
		{Id: "bar"},
	})
	assert.Len(t, blobs.ByID, 2)
	assert.Contains(t, blobs.ByID, "foo")
	assert.Contains(t, blobs.ByID, "bar")
}
