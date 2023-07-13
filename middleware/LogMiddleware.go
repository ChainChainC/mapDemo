//write Logger middleware for golang

package middleware

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Enum for log level
type LogLevel int

const (
	Debug LogLevel = iota
	Info  LogLevel = iota
	Warn  LogLevel = iota
	Error LogLevel = iota
	Fatal LogLevel = iota
)

func Logger() *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return log.WithField("file", filename).WithField("function", fn)
}

func LoggerMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
