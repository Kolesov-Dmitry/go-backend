package dir

import (
	"os"
	"path/filepath"

	"file-srv/internal/file"
)

// Reader is file system directory reader
// implements srv.DirReader
type Reader struct{}

// Read read file system directory
// Inputs:
//   path - path to the directory to read
// Output:
//   Returns list of files in case of success, otherwise returns error
func (r *Reader) Read(path string) ([]*file.FileInfo, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]*file.FileInfo, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			// skip sirectories
			continue
		}

		info, err := f.Info()
		if err != nil {
			continue
		}

		result = append(result, &file.FileInfo{
			Name: f.Name(),
			Ext:  filepath.Ext(f.Name()),
			Size: info.Size(),
		})
	}

	return result, nil
}
