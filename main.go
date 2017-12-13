package main

import (
	"flag"
	"fmt"
	"lask3802/simple-OTA/folder"
)

/*
  dir/{CI_COMMIT_SHA}/(some content...)
*/
func main() {
	var dir string
	flag.StringVar(&dir, "dir", "public", "Search Directory")
	f := folder.Recent(dir, 0, 20)
	for _, info := range f {
		fmt.Println(info.Name())
	}
}
