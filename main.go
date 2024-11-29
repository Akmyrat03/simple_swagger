package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joefazee/api-doc/docs"
	"github.com/joefazee/api-doc/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Simple API
// @version 1
// @Description Sample description
// @host localhost:8080
// @BasedPath /api/v1
func main() {

	r := gin.Default()

	user := r.Group("/users")
	{
		user.POST("/", handlers.CreateUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}

}
