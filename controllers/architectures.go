package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/greeneg/update-reporterd/model"
)

func (u *UpdateReporter) CreateArchitecture(c *gin.Context) {
	var a model.Architecture
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := model.CreateArchitecture(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create architecture"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Architecture created"})
}

func (u *UpdateReporter) DeleteArchitecture(c *gin.Context) {
	archId, err := strconv.Atoi(c.Param("archId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid architecture Id"})
		return
	}

	ok, err := model.DeleteArchitecture(archId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete architecture"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Architecture deleted"})
}

func (u *UpdateReporter) GetArchitectureById(c *gin.Context) {
	archId, err := strconv.Atoi(c.Param("archId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid architecture Id"})
		return
	}

	arch, err := model.GetArchitectureById(archId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, arch)
}

func (u *UpdateReporter) GetArchitectureByName(c *gin.Context) {
	archName := c.Param("archName")

	arch, err := model.GetArchitectureByName(archName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, arch)
}

func (u *UpdateReporter) GetArchitectures(c *gin.Context) {
	archs, err := model.GetArchitectures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, archs)
}
