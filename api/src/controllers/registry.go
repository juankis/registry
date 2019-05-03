package controllers

import (
	"strconv"

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

//GetRegistryAll get all registries
func GetRegistryAll(c *gin.Context) {
	var regestries []models.Registry
	regestries, err := dao.GetRegistryAll()
	if err != nil {
		utils.CustomResponse(c, "getting regestries", err, 404)
		return
	}
	c.JSON(200, regestries)
	return
}

// DeleteRegistry delete some job in the db
func DeleteRegistry(c *gin.Context) {
	var registry models.Registry
	RegistryID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	registry.ID = RegistryID
	err = dao.DeleteRegistry(&registry)
	if err != nil {
		utils.CustomResponse(c, "deleting registry", err, 400)
		return
	}
	c.Status(200)
	return
}

//PutRegistry actualizar registro
func PutRegistry(c *gin.Context) {
	var reg models.Response
	RegistryID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	err = c.ShouldBindJSON(&reg)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	registry := reg.Data
	registry.ID = RegistryID
	err = dao.UpdateRegistry(&registry)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &registry)
	return
}

//GetRegistry obtener registro
func GetRegistry(c *gin.Context) {
	var registry models.Registry
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	registry.ID = id
	err = dao.GetRegistry(&registry)
	if err != nil {
		utils.CustomResponse(c, "getting registry", err, 404)
		return
	}
	c.JSON(200, models.Response{Data: registry})
	return
}
