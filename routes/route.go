package routes

import (
	sql "goTest/datatable"
	"goTest/middleware"

	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	server := gin.Default()
	server.Static("/css", "./templates/asset/css")
	server.LoadHTMLGlob("templates/html/*/*.html")
	server.Use(middleware.GenerateLogFile())

	apiRoute := server.Group("/api")

	mongo := sql.MongodbDefault()
	//mydb := sql.MysqlDefault()

	LoginRoutes(apiRoute, mongo)
	CoffeeAPIRoutes(apiRoute, mongo)

	err := server.Run(":8080")

	if err != nil {
		panic("start server failed, because:" + err.Error())
	}
}
