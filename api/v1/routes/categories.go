package routes

import (
	"github.com/gin-gonic/gin"
	"salawat/api/api/v1/controllers"
)

func Categories(r *gin.Engine) {
	categoryRouter := r.Group("/api/categories")
	{
		categoryRouter.GET("/", controllers.CategoryController{}.CategoryIndex)
		categoryRouter.POST("/", controllers.CategoryController{}.CreateCategory)
		categoryRouter.GET("/:category_id", controllers.CategoryController{}.CategorySingle)
	}
}
