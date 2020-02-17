package archivist_test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	archivist "github.com/thepwagner/archivist/proto"
)

const (
	testMessage = "hello world"
	testSha     = "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="
	testBlake   = "Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A=="
)

func TestNewIntegrity(t *testing.T) {
	i, err := archivist.NewIntegrity(strings.NewReader(testMessage))
	require.NoError(t, err)
	assertTestMessageIntegrity(t, i)
}

func TestNewFileIntegrity(t *testing.T) {
	f := testMessageFile(t)
	defer os.Remove(f.Name())

	i, err := archivist.NewFileIntegrity(f.Name())
	require.NoError(t, err)
	assertTestMessageIntegrity(t, i)
}

func TestNewBlob(t *testing.T) {
	f := testMessageFile(t)
	defer os.Remove(f.Name())

	b, err := archivist.NewBlob(f.Name())
	require.NoError(t, err)
	assert.Equal(t, uint64(len(testMessage)), b.Size)
	assertTestMessageIntegrity(t, b.GetIntegrity())
}

func TestNewBlob_NotFound(t *testing.T) {
	_, err := archivist.NewBlob("/must-not-exist")
	assert.Error(t, err)
}

func TestNewBlob_NotFile(t *testing.T) {
	tmp, err := ioutil.TempDir("", "new-blob-not-file-")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)
	_, err = archivist.NewBlob(tmp)
	assert.Error(t, err)
}

func testMessageFile(t *testing.T) *os.File {
	f, err := ioutil.TempFile("", "test-*")
	require.NoError(t, err)
	_, err = f.WriteString(testMessage)
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)
	return f
}

func assertTestMessageIntegrity(t *testing.T, i *archivist.Integrity) {
	assert.Equal(t, testSha, base64.StdEncoding.EncodeToString(i.GetSha512()))
	assert.Equal(t, testBlake, base64.StdEncoding.EncodeToString(i.GetBlake2B512()))
}
