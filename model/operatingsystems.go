package model

import (
	"log"
	"strconv"
)

func CreateOperatingSystem(o OperatingSystem) (bool, error) {
	log.Println("INFO: OS creation requested: " + o.OsName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("INSERT INTO OperatingSystems (OsName, OsVersion, OsFamilyId, OsArchId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(o.OsName, o.OsVersion, o.OsFamilyId, o.ArchitectureId)
	if err != nil {
		log.Println("ERROR: Cannot create OS '" + o.OsName + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: OS '" + o.OsName + "' created")
	return true, nil
}

func DeleteOperatingSystem(osId int) (bool, error) {
	osIdStr := strconv.Itoa(osId)
	log.Println("INFO: OS deletion requested: " + osIdStr)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("DELETE FROM OperatingSystems WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(osId)
	if err != nil {
		log.Println("ERROR: Cannot delete OS '" + osIdStr + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: OS '" + osIdStr + "' deleted")
	return true, nil
}

func GetOperatingSystemById(osId int) (OperatingSystem, error) {
	osIdStr := strconv.Itoa(osId)
	log.Println("INFO: OS retrieval requested: " + osIdStr)
	var osStruct OperatingSystem

	q := "SELECT Id, OsName, OsVersion, OsFamilyId, OsArchId, CreationDate FROM OperatingSystems WHERE Id = ?"
	err := DB.QueryRow(q, osId).Scan(&osStruct.Id, &osStruct.OsName, &osStruct.OsVersion, &osStruct.OsFamilyId, &osStruct.ArchitectureId, &osStruct.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS '" + osIdStr + "': " + string(err.Error()))
		return osStruct, err
	}

	log.Println("INFO: OS '" + osIdStr + "' retrieved")
	return osStruct, nil
}

func GetOperatingSystemByName(osName string) (OperatingSystem, error) {
	log.Println("INFO: OS retrieval requested: " + osName)
	var osStruct OperatingSystem

	q := "SELECT Id, OsName, OsVersion, OsFamilyId, OsArchId, CreationDate FROM OperatingSystems WHERE OSName = ?"
	err := DB.QueryRow(q, osName).Scan(&osStruct.Id, &osStruct.OsName, &osStruct.OsVersion, &osStruct.OsFamilyId, &osStruct.ArchitectureId, &osStruct.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS '" + osName + "': " + string(err.Error()))
		return osStruct, err
	}

	log.Println("INFO: OS '" + osName + "' retrieved")
	return osStruct, nil
}

func GetOperatingSystems() ([]OperatingSystem, error) {
	log.Println("INFO: OS list retrieval requested")
	var osStruct OperatingSystem
	var osList []OperatingSystem

	q := "SELECT Id, OsName, OsVersion, OsFamilyId, OsArchId, CreationDate FROM OperatingSystems"
	rows, err := DB.Query(q)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS list: " + string(err.Error()))
		return osList, err
	}

	for rows.Next() {
		err = rows.Scan(&osStruct.Id, &osStruct.OsName, &osStruct.OsVersion, &osStruct.OsFamilyId, &osStruct.ArchitectureId, &osStruct.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve OS list: " + string(err.Error()))
			return osList, err
		}
		osList = append(osList, osStruct)
	}

	log.Println("INFO: OS list retrieved")
	return osList, nil
}
