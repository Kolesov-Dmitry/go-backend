package srv

import (
	"file-srv/internal/file"
)

// DirReader is an abstraction above file system directory reader
type DirReader interface {
	Read(path string) ([]*file.FileInfo, error)
}
