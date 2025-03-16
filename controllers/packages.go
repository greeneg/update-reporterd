package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greeneg/update-reporterd/model"
)

func (u *UpdateReporter) GetPackages(c *gin.Context) {
	pkgs, err := model.GetPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

func (u *UpdateReporter) GetPkgsByArch(c *gin.Context) {
	pkgs, err := model.GetPkgsByArch(c.Param("pkgArch"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

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

func (u *UpdateReporter) GetPkgsByPlatform(c *gin.Context) {
	pkgs, err := model.GetPkgsByPlatform(c.Param("pkgPlatform"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

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

func (u *UpdateReporter) GetPkgsByType(c *gin.Context) {
	pkgs, err := model.GetPkgsByType(c.Param("pkgType"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgs)
}

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

func (u *UpdateReporter) GetPkgArch(c *gin.Context) {
	pkgArch, err := model.GetArchitectureByName(c.Param("pkgArch"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgArch)
}

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

func (u *UpdateReporter) GetVersions(c *gin.Context) {
	versions, err := model.GetPkgVersionsById(c.Param("pkgId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, versions)
}

func (u *UpdateReporter) GetPackage(c *gin.Context) {
	pkg, err := model.GetPackage(c.Param("pkgName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

func (u *UpdateReporter) GetVersionsByName(c *gin.Context) {
	versions, err := model.GetVersionsByName(c.Param("pkgName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, versions)
}

func (u *UpdateReporter) GetPkgPlatform(c *gin.Context) {
	pkgPlatform, err := model.GetPkgPlatform(c.Param("pkgPlatform"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgPlatform)
}

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

func (u *UpdateReporter) GetPkgType(c *gin.Context) {
	pkgType, err := model.GetPkgTypeByName(c.Param("pkgType"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkgType)
}

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

func (u *UpdateReporter) DeletePackageByName(c *gin.Context) {
	if err := model.DeletePkgByName(c.Param("pkgName")); err != nil {
		c.JSON(http.StatusInternalServerError, model.FailureMsg{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

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
