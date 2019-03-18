package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/registry/api/src/models"
	"github.com/juankis/registry/api/src/utils"
)

func InsertRegistry(c *gin.Context) {
	var registry models.Registry
	err := c.ShouldBindJSON(&registry)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &registry)
	return
}
