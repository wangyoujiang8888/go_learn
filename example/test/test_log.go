package main

import (
	mylog "example/util/log"
	"go.uber.org/zap"
)

func main()  {
	TestLogWrite()
}

func TestLogWrite()  {
	url := "http://baidu.com"
	mylog.Logger.LogInfo("log_info",zap.String("url", url))
	mylog.Logger.LogError("test_error",zap.String("url", url))
}