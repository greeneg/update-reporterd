package model

import (
	"log"
	"strconv"
	"time"
)

func GetPackages() ([]Package, error) {
	log.Println("INFO: Packages retrieval requested")

	var pkgs []Package
	q := "SELECT * FROM Packages"
	rows, err := DB.Query(q)
	if err != nil {
		log.Println("ERROR: Cannot retrieve os families: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Package
		err = rows.Scan(&p.Id, &p.PackageName, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve OS Family: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, p)
	}

	log.Println("INFO: All operating system families retrieved")
	return pkgs, nil
}

func GetPkgsByArch(pkgArch string) ([]Package, error) {
	log.Println("INFO: Packages retrieval requested: " + pkgArch)

	archId, err := GetArchitectureByName(pkgArch)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture: " + string(err.Error()))
		return nil, err
	}

	var pkgs []Package
	q := "SELECT * FROM PackagesByArchitecture WHERE ArchitectureId IS ?"
	rows, err := DB.Query(q, archId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByArch
		err = rows.Scan(&p.Id, &p.PackageId, &p.ArchId, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPkgsByArchId(pkgArchId int) ([]Package, error) {
	pkgArchStr := strconv.Itoa(pkgArchId)
	log.Println("INFO: Packages retrieval requested: " + pkgArchStr)

	arch, err := GetArchitectureById(pkgArchId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package architecture: " + string(err.Error()))
		return nil, err
	}

	var pkgs []Package
	q := "SELECT * FROM PackagesByArchitecture WHERE ArchitectureId IS ?"
	rows, err := DB.Query(q, arch.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByArch
		err = rows.Scan(&p.Id, &p.PackageId, &p.ArchId, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPkgsByPlatform(pkgPlatform string) ([]Package, error) {
	log.Println("INFO: Packages retrieval requested: " + pkgPlatform)

	plat, err := GetOperatingSystemByName(pkgPlatform)
	if err != nil {
		log.Println("ERROR: Cannot retrieve operating system: " + string(err.Error()))
		return nil, err
	}

	var pkgs []Package
	q := "SELECT * FROM PackagesByPlatform WHERE PlatformId IS ?"
	rows, err := DB.Query(q, plat.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var pkgPlat PackageByPlatform
		err = rows.Scan(&pkgPlat.Id, &pkgPlat.PackageId, &pkgPlat.PlatformId, &pkgPlat.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(pkgPlat.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPkgsByPlatformId(pkgPlatformId int) ([]Package, error) {
	pkgPlatform := strconv.Itoa(pkgPlatformId)
	log.Println("INFO: Packages retrieval requested: " + pkgPlatform)

	plat, err := GetOperatingSystemById(pkgPlatformId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package platform: " + string(err.Error()))
		return nil, err
	}

	var pkgs []Package
	q := "SELECT * FROM PackagesByPlatform WHERE PackagePlat IS ?"
	rows, err := DB.Query(q, plat.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var pkgPlat PackageByPlatform
		err = rows.Scan(&pkgPlat.Id, &pkgPlat.PackageId, &pkgPlat.PlatformId, &pkgPlat.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(pkgPlat.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPkgsByType(pkgType string) ([]Package, error) {
	log.Println("INFO: Packages retrieval requested: " + pkgType)

	pkgTypeStruct, err := GetPkgTypeByName(pkgType)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type: " + string(err.Error()))
		return nil, err
	}

	var pkgs []Package
	q := "SELECT * FROM PackagesByType WHERE PackageTypeId IS ?"
	rows, err := DB.Query(q, pkgTypeStruct.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByType
		err = rows.Scan(&p.Id, &p.PackageId, &p.PackageTypeId, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPkgsByTypeId(pkgTypeId int) ([]Package, error) {
	pkgTypeIdStr := strconv.Itoa(pkgTypeId)
	log.Println("INFO: Packages retrieval requested: " + pkgTypeIdStr)

	var pkgs []Package
	q := "SELECT * FROM PackagesByType WHERE PackageTypeId IS ?"
	rows, err := DB.Query(q, pkgTypeId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve packages: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByType
		err = rows.Scan(&p.Id, &p.PackageId, &p.PackageTypeId, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}

	log.Println("INFO: All packages retrieved")
	return pkgs, nil
}

func GetPackageById(pkgId int) (Package, error) {
	pkgIdStr := strconv.Itoa(pkgId)
	log.Println("INFO: Package retrieval requested: " + pkgIdStr)

	var p Package
	q := "SELECT * FROM Packages WHERE Id IS ?"
	r, err := DB.Query(q, pkgId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package '" + pkgIdStr + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	err = r.Scan(&p.Id, &p.PackageName, &p.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package '" + pkgIdStr + "': " + string(err.Error()))
		return p, err
	}

	log.Println("INFO: Package '" + p.PackageName + "' retrieved")
	return p, nil
}

func GetPackage(pkgName string) (Package, error) {
	log.Println("INFO: Package retrieval requested: " + pkgName)

	var p Package
	q := "SELECT * FROM Packages WHERE PackageName IS ?"
	r, err := DB.Query(q, pkgName)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package '" + pkgName + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	r.Scan(&p.Id, &p.PackageName, &p.CreationDate)

	log.Println("INFO: Package '" + p.PackageName + "' retrieved")
	return p, nil
}

func GetPkgVersionsById(pkgId string) ([]PackageByVersionName, error) {
	log.Println("INFO: Package versions retrieval requested: " + pkgId)

	var pkgs []PackageByVersionName
	q := "SELECT * FROM PackagesByVersion WHERE PackageId IS ?"
	rows, err := DB.Query(q, pkgId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package versions: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByVersion
		err = rows.Scan(&p.Id, &p.PackageId, &p.Version, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package version: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		namedPkg := PackageByVersionName{Id: pkg.Id, PackageName: pkg.PackageName, Version: p.Version, CreationDate: p.CreationDate}
		pkgs = append(pkgs, namedPkg)
	}

	log.Println("INFO: All package versions retrieved")
	return pkgs, nil
}

func GetVersionsByName(pkgName string) ([]PackageByVersionName, error) {
	log.Println("INFO: Package versions retrieval requested: " + pkgName)

	pkg, err := GetPackage(pkgName)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return nil, err
	}

	var pkgs []PackageByVersionName
	q := "SELECT * FROM PackagesByVersion WHERE PackageId IS ?"
	rows, err := DB.Query(q, pkg.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package versions: " + string(err.Error()))
		return pkgs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PackageByVersion
		err = rows.Scan(&p.Id, &p.PackageId, &p.Version, &p.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package version: " + string(err.Error()))
			return pkgs, err
		}
		pkg, err := GetPackageById(p.PackageId)
		if err != nil {
			log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
			return pkgs, err
		}
		namedPkg := PackageByVersionName{Id: pkg.Id, PackageName: pkg.PackageName, Version: p.Version, CreationDate: p.CreationDate}
		pkgs = append(pkgs, namedPkg)
	}

	log.Println("INFO: All package versions retrieved")
	return pkgs, nil
}

func GetPkgTypeByName(pkgType string) (PackageType, error) {
	log.Println("INFO: Package type retrieval requested: " + pkgType)

	var p PackageType

	q := "SELECT * FROM PackageTypes WHERE Name IS ?"
	r, err := DB.Query(q, pkgType)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type '" + pkgType + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	err = r.Scan(&p.Id, &p.TypeName, &p.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type '" + pkgType + "': " + string(err.Error()))
		return p, err
	}

	log.Println("INFO: Package type '" + p.TypeName + "' retrieved")
	return p, nil
}

func GetPkgTypeById(pkgTypeId int) (PackageType, error) {
	pkgTypeStr := strconv.Itoa(pkgTypeId)
	log.Println("INFO: Package type retrieval requested: " + pkgTypeStr)

	var p PackageType
	q := "SELECT * FROM PackageTypes WHERE Id IS ?"
	r, err := DB.Query(q, pkgTypeId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type '" + pkgTypeStr + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	err = r.Scan(&p.Id, &p.TypeName, &p.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type '" + pkgTypeStr + "': " + string(err.Error()))
		return p, err
	}

	log.Println("INFO: Package type '" + p.TypeName + "' retrieved")
	return p, nil
}

func GetPkgPlatform(pkgPlatform string) (OperatingSystem, error) {
	log.Println("INFO: Package platform retrieval requested: " + pkgPlatform)

	var p OperatingSystem
	q := "SELECT * FROM OperatingSystems WHERE OSName IS ?"
	r, err := DB.Query(q, pkgPlatform)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package platform '" + pkgPlatform + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	err = r.Scan(&p.Id, &p.OsName, &p.OsVersion, &p.OsFamilyId, &p.ArchitectureId, &p.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package platform '" + pkgPlatform + "': " + string(err.Error()))
		return p, err
	}

	log.Println("INFO: Package platform '" + p.OsName + "' retrieved")
	return p, nil
}

func GetPkgPlatformById(pkgPlatformId int) (OperatingSystem, error) {
	pkgPlatformStr := strconv.Itoa(pkgPlatformId)
	log.Println("INFO: Package platform retrieval requested: " + pkgPlatformStr)

	var p OperatingSystem

	q := "SELECT * FROM OperatingSystems WHERE Id IS ?"
	r, err := DB.Query(q, pkgPlatformId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package platform '" + pkgPlatformStr + "': " + string(err.Error()))
		return p, err
	}
	defer r.Close()

	err = r.Scan(&p.Id, &p.OsName, &p.OsVersion, &p.OsFamilyId, &p.ArchitectureId, &p.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package platform '" + pkgPlatformStr + "': " + string(err.Error()))
		return p, err
	}

	log.Println("INFO: Package platform '" + p.OsName + "' retrieved")
	return p, nil
}

func CreatePackage(pkg *Package) error {
	log.Println("INFO: Package creation requested: " + pkg.PackageName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO Packages (PackageName, CreationDate) VALUES (?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkg.PackageName, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkg.PackageName + "' created")
	return nil
}

func CreatePackageByArch(pkg *PackageByArch) error {
	archIdStr := strconv.Itoa(pkg.ArchId)
	arch, err := GetArchitectureById(pkg.ArchId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture: " + string(err.Error()))
		return err
	}
	log.Println("INFO: Package creation requested: " + archIdStr + " for architecture: " + arch.ArchitectureName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackagesByArchitecture (PackageId, ArchId, CreationDate) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkg.PackageId, pkg.ArchId, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	pkgStruct, err := GetPackageById(pkg.PackageId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgStruct.PackageName + "' created")
	return nil
}

func CreatePackageByPlatform(pkg *PackageByPlatform) error {
	platIdStr := strconv.Itoa(pkg.PlatformId)
	plat, err := GetOperatingSystemById(pkg.PlatformId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve platform: " + string(err.Error()))
		return err
	}
	log.Println("INFO: Package creation requested: " + platIdStr + " for platform: " + plat.OsName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackagesByPlatform (PackageId, PlatformId, CreationDate) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package by platform creation: " + string(err.Error()))
		return err
	}
	_, err = q.Exec(pkg.PackageId, pkg.PlatformId, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package by platform creation: " + string(err.Error()))
		return err
	}
	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	pkgStruct, err := GetPackageById(pkg.PackageId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgStruct.PackageName + "' created")
	return nil
}

func CreatePackageByType(pkg *PackageByType) error {
	pkgTypeIdStr := strconv.Itoa(pkg.PackageTypeId)
	pkgType, err := GetPkgTypeById(pkg.PackageTypeId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type: " + string(err.Error()))
		return err
	}
	log.Println("INFO: Package creation requested: " + pkgTypeIdStr + " for type: " + pkgType.TypeName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackagesByType (PackageId, PackageTypeId, CreationDate) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package by type creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkg.PackageId, pkg.PackageTypeId, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package by type creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	pkgStruct, err := GetPackageById(pkg.PackageId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgStruct.PackageName + "' created")
	return nil
}

func CreatePackageVersion(pkgVersion *PackageByVersion) error {
	pkgIdStr := strconv.Itoa(pkgVersion.PackageId)
	log.Println("INFO: Package version creation requested: " + pkgIdStr)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package version creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package version: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackagesByVersion (PackageId, Version, CreationDate) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package version creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgVersion.PackageId, pkgVersion.Version, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package version creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	pkg, err := GetPackageById(pkgVersion.PackageId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkg.PackageName + "' version '" + pkgVersion.Version + "' created")
	return nil
}

func CreatePackageVersionByName(pkgName, pkgVersion string) error {
	log.Println("INFO: Package version creation requested: " + pkgName)

	pkg, err := GetPackage(pkgName)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package version creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package version: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackagesByVersion (PackageId, Version, CreationDate) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package version creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkg.Id, pkgVersion, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package version creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	pkgStruct, err := GetPackageById(pkg.Id)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgStruct.PackageName + "' version '" + pkgVersion + "' created")
	return nil
}

func CreatePkgType(pkgType *PackageType) error {
	log.Println("INFO: Package type creation requested: " + pkgType.TypeName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package type creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package type: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackageTypes (Name, CreationDate) VALUES (?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package type creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgType.TypeName, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package type creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package type '" + pkgType.TypeName + "' created")
	return nil
}

func CreatePkgTypeByName(pkgTypeName string) error {
	log.Println("INFO: Package type creation requested: " + pkgTypeName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package type creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package type: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO PackageTypes (Name, CreationDate) VALUES (?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package type creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgTypeName, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package type creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package type '" + pkgTypeName + "' created")
	return nil
}

func CreatePkgPlatform(pkgPlatform *OperatingSystem) error {
	log.Println("INFO: Package platform creation requested: " + pkgPlatform.OsName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package platform creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package platform: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO OperatingSystems (OSName, OSVersion, OSFamilyId, ArchitectureId, CreationDate) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package platform creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgPlatform.OsName, pkgPlatform.OsVersion, pkgPlatform.OsFamilyId, pkgPlatform.ArchitectureId, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package platform creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package platform '" + pkgPlatform.OsName + "' created")
	return nil
}

func CreatePkgPlatformByName(pkgPlatformName string) error {
	log.Println("INFO: Package platform creation requested: " + pkgPlatformName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package platform creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot create package platform: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("INSERT INTO OperatingSystems (OSName, CreationDate) VALUES (?, ?)")
	if err != nil {
		log.Println("ERROR: Cannot prepare package platform creation: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgPlatformName, sqliteTimeStamp)
	if err != nil {
		log.Println("ERROR: Cannot execute package platform creation: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package platform '" + pkgPlatformName + "' created")
	return nil
}

func UpdatePkg(pkg *Package) error {
	log.Println("INFO: Package update requested: " + pkg.PackageName)

	sqliteTimeStamp := ConvertSqliteTimestamp(time.Now().String())

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package update: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot update package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("UPDATE Packages SET PackageName = ?, CreationDate = ? WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Cannot prepare package update: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkg.PackageName, sqliteTimeStamp, pkg.Id)
	if err != nil {
		log.Println("ERROR: Cannot execute package update: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkg.PackageName + "' updated")
	return nil
}

func DeletePkg(pkgId int) error {
	pkgIdStr := strconv.Itoa(pkgId)
	log.Println("INFO: Package deletion requested: " + pkgIdStr)

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot delete package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("DELETE FROM Packages WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Cannot prepare package deletion: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgId)
	if err != nil {
		log.Println("ERROR: Cannot execute package deletion: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgIdStr + "' deleted")
	return nil
}

func DeletePkgByName(pkgName string) error {
	log.Println("INFO: Package deletion requested: " + pkgName)

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot delete package: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("DELETE FROM Packages WHERE PackageName IS ?")
	if err != nil {
		log.Println("ERROR: Cannot prepare package deletion: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgName)
	if err != nil {
		log.Println("ERROR: Cannot delete package: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}

	log.Println("INFO: Package '" + pkgName + "' deleted")
	return nil
}

func DeletePkgType(pkgTypeId int) error {
	pkgTypeIdStr := strconv.Itoa(pkgTypeId)
	log.Println("INFO: Package type deletion requested: " + pkgTypeIdStr)

	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Cannot begin transaction: " + string(err.Error()))
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during package type deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Cannot delete package type: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("DELETE FROM PackageTypes WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Cannot prepare package type deletion: " + string(err.Error()))
		return err
	}

	_, err = q.Exec(pkgTypeId)
	if err != nil {
		log.Println("ERROR: Cannot execute package type deletion: " + string(err.Error()))
		return err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit transaction: " + string(err.Error()))
		return err
	}
	pkgType, err := GetPkgTypeById(pkgTypeId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve package type: " + string(err.Error()))
		return err
	}
	log.Println("INFO: Package type '" + pkgType.TypeName + "' deleted")
	return nil
}
