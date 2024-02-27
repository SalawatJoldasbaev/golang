package routes

import (
	"github.com/gin-gonic/gin"
	"salawat/api/api/v1/controllers"
)

func Tasks(r *gin.Engine) {
	tasksRouter := r.Group("/api/tasks")
	{
		tasksRouter.GET("/", controllers.TaskController{}.TaskIndex)
		tasksRouter.GET("/:task_id", controllers.TaskController{}.TaskSingle)
	}
}
