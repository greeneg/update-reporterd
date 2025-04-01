package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

// CreateSystem creates a new system
//
//	@Summary		Create a new system
//	@Description	Create a new system
//	@Tags			systems
//	@Accept			json
//	@Produce		json
//	@Param			system	body	model.System	true	"System data"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/system [post]
func (u *UpdateReporter) CreateSystem(c *gin.Context) {
	var sys model.ProposedSystem
	if err := c.ShouldBindJSON(&sys); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := model.CreateSystem(sys)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: Systems.FQDN" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "System already exists"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create system"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "System created"})
	}
}

// DeleteSystem deletes a system
//
//	@Summary		Delete system
//	@Description	Delete a system
//	@Tags			systems
//	@Accept			json
//	@Produce		json
//	@Param			sysId	path	int	true	"System Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Router			/system/{sysId} [delete]
func (u *UpdateReporter) DeleteSystem(c *gin.Context) {
	sysId, err := strconv.Atoi(c.Param("sysId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system Id"})
		return
	}

	ok, err := model.DeleteSystem(sysId)
	if err != nil {
		if err.Error() == "Unknown system!" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "System not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "System deleted"})
}

func (u *UpdateReporter) GetSystems(c *gin.Context) {
	systems, err := model.GetSystems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, systems)
}

func (u *UpdateReporter) GetSystemById(c *gin.Context) {
	sysId, err := strconv.Atoi(c.Param("sysId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system Id"})
		return
	}

	sys, err := model.GetSystemById(sysId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sys)
}

func (u *UpdateReporter) GetSystemByName(c *gin.Context) {
	sysName := c.Param("sysName")

	sys, err := model.GetSystemByName(sysName)
	if err != nil {
		if err.Error() == "Unknown system!" {
			c.JSON(http.StatusNotFound, gin.H{"error": "System not found"})
			return
		} else {
			// If it's a different error, return a 500
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, sys)
}

func (u *UpdateReporter) GetSystemsByOsFamilyId(c *gin.Context) {
	osfId, err := strconv.Atoi(c.Param("osfId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OS family Id"})
		return
	}

	systems, err := model.GetSystemsByOsFamilyId(osfId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, systems)
}

func (u *UpdateReporter) GetSystemsByArchId(c *gin.Context) {
	archId, err := strconv.Atoi(c.Param("archId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid architecture Id"})
		return
	}

	systems, err := model.GetSystemsByArchId(archId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, systems)
}
