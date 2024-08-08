package router

import (
	"github.com/Hailemari/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.DELETE("/tasks/:id", controllers.RemoveTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.POST("/tasks", controllers.AddTask)

	return r
}
