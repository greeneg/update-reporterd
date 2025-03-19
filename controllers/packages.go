package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

// GetPackages returns all packages
//
//	@Summary		Get all packages
//	@Description	Get all packages
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs [get]
func (u *UpdateReporter) GetPackages(c *gin.Context) {
	pkgs, err := model.GetPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByArch returns all packages by architecture
//
//	@Summary		Get all packages by architecture
//	@Description	Get all packages by architecture
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgArch	path	string	true	"Package Architecture"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/arch/{pkgArch} [get]
func (u *UpdateReporter) GetPkgsByArch(c *gin.Context) {
	pkgs, err := model.GetPkgsByArch(c.Param("pkgArch"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByArchId returns all packages by architecture Id
//
//	@Summary		Get all packages by architecture Id
//	@Description	Get all packages by architecture Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgArchId	path	int	true	"Package Architecture Id"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/archid/{pkgArchId} [get]
func (u *UpdateReporter) GetPkgsByArchId(c *gin.Context) {
	archId, err := strconv.Atoi(c.Param("pkgArchId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgs, err := model.GetPkgsByArchId(archId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByPlatform returns all packages by platform
//
//	@Summary		Get all packages by platform
//	@Description	Get all packages by platform
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgPlatform	path	string	true	"Package Platform"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/platform/{pkgPlatform} [get]
func (u *UpdateReporter) GetPkgsByPlatform(c *gin.Context) {
	pkgs, err := model.GetPkgsByPlatform(c.Param("pkgPlatform"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByPlatformId returns all packages by platform Id
//
//	@Summary		Get all packages by platform Id
//	@Description	Get all packages by platform Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgPlatformId	path	int	true	"Package Platform Id"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/platformid/{pkgPlatformId} [get]
func (u *UpdateReporter) GetPkgsByPlatformId(c *gin.Context) {
	platId, err := strconv.Atoi(c.Param("pkgPlatformId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgs, err := model.GetPkgsByPlatformId(platId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByType returns all packages by type
//
//	@Summary		Get all packages by type
//	@Description	Get all packages by type
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgType	path	string	true	"Package Type"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/type/{pkgType} [get]
func (u *UpdateReporter) GetPkgsByType(c *gin.Context) {
	pkgs, err := model.GetPkgsByType(c.Param("pkgType"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgsByTypeId returns all packages by type Id
//
//	@Summary		Get all packages by type Id
//	@Description	Get all packages by type Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgTypeId	path	int	true	"Package Type Id"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.Package
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkgs/typeid/{pkgTypeId} [get]
func (u *UpdateReporter) GetPkgsByTypeId(c *gin.Context) {
	typeId, err := strconv.Atoi(c.Param("pkgTypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgs, err := model.GetPkgsByTypeId(typeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

// GetPkgArch returns a package architecture
//
//	@Summary		Get package architecture
//	@Description	Get package architecture
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgArch	path	string	true	"Package Architecture"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Architecture
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/arch/{pkgArch} [get]
func (u *UpdateReporter) GetPkgArch(c *gin.Context) {
	pkgArch, err := model.GetArchitectureByName(c.Param("pkgArch"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgArch)
}

// GetPkgArchById returns a package architecture by Id
//
//	@Summary		Get package architecture by Id
//	@Description	Get package architecture by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgArchId	path	int	true	"Package Architecture Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Architecture
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/archid/{pkgArchId} [get]
func (u *UpdateReporter) GetPkgArchById(c *gin.Context) {
	pkgArchId, err := strconv.Atoi(c.Param("pkgArchId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgArch, err := model.GetArchitectureById(pkgArchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgArch)
}

// GetPackageById returns a package by Id
//
//	@Summary		Get package by Id
//	@Description	Get package by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgId	path	int	true	"Package Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Package
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/id/{pkgId} [get]
func (u *UpdateReporter) GetPackageById(c *gin.Context) {
	pkgId, err := strconv.Atoi(c.Param("pkgId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkg, err := model.GetPackageById(pkgId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// GetVersions returns all versions of a package
//
//	@Summary		Get all versions of a package
//	@Description	Get all versions of a package
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgId	path	int	true	"Package Id"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.PackageVersion
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/id/{pkgId}/versions [get]
func (u *UpdateReporter) GetVersions(c *gin.Context) {
	versions, err := model.GetPkgVersionsById(c.Param("pkgId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, versions)
}

// GetPackage returns a package by name
//
//	@Summary		Get package by name
//	@Description	Get package by name
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgName	path	string	true	"Package Name"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Package
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/name/{pkgName} [get]
func (u *UpdateReporter) GetPackage(c *gin.Context) {
	pkg, err := model.GetPackage(c.Param("pkgName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// GetVersionsByName returns all versions of a package by name
//
//	@Summary		Get all versions of a package by name
//	@Description	Get all versions of a package by name
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgName	path	string	true	"Package Name"
//	@Security		BasicAuth
//	@Success		200	{object}	[]model.PackageVersion
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/name/{pkgName}/versions [get]
func (u *UpdateReporter) GetVersionsByName(c *gin.Context) {
	versions, err := model.GetVersionsByName(c.Param("pkgName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, versions)
}

// GetPkgPlatform returns a package platform
//
//	@Summary		Get package platform
//	@Description	Get package platform
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgPlatform	path	string	true	"Package Platform"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Platform
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/platform/{pkgPlatform} [get]
func (u *UpdateReporter) GetPkgPlatform(c *gin.Context) {
	pkgPlatform, err := model.GetPkgPlatform(c.Param("pkgPlatform"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgPlatform)
}

// GetPkgPlatformById returns a package platform by Id
//
//	@Summary		Get package platform by Id
//	@Description	Get package platform by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgPlatformId	path	int	true	"Package Platform Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Platform
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/platformid/{pkgPlatformId} [get]
func (u *UpdateReporter) GetPkgPlatformById(c *gin.Context) {
	pkgPlatformId, err := strconv.Atoi(c.Param("pkgPlatformId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgPlatform, err := model.GetPkgPlatformById(pkgPlatformId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgPlatform)
}

// GetPkgType returns a package type
//
//	@Summary		Get package type
//	@Description	Get package type
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgType	path	string	true	"Package Type"
//	@Security		BasicAuth
//	@Success		200	{object}	model.PackageType
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/type/{pkgType} [get]
func (u *UpdateReporter) GetPkgType(c *gin.Context) {
	pkgType, err := model.GetPkgTypeByName(c.Param("pkgType"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgType)
}

// GetPkgTypeById returns a package type by Id
//
//	@Summary		Get package type by Id
//	@Description	Get package type by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgTypeId	path	int	true	"Package Type Id"
//	@Security		BasicAuth
//	@Success		200	{object}	model.PackageType
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/typeid/{pkgTypeId} [get]
func (u *UpdateReporter) GetPkgTypeById(c *gin.Context) {
	pkgTypeId, err := strconv.Atoi(c.Param("pkgTypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	pkgType, err := model.GetPkgTypeById(pkgTypeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgType)
}

// CreatePackage creates a new package
//
//	@Summary		Create new package
//	@Description	Create new package
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkg	body	model.Package	true	"Package object"
//	@Security		BasicAuth
//	@Success		201	{object}	model.Package
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg [post]
func (u *UpdateReporter) CreatePackage(c *gin.Context) {
	var pkg model.Package
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	if err := model.CreatePackage(&pkg); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pkg)
}

// CreatePkgType creates a new package type
//
//	@Summary		Create new package type
//	@Description	Create new package type
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgType	body	model.PackageType	true	"Package Type object"
//	@Security		BasicAuth
//	@Success		201	{object}	model.PackageType
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/type/{pkgType} [post]
func (u *UpdateReporter) CreatePkgType(c *gin.Context) {
	var pkgType model.PackageType
	if err := c.ShouldBindJSON(&pkgType); err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	if err := model.CreatePkgType(&pkgType); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pkgType)
}

// UpdatePkg updates a package
//
//	@Summary		Update a package
//	@Description	Update a package
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgName	path	string	true	"Package Name"
//	@Security		BasicAuth
//	@Success		200	{object}	model.Package
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/name/{pkgName} [patch]
func (u *UpdateReporter) UpdatePkg(c *gin.Context) {
	var pkg model.Package
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	if err := model.UpdatePkg(&pkg); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// DeletePackage deletes a package by Id
//
//	@Summary		Delete a package by Id
//	@Description	Delete a package by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgId	path	int	true	"Package Id"
//	@Security		BasicAuth
//	@Success		204	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/id/{pkgId} [delete]
func (u *UpdateReporter) DeletePackage(c *gin.Context) {
	pkgId, err := strconv.Atoi(c.Param("pkgId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	if err := model.DeletePkg(pkgId); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeletePackageByName deletes a package by name
//
//	@Summary		Delete a package by name
//	@Description	Delete a package by name
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgName	path	string	true	"Package Name"
//	@Security		BasicAuth
//	@Success		204	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/name/{pkgName} [delete]
func (u *UpdateReporter) DeletePackageByName(c *gin.Context) {
	if err := model.DeletePkgByName(c.Param("pkgName")); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// DeletePkgType deletes a package type by Id
//
//	@Summary		Delete a package type by Id
//	@Description	Delete a package type by Id
//	@Tags			packages
//	@Accept			json
//	@Produce		json
//	@Param			pkgType	path	int	true	"Package Type Id"
//	@Security		BasicAuth
//	@Success		204	{object}	model.SuccessMsg
//	@Failure		400	{object}	model.FailureMsg
//	@Failure		500	{object}	model.FailureMsg
//	@Router			/pkg/type/{pkgType} [delete]
func (u *UpdateReporter) DeletePkgType(c *gin.Context) {
	typeId, err := strconv.Atoi(c.Param("pkgType"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureMsg{Error: err.Error()})
		return
	}
	if err := model.DeletePkgType(typeId); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
