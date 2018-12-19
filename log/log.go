package log

import (
	"log"
	"os"
	"io"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)

func init(){
	errFile, err := os.OpenFile("./errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	Info = log.New(os.Stdout, "[COLUMBUS] [INFO] ", log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout, "[COLUMBUS] [WARNING] ", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr,errFile), "[COLUMBUS] [ERROR] ", log.Ldate | log.Ltime | log.Lshortfile)

}
