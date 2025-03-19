package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/greeneg/update-reporterd/model"
)

// CreateArchitecture creaets a new architecture
//
//	@Summary		Create a new architecture
//	@Description	Create a new architecture
//	@Tags			architectures
//	@Accept			json
//	@Produce		json
//	@Param			architecture	body	model.Architecture	true	"Architecture data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/architecture [post]
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

// DeleteArchitecture deletes an architecture
//
//	@Summary		Delete architecture
//	@Description	Delete an architecture
//	@Tags			architectures
//	@Accept			json
//	@Produce		json
//	@Param			archId	path	int	true	"Architecture Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/architecture/{archId} [delete]
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

// GetArchitectureById gets an architecture by Id
//
//	@Summary		Get architecture by Id
//	@Description	Get architecture by Id
//	@Tags			architectures
//	@Accept			json
//	@Produce		json
//	@Param			archId	path	int	true	"Architecture Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Architecture
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/architecture/id/{archId} [get]
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

// GetArchitectureByName gets an architecture by name
//
//	@Summary		Get architecture by name
//	@Description	Get architecture by name
//	@Tags			architectures
//	@Accept			json
//	@Produce		json
//	@Param			archName	path	string	true	"Architecture Name"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Architecture
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/architecture/name/{archName} [get]
func (u *UpdateReporter) GetArchitectureByName(c *gin.Context) {
	archName := c.Param("archName")

	arch, err := model.GetArchitectureByName(archName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, arch)
}

// GetArchitectures gets all architectures
//
//	@Summary		Get all architectures
//	@Description	Get all architectures
//	@Tags			architectures
//	@Accept			json
//	@Produce		json
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Architecture
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/architectures [get]
func (u *UpdateReporter) GetArchitectures(c *gin.Context) {
	archs, err := model.GetArchitectures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, archs)
}
