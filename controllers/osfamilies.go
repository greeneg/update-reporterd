package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

// CreateOSFamily creates a new operating system family
//
//	@Summary		Create a new operating system family
//	@Description	Create a new operating system family
//	@Tags			osfamilies
//	@Accept			json
//	@Produce		json
//	@Param			osfamily	body	model.OSFamily	true	"Operating System Family data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/osfamily [post]
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

// DeleteOSFamily deletes an operating system family
//
//	@Summary		Delete operating system family
//	@Description	Delete an operating system family
//	@Tags			osfamilies
//	@Accept			json
//	@Produce		json
//	@Param			osfId	path	int	true	"Operating System Family Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/osfamily/{osfId} [delete]
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

// GetOSFamilyById gets an operating system family by Id
//
//	@Summary		Get operating system family by Id
//	@Description	Get operating system family by Id
//	@Tags			osfamilies
//	@Accept			json
//	@Produce		json
//	@Param			osfId	path	int	true	"Operating System Family Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.OSFamily
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/osfamily/id/{osfId} [get]
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

// GetOSFamilyByName gets an operating system family by name
//
//	@Summary		Get operating system family by name
//	@Description	Get operating system family by name
//	@Tags			osfamilies
//	@Accept			json
//	@Produce		json
//	@Param			osfName	path	string	true	"Operating System Family Name"
//	@Security		BasicAuth
//	@Success		200	{object}	model.OSFamily
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/osfamily/name/{osfName} [get]
func (u *UpdateReporter) GetOSFamilyByName(c *gin.Context) {
	osfName := c.Param("osfName")

	osf, err := model.GetOSFamilyByName(osfName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osf)
}

// GetOSFamilies gets all operating system families
//
//	@Summary		Get all operating system families
//	@Description	Get all operating system families
//	@Tags			osfamilies
//	@Accept			json
//	@Produce		json
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.OSFamily
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/osfamilies [get]
func (u *UpdateReporter) GetOSFamilies(c *gin.Context) {
	osf, err := model.GetOSFamilies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, osf)
}
