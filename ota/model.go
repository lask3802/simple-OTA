package ota

import (
	"log"
	"path/filepath"
	"strings"
	"github.com/dustin/go-humanize"
	"encoding/json"
	"io/ioutil"
	"time"
	"os"
	"html/template"
)

type CommitBlock struct {
	IPALink     template.URL
	APKUrl      string
	Message     string
	ProjectName string
	Branch      string
	Icon        string
	Commit      string
	Time        string
	UnixTime    int64
}

type CommitBlocks []CommitBlock

func (c CommitBlocks) Len() int {
	return len(c)
}

func (c CommitBlocks) Less(i, j int) bool {
	t1, _:= time.Parse(time.UnixDate,c[i].Time)
	t2, _:= time.Parse(time.UnixDate,c[j].Time)
	return  t1.Unix() > t2.Unix()
}

func (c CommitBlocks) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func FindCommits(targetDir string, start int, count int) CommitBlocks {
	fullPath, err := filepath.Abs(targetDir)

	if err != nil {
		log.Fatal(err)
	}

	folders := Recent(fullPath, start, count)
	result := make([]CommitBlock, len(folders))
	available := 0
	for _, d := range folders {
		env, err := Env(targetDir, d)
		if err != nil{
			log.Println(err)
			continue
		}

		apk, err := FindAPK(filepath.Join(targetDir, d.Name()))
		if err != nil {
			log.Println(err)
		}
		ipa, err := FindIPA(filepath.Join(targetDir, d.Name()))
		if err != nil {
			log.Println(err)
		}

		time, err := time.Parse(time.UnixDate, env["BUILD_TIME"])
		if err != nil {
			log.Println(err)
		}
		ipaLink := strings.Replace(ipa, "\\", "/", -1)
		result[available] = CommitBlock{
			IPALink:     template.URL(ipaLink),
			APKUrl:      strings.Replace(apk, "\\", "/", -1),
			Commit:      env["CI_COMMIT_SHA"],
			Time:        humanize.Time(time),
			Branch:      env["CI_COMMIT_REF_NAME"],
			ProjectName: env["CI_PROJECT_PATH"],
			Message:     env["BUILD_MESSAGE"],
			UnixTime: time.Unix(),
		}
		available++
	}
	return result[:available]
}
func Env(targetDir string, d os.FileInfo) (map[string]string, error) {
	cienv, err := ioutil.ReadFile(filepath.Join(targetDir, d.Name(), "ci.json"))
	if err != nil {
		log.Print(err)
	}
	var env map[string]string
	err = json.Unmarshal(cienv, &env)
	if err != nil {
		log.Println(err)
	}
	return env,err
}
