package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "wefile/log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/tran/:filetype", func(c *gin.Context) {
		// 读取文件
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// 建临时目录
		tempfile, err := ioutil.TempFile(os.TempDir(), "wefile")
		if err != nil {
			fmt.Println(err)
			return
		}

		io.Copy(tempfile, file)
		tempfile.Close()

		filename := tempfile.Name() + filepath.Ext(handler.Filename)
		filetype := c.Param("filetype")
		cmd := exec.Command("unoconv", "-f", filetype, filename)
		out, err := cmd.CombinedOutput()
		if err != nil {
			Error.Println(err)
		}
		fmt.Println(string(out))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
