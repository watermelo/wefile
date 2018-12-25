package main

import (
	"fmt"
	"io"
	"strings"

	"io/ioutil"
	"os"
	"os/exec"

	. "wefile/log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/tran/:filetype", func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			Error.Println(err)
			return
		}
		defer file.Close()

		tempfile, err := os.Create(handler.Filename)

		io.Copy(tempfile, file)

		filetype := c.Param("filetype")
		cmd := exec.Command("unoconv", "-f", filetype, tempfile.Name())
		out, err := cmd.CombinedOutput()
		if err != nil {
			Error.Println(err)
		}
		defer os.Remove(tempfile.Name())

		outfileName := strings.Split(tempfile.Name(), ".")[0] + "." + filetype

		defer os.Remove(outfileName)

		dat, err := ioutil.ReadFile(outfileName)

		c.Header("Content-Disposition", "attachment; filename=" + outfileName)
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(dat)))
		c.Writer.Write([]byte(dat))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
