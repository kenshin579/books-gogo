package filesystem

import (
	"fmt"
	"os"
)

type FileSystem interface {
	Rename(oldPath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	fmt.Printf("rename %s -> %s\n", oldpath, newpath)

	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	fmt.Printf("remove %s\n", name)
	return os.Remove(name)
}

//인터페이스를 사용함
func ManageFiles(fs FileSystem) {
	//todo : 이건 어떻게 사용할 있나?
	fmt.Println("ManagedFiles called")
	err := fs.Rename("oldpath", "newpath")
	if err != nil {
		fmt.Errorf("error while calling rename: %w", err)
	}
	err = fs.Remove("file")
	if err != nil {
		fmt.Errorf("error while calling remove: %w", err)
	}
}
