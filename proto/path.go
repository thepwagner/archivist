package archivist

import (
	"regexp"
	"sort"
	"strings"
)

var (
	tvRe    = regexp.MustCompile("tv/([A-Za-z0-9 ()]+)/")
	movieRe = regexp.MustCompile("movies/([A-Za-z0-9 ()]+)/")
)

// FindTV searches for a TV show in the archive.
// Returns a map of show name to filesystems storing _some_ part of it.
func FindTV(idx *Index, re *regexp.Regexp) map[string][]string {
	return findMedia(idx, tvRe, re)
}

// FindMovies searches for a movies in the archive.
// Returns a map of Movie name to filesystems storing _some_ part of it.
func FindMovies(idx *Index, re *regexp.Regexp) map[string][]string {
	return findMedia(idx, movieRe, re)
}

func findMedia(idx *Index, mediaRe *regexp.Regexp, pathRe *regexp.Regexp) map[string][]string {
	res := map[string][]string{}
	for fsName, fs := range idx.GetFilesystems() {
		// We operate on files, so we'll see the same shows (in directory path) 3x the number of episodes.
		done := map[string]struct{}{}
		for p := range fs.GetPaths() {
			if m := mediaRe.FindStringSubmatch(p); len(m) > 0 {
				showName := m[1]
				if _, ok := done[showName]; ok {
					continue
				}
				if pathRe.MatchString(showName) {
					res[showName] = append(res[showName], fsName)
				}
				done[showName] = struct{}{}
			}
		}
	}

	for _, fses := range res {
		sort.Strings(fses)
	}

	return res
}

type PathSummary struct {
	FileCount   uint64
	FileSizeSum uint64
}

func Summarize(idx *Index, filesystems []string, prefix string) PathSummary {
	// If no filesystems are specified, search all filesystems:
	if len(filesystems) == 0 {
		idxFilesystems := idx.GetFilesystems()
		filesystems = make([]string, 0, len(idxFilesystems))
		for fs := range idxFilesystems {
			filesystems = append(filesystems, fs)
		}
	}

	blobs := NewBlobIndex(idx.GetBlobs())

	var res PathSummary
	for _, fsName := range filesystems {
		fs := idx.GetFilesystem(fsName)
		for p, f := range fs.GetPaths() {
			if !strings.HasPrefix(p, prefix) {
				continue
			}
			res.FileCount += 1
			res.FileSizeSum += blobs.ByID[f.BlobId].Size
		}
	}
	return res
}
