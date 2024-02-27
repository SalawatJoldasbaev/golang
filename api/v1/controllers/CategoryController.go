package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"salawat/api/database"
	"salawat/api/models"
	"salawat/api/requests"
)

type CategoryController struct {
	BaseController
}

func (con CategoryController) CategoryIndex(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (con CategoryController) CategorySingle(c *gin.Context) {
	var category models.Category
	if err := database.DB.Where("id = ?", c.Param("category_id")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func (con CategoryController) CreateCategory(c *gin.Context) {
	var input requests.CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{Title: input.Title}
	database.DB.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}
