package middleware

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"guoke-helper-golang/config"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LogConf.LogfilePath
	logFileName := config.LogConf.LogFileName
	fileName := path.Join(logFilePath, logFileName)

	//src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err", err)
	//}

	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 设置最大保存时间(30天)
		rotatelogs.WithMaxAge(30*24*time.Hour),

		// 生成软链，指向最新日志文件
		// rotatelogs.WithLinkName(fileName),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
		)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.WithFields(logrus.Fields{
			"status":	statusCode,
			"latency":	latency,
			"client":	clientIP,
			"method":	reqMethod,
			"uri":		reqUri,
		}).Info()
	}
}
