package filesystem

import "os"

type FileSystem interface {
	Rename(oldPath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}

//인터페이스를 사용함
func ManageFiles(fs FileSystem) {
	//todo : 이건 어떻게 사용할 있나?
}
