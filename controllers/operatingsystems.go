package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

func (u *UpdateReporter) CreateOS(c *gin.Context) {
	var osf model.OperatingSystem
	if err := c.ShouldBindJSON(&osf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := model.CreateOperatingSystem(osf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OS"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OS created"})
}

func (u *UpdateReporter) DeleteOS(c *gin.Context) {
	osfId, err := strconv.Atoi(c.Param("osId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OS Id"})
		return
	}

	ok, err := model.DeleteOperatingSystem(osfId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete OS"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OS deleted"})
}

func (u *UpdateReporter) GetOSById(c *gin.Context) {
	osId, err := strconv.Atoi(c.Param("osId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OS Id"})
		return
	}

	osStruct, err := model.GetOperatingSystemById(osId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osStruct)
}

func (u *UpdateReporter) GetOSByName(c *gin.Context) {
	osName := c.Param("osName")

	osStruct, err := model.GetOperatingSystemByName(osName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osStruct)
}

func (u *UpdateReporter) GetOSes(c *gin.Context) {
	osStruct, err := model.GetOperatingSystems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osStruct)
}
