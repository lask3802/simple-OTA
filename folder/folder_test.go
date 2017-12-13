package folder

import (
	"os"
	"testing"
	"time"
)

type FileInfoMock struct {
	name    string
	modTime time.Time
}

func (f FileInfoMock) Name() string {
	return f.name
}

func (FileInfoMock) Size() int64 {
	panic("implement me")
}

func (FileInfoMock) Mode() os.FileMode {
	panic("implement me")
}

func (f FileInfoMock) ModTime() time.Time {
	return f.modTime
}

func (FileInfoMock) IsDir() bool {
	return true
}

func (FileInfoMock) Sys() interface{} {
	panic("implement me")
}

func TestRecent(t *testing.T) {
	files := make(FileInfoSlice, 3)
	files[0] = FileInfoMock{
		name:    "file 1",
		modTime: time.Unix(1, 0),
	}
	files[1] = FileInfoMock{
		name:    "file 0",
		modTime: time.Unix(0, 0),
	}
	files[2] = FileInfoMock{
		name:    "file 2",
		modTime: time.Unix(2, 0),
	}
	result := RecentInternal(files, 0, 3)
	if result[0].Name() != "file 0" || result[1].Name() != "file 1" || result[2].Name() != "file 2" {
		t.Error("file name not match")
	}
	result = RecentInternal(files, 0, 2)
	if len(result) != 2 {
		t.Error("Wrong file info count")
	}

	result = RecentInternal(files, 0, 5)
	if len(result) != 3 {
		t.Error("Wrong file info count")
	}
}
