package controllers

import "github.com/gin-gonic/gin"

type TaskController struct {
	BaseController
}

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

func (con TaskController) TaskIndex(c *gin.Context) {
	c.String(200, "api TaskIndex")
}

func (con TaskController) TaskSingle(c *gin.Context) {
	c.String(200, "api Task single")
}
