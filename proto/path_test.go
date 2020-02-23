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
					"video/tv/Show 1/s1e1.mkv":              {},
					"video/movies/Movie 1 (2020)/movie.mkv": {},
				},
			},
		},
	}

	t.Run("movies ignored", func(t *testing.T) {
		res := archivist.FindTV(idx, regexp.MustCompile(".*ovie.*"))
		assert.Empty(t, res)
	})

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

func TestFindMovies(t *testing.T) {
	idx := &archivist.Index{
		Filesystems: map[string]*archivist.Filesystem{
			"fs1": {
				Paths: map[string]*archivist.File{
					"video/movies/Movie 1 (2020)/movie.mkv": {},
					"video/movies/Movie 2 (2020)/movie.mkv": {},
				},
			},
			"fs2": {
				Paths: map[string]*archivist.File{
					"video/movies/Movie 1 (2020)/movie.mkv": {},
					"video/tv/Show 1/s1e1.mkv":              {},
				},
			},
		},
	}

	t.Run("tv ignored", func(t *testing.T) {
		res := archivist.FindMovies(idx, regexp.MustCompile(".*how 1.*"))
		assert.Empty(t, res)
	})

	t.Run("single filesystem", func(t *testing.T) {
		res := archivist.FindMovies(idx, regexp.MustCompile(".*vie 2.*"))
		assert.Equal(t, map[string][]string{
			"Movie 2 (2020)": {"fs1"},
		}, res)
	})

	t.Run("multiple filesystems", func(t *testing.T) {
		res := archivist.FindMovies(idx, regexp.MustCompile(".*vie 1.*"))
		assert.Equal(t, map[string][]string{
			"Movie 1 (2020)": {"fs1", "fs2"},
		}, res)

		res = archivist.FindMovies(idx, regexp.MustCompile(".*"))
		assert.Equal(t, map[string][]string{
			"Movie 1 (2020)": {"fs1", "fs2"},
			"Movie 2 (2020)": {"fs1"},
		}, res)
	})
}
