package filesystem

import "testing"

//todo : 이걸 정확하게 어떻게 작성을 해야 하나?

func Test(t *testing.T) {
	var fileSystem FileSystem = OSFileSystem{}
	ManageFiles(fileSystem)
}
