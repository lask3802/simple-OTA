package ota

import (
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CommitBlock struct {
	IPALink     string
	APKUrl      string
	Message     string
	ProjectName string
	Branch      string
	Icon        string
	Commit      string
	Time        time.Time
}

func FindCommits(targetDir string, start int, count int) []CommitBlock {
	fullPath, err := filepath.Abs(targetDir)

	if err != nil {
		log.Fatal(err)
	}

	folders := Recent(fullPath, start, count)
	result := make([]CommitBlock, len(folders))
	for idx, d := range folders {
		apk, err := FindAPK(filepath.Join(targetDir, d.Name()))
		if err != nil {
			log.Println(err)
		}
		ipa, err := FindIPA(filepath.Join(targetDir, d.Name()))
		if err != nil {
			log.Println(err)
		}
		result[idx] = CommitBlock{
			IPALink: strings.Replace(ipa, "\\", "/", -1),
			APKUrl:  strings.Replace(apk, "\\", "/", -1),
			Commit:  d.Name(),
			Time:    d.ModTime(),
		}
	}
	return result
}
