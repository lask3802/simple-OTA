package ota

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
		panic(err)
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

func FindAPK(dir string) (string, error) {
	return FindPattern(dir, "*.apk")
}

func FindIPA(dir string) (string, error) {
	return FindPattern(dir, "*.ipa")
}

func FindPattern(dir string, pattern string) (string, error) {
	path := filepath.Join(dir, pattern)
	file, err := filepath.Glob(path)
	if err != nil {
		return "", err
	}
	if len(file) > 0 {
		return file[0], nil
	}
	return "", os.ErrNotExist
}
