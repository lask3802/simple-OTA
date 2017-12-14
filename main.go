package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"lask3802/simple-OTA/ota"
	"net/http"
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
	r.GET("/", func(c *gin.Context) {
		blocks := ota.FindCommits("public/", 0, 10)
		c.HTML(http.StatusOK, "index.html", blocks)
	})
	r.Run(":8080")
}
