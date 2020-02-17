package archivist

func (i *Index) GetFilesystem(root string) *Filesystem {
	if i.Filesystems == nil {
		i.Filesystems = make(map[string]*Filesystem)
	} else if fs, ok := i.Filesystems[root]; ok {
		return fs
	}
	fs := &Filesystem{}
	i.Filesystems[root] = fs
	return fs
}
