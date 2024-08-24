package handlers

import "github.com/gin-gonic/gin"

type Orders interface {
	GetByID(c *gin.Context)
}
