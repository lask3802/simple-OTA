package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"lask3802/simple-OTA/ota"
	"net/http"
	"sort"
)

/*
  dir/{CI_COMMIT_SHA}/(some content...)
*/

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "public", "Search Directory")

	//f := ota.Recent(dir, 0, 20)

	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.StaticFS("/public", http.Dir("public"))
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		blocks := ota.FindCommits("public/", 0, 10)
		sort.Sort(blocks)
		c.HTML(http.StatusOK, "tables.html", blocks)
	})
	r.RunTLS("127.0.0.1:443", "crt/server.crt", "crt/server.key")
}
