package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeksilaen/api-builder/db"
	middlewares "github.com/jeksilaen/api-builder/middlewares"
	userHandler "github.com/jeksilaen/api-builder/modules/user/handlers"
	collectionHandler "github.com/jeksilaen/api-builder/modules/collection/handlers"
	requestHandler "github.com/jeksilaen/api-builder/modules/request/handlers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middlewares.SetJSONContentTypeMiddleware())

	userHandler.InitUserHttpHandler(router)
	collectionHandler.InitCollectionHttpHandler(router)
	requestHandler.InitRequestHttpHandler(router)

	router.Run("localhost:8080")
}
