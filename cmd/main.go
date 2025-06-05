package main

import (
	"acronis/controller"
	"acronis/repository"
	"acronis/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	repo := repository.NewStoreDataMap()

	go repo.StartCleanupJob()

	svc := service.NewInMemoryStore(repo)
	controller := controller.NewController(svc)

	api := server.Group("/api")
	{
		api.GET("/data", controller.GetAllData)
		api.GET("/data/:key", controller.GetDataByKey)
		api.POST("/data", controller.SetData)
		api.PUT("/data/:key", controller.UpdateData)
		api.DELETE("/data/:key", controller.DeleteData)
		api.PUT("/data/:key/push", controller.PushToList)
		api.PUT("/data/:key/pop", controller.PopFromList)
	}

	server.Run(":8081")
}
