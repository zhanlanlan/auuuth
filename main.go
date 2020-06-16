package main

import (
	"auuuth/routers"

	_ "auuuth/models"
	_ "auuuth/utils/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	// get the server instance
	s := routers.NewServer()

	// use recover middleware
	s.Use(gin.Recovery())

	// // use log format
	// s.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf(
	// 		"%s - [%s] \"%s %s %s %d %s \" \" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.ErrorMessage,
	// 	)
	// }))

	// use default log middleware
	// todo use global logger rather than gin default logger
	s.Use(gin.Logger())

	// run and linsten
	s.Run(":8080")
}
