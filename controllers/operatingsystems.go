package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

// CreateOS creates a new operating system
//
//	@Summary		Create a new operating system
//	@Description	Create a new operating system
//	@Tags			operatingsystems
//	@Accept			json
//	@Produce		json
//	@Param			operatingsystem	body	model.OperatingSystem	true	"Operating System data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/os [post]
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

// DeleteOS deletes an operating system
//
//	@Summary		Delete operating system
//	@Description	Delete an operating system
//	@Tags			operatingsystems
//	@Accept			json
//	@Produce		json
//	@Param			osId	path	int	true	"Operating System Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/os/{osId} [delete]
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

// GetOSById gets an operating system by Id
//
//	@Summary		Get operating system by Id
//	@Description	Get operating system by Id
//	@Tags			operatingsystems
//	@Accept			json
//	@Produce		json
//	@Param			osId	path	int	true	"Operating System Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.OperatingSystem
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/os/id/{osId} [get]
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

// GetOSByName gets an operating system by name
//
//	@Summary		Get operating system by name
//	@Description	Get operating system by name
//	@Tags			operatingsystems
//	@Accept			json
//	@Produce		json
//	@Param			osName	path	string	true	"Operating System Name"
//	@Security		BasicAuth
//	@Success		200	{object}	model.OperatingSystem
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/os/name/{osName} [get]
func (u *UpdateReporter) GetOSByName(c *gin.Context) {
	osName := c.Param("osName")

	osStruct, err := model.GetOperatingSystemByName(osName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osStruct)
}

// GetOSes gets all operating systems
//
//	@Summary		Get all operating systems
//	@Description	Get all operating systems
//	@Tags			operatingsystems
//	@Accept			json
//	@Produce		json
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.OperatingSystem
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/oses [get]
func (u *UpdateReporter) GetOSes(c *gin.Context) {
	osStruct, err := model.GetOperatingSystems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osStruct)
}
