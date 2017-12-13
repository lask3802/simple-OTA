package folder

import (
	"io/ioutil"
	"os"
	"sort"
)

type FileInfoSlice []os.FileInfo

func (s FileInfoSlice) Less(i, j int) bool {
	return s[i].ModTime().Unix() < s[j].ModTime().Unix()
}

func (s FileInfoSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s FileInfoSlice) Len() int { return len(s) }

func Recent(targetDir string, start int, count int) []os.FileInfo {
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		panic("No public folder")
	}
	return RecentInternal(files, start, count)
}

func RecentInternal(slice []os.FileInfo, start int, count int) []os.FileInfo {
	commitFolders := make(FileInfoSlice, len(slice))

	for idx, f := range slice {
		if f.IsDir() {
			commitFolders[idx] = f
		}
	}
	sort.Sort(commitFolders)
	var end int
	if end = start + count; end > len(commitFolders) {
		end = len(commitFolders)
	}
	return commitFolders[start:end]
}
