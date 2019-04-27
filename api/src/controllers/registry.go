package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/registry/api/src/dao"
	"github.com/juankis/registry/api/src/models"
	"github.com/juankis/registry/api/src/utils"
)

func InsertRegistry(c *gin.Context) {
	var reg models.Response
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	registry := reg.Data
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
