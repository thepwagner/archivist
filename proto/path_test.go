package archivist_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	archivist "github.com/thepwagner/archivist/proto"
)

func TestFindTV(t *testing.T) {
	t.Run("nilsafe", func(t *testing.T) {
		res := archivist.FindTV(nil, regexp.MustCompile(".*"))
		assert.Empty(t, res)
	})

	// "show 1" mirrored in {fs1,fs2}
	// "show 2" stored on {fs1}
	// "movie" ignored
	idx := &archivist.Index{
		Filesystems: map[string]*archivist.Filesystem{
			"fs1": {
				Paths: map[string]*archivist.File{
					"video/tv/Show 1/s1e1.mkv":      {},
					"video/tv/Show 2 (UK)/s1e1.mkv": {},
				},
			},
			"fs2": {
				Paths: map[string]*archivist.File{
					"video/tv/Show 1/s1e1.mkv":            {},
					"video/movies/Movie (2020)/movie.mkv": {},
				},
			},
		},
	}

	t.Run("single filesystem", func(t *testing.T) {
		res := archivist.FindTV(idx, regexp.MustCompile(".*ow 2.*"))
		assert.Equal(t, map[string][]string{
			"Show 2 (UK)": {"fs1"},
		}, res)
	})

	t.Run("multiple filesystems", func(t *testing.T) {
		res := archivist.FindTV(idx, regexp.MustCompile(".*ow 1.*"))
		assert.Equal(t, map[string][]string{
			"Show 1": {"fs1", "fs2"},
		}, res)

		res = archivist.FindTV(idx, regexp.MustCompile(".*"))
		assert.Equal(t, map[string][]string{
			"Show 1":      {"fs1", "fs2"},
			"Show 2 (UK)": {"fs1"},
		}, res)
	})
}
