package routes

/*

  Copyright 2024, YggdrasilSoft, LLC.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

import (
	"github.com/gin-gonic/gin"

	"github.com/greeneg/update-reporterd/controllers"
)

func PrivateRoutes(g *gin.RouterGroup, u *controllers.UpdateReporter) {
	// Architectures
	g.GET("/architectures", u.GetArchitectures)                    // get all architectures
	g.GET("/architecture/id/:archId", u.GetArchitectureById)       // get architecture by Id
	g.GET("/architecture/name/:archName", u.GetArchitectureByName) // get architecture by name
	g.POST("/architecture", u.CreateArchitecture)                  // create new architecture
	g.DELETE("/architecture/:archId", u.DeleteArchitecture)        // delete an architecture by Id
	// OS Families
	g.GET("/osfamilies", u.GetOSFamilies)                 // get all OS families
	g.GET("/osfamily/id/:osfId", u.GetOSFamilyById)       // get OS family by Id
	g.GET("/osfamily/name/:osfName", u.GetOSFamilyByName) // get OS family by name
	g.POST("/osfamily", u.CreateOSFamily)                 // create new OS family
	g.DELETE("/osfamily/:osfId", u.DeleteOSFamily)        // delete an OS family by Id
	// OSes
	g.GET("/oses", u.GetOSes)                // get all OSes
	g.GET("/os/id/:osId", u.GetOSById)       // get OS by Id
	g.GET("/os/name/:osName", u.GetOSByName) // get OS by name
	g.POST("/os", u.CreateOS)                // create new OS
	g.DELETE("/os/:osId", u.DeleteOS)        // delete an OS by Id
	// packages
	g.GET("/pkgs", u.GetPackages)                                   // get all packages
	g.GET("/pkgs/arch/:pkgArch", u.GetPkgsByArch)                   // get all packages by architecture
	g.GET("/pkgs/archid/:pkgArchId", u.GetPkgsByArchId)             // get all packages by architecture Id
	g.GET("/pkgs/platform/:pkgPlatform", u.GetPkgsByPlatform)       // get all packages by platform
	g.GET("/pkgs/platformid/:pkgPlatformId", u.GetPkgsByPlatformId) // get all packages by platform Id
	g.GET("/pkgs/type/:pkgType", u.GetPkgsByType)                   // get all packages by type
	g.GET("/pkgs/typeid/:pkgTypeId", u.GetPkgsByTypeId)             // get all packages by type Id
	g.GET("/pkg/arch/:pkgArch", u.GetPkgArch)                       // get package architecture
	g.GET("/pkg/archid/:pkgArchId", u.GetPkgArchById)               // get package architecture by Id
	g.GET("/pkg/id/:pkgId", u.GetPackageById)                       // get package by Id
	g.GET("/pkg/id/:pkgId/versions", u.GetVersions)                 // get all versions of a package
	g.GET("/pkg/name/:pkgName", u.GetPackage)                       // get package by name
	g.GET("/pkg/name/:pkgName/versions", u.GetVersionsByName)       // get all versions of a package by name
	g.GET("/pkg/platform/:pkgPlatform", u.GetPkgPlatform)           // get package platform
	g.GET("/pkg/platformid/:pkgPlatformId", u.GetPkgPlatformById)   // get package platform by Id
	g.GET("/pkg/type/:pkgType", u.GetPkgType)                       // get package type
	g.GET("/pkg/typeid/:pkgTypeId", u.GetPkgTypeById)               // get package type by Id
	g.POST("/pkg", u.CreatePackage)                                 // create new package
	g.POST("/pkg/type/:pkgType", u.CreatePkgType)                   // create new package type
	g.PATCH("/pkg/name/:pkgName", u.UpdatePkg)                      // update a package
	g.DELETE("/pkg/id/:pkgId", u.DeletePackage)                     // delete a package by Id
	g.DELETE("/pkg/name/:pkgName", u.DeletePackageByName)           // delete a package by name
	g.DELETE("/pkg/type/:pkgType", u.DeletePkgType)                 // delete a package type
	// Roles
	g.GET("/roles", u.GetRoles)                    // get all roles
	g.GET("/role/id/:roleId", u.GetRoleById)       // get role by Id
	g.GET("/role/name/:roleName", u.GetRoleByName) // get role by name
	g.POST("/role", u.CreateRole)                  // create new role
	g.DELETE("/role/:roleId", u.DeleteRole)        // delete a role by Id
	// systems
	g.GET("/systems", u.GetSystems)                              // get all systems
	g.GET("/system/id/:sysId", u.GetSystemById)                  // get system by Id
	g.GET("/system/name/:sysName", u.GetSystemByName)            // get system by name
	g.GET("/system/os/:sysOsFamilyId", u.GetSystemsByOsFamilyId) // get all systems by OS family Id
	g.GET("/system/arch/:sysArchId", u.GetSystemsByArchId)       // get all systems by architecture Id
	g.POST("/system", u.CreateSystem)                            // create new system
	g.DELETE("/system/:sysId", u.DeleteSystem)                   // delete a system by Id
	// user related routes
	g.GET("/users", u.GetUsers)                          // get all users
	g.GET("/users/roleId/:roleId", u.GetUsersByRoleId)   // get all users by role Id
	g.GET("/user/name/:name", u.GetUserByUserName)       // get a user by username
	g.GET("/user/name/:name/status", u.GetUserStatus)    // get whether a user is locked or not
	g.GET("/user/id/:id", u.GetUserById)                 // get a user by Id
	g.POST("/user", u.CreateUser)                        // create new user
	g.PATCH("/user/name/:name", u.ChangeAccountPassword) // update a user password
	g.PATCH("/user/name/:name/status", u.SetUserStatus)  // lock a user
	g.PATCH("/user/name/:name/roleId", u.SetUserRoleId)  // set a user's role Id
	g.DELETE("/user/name/:name", u.DeleteUser)           // trash a user
}

func PublicRoutes(g *gin.RouterGroup, u *controllers.UpdateReporter) {
	// service related routes
	g.GET("/health", u.GetHealth) // service health API
}
