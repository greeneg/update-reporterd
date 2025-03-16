package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

func (u *UpdateReporter) CreateOSFamily(c *gin.Context) {
	var osf model.OSFamily
	if err := c.ShouldBindJSON(&osf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := model.CreateOSFamily(osf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OS family"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OS family created"})
}

func (u *UpdateReporter) DeleteOSFamily(c *gin.Context) {
	osfId, err := strconv.Atoi(c.Param("osfId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OS family Id"})
		return
	}

	ok, err := model.DeleteOSFamily(osfId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete OS family"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OS family deleted"})
}

func (u *UpdateReporter) GetOSFamilyById(c *gin.Context) {
	osfId, err := strconv.Atoi(c.Param("osfId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OS family Id"})
		return
	}

	osf, err := model.GetOSFamilyById(osfId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osf)
}

func (u *UpdateReporter) GetOSFamilyByName(c *gin.Context) {
	osfName := c.Param("osfName")

	osf, err := model.GetOSFamilyByName(osfName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osf)
}

func (u *UpdateReporter) GetOSFamilies(c *gin.Context) {
	osf, err := model.GetOSFamilies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osf)
}
