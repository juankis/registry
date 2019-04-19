package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/registry/api/src/dao"
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
	err = dao.InsertRegistry(&registry)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &registry)
	return
}

func GetRegistry(c *gin.Context) {
	var checks []models.Registry
	checks, err := dao.GetRegistryAll()
	if err != nil {
		utils.CustomResponse(c, "getting checks", err, 404)
		return
	}
	c.JSON(200, checks)
	return
}
