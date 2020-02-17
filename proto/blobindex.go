package archivist

type BlobIndex struct {
	ByID map[string]*Blob
}

func NewBlobIndex(blobs []*Blob) *BlobIndex {
	bi := &BlobIndex{
		ByID: make(map[string]*Blob, len(blobs)),
	}
	for _, b := range blobs {
		bi.ByID[b.Id] = b
	}
	return bi
}
