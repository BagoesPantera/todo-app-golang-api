package main

import (
	"github.com/gin-gonic/gin"
	todoController "restAPI/controllers"
	"restAPI/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/todo", todoController.FindAll)
	r.GET("/todo/:id", todoController.FindById)
	r.POST("/todo/add", todoController.Create)
	r.PUT("/todo/update/:id", todoController.Update)
	r.DELETE("/todo/delete/:id", todoController.Delete)

	r.Run()

}
