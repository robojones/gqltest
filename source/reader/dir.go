package reader

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
)

// Dir represents the contents of a directory on the disk.
// All contained paths include the directory name.
type Dir struct {
	Name    string
	Files   []string
	Subdirs []string
}

func newDir(name string, infos []os.FileInfo) *Dir {
	dir := &Dir{
		Name: name,
	}

	for _, info := range infos {
		n := path.Join(name, info.Name())

		if info.IsDir() {
			dir.Subdirs = append(dir.Subdirs, n)
		} else {
			dir.Files = append(dir.Files, n)
		}
	}

	return dir
}

func (r *Reader) ReadDir(dirname string) *Dir {
	infos, err := ioutil.ReadDir(dirname)

	if err != nil {
		panic(errors.Wrap(err, dirname))
	}

	return newDir(dirname, infos)
}
