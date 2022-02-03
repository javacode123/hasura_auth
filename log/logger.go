package log

import (
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// InitLogger TODO 日志分割
func InitLogger() {
	stdWriter := os.Stdout
	logPath := "log_files"
	err := os.Mkdir(logPath, 0755)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			panic(err)
		}
	}
	fileWriter, err := os.OpenFile(logPath+"/log_"+time.Now().Format("2006-01-02 15:04:05")+".txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(io.MultiWriter(stdWriter, fileWriter))
	logrus.SetReportCaller(true)
}
