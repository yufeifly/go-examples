package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	fileName := "/home/alan/gin.log"
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.Info("fuck you")
}
