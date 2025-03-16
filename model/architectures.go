package model

import (
	"log"
	"strconv"
)

func CreateArchitecture(a Architecture) (bool, error) {
	log.Println("INFO: Architecture creation requested: " + a.ArchitectureName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("INSERT INTO Architectures (ArchitectureName) VALUES (?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(a.ArchitectureName)
	if err != nil {
		log.Println("ERROR: Cannot create architecture '" + a.ArchitectureName + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' created")
	return true, nil
}

func DeleteArchitecture(archId int) (bool, error) {
	archIdStr := strconv.Itoa(archId)
	log.Println("INFO: Architecture deletion requested: " + archIdStr)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("DELETE FROM Architectures WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(archId)
	if err != nil {
		log.Println("ERROR: Cannot delete architecture '" + archIdStr + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: Architecture with Id '" + archIdStr + "' has been deleted")
	return true, nil
}

func GetArchitectureById(archId int) (Architecture, error) {
	archIdStr := strconv.Itoa(archId)
	log.Println("INFO: Architecture retrieval requested: " + archIdStr)
	var a Architecture
	q := "SELECT * FROM Architectures WHERE Id IS ?"
	err := DB.QueryRow(q, archId).Scan(&a.Id, &a.ArchitectureName, &a.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archIdStr + "': " + string(err.Error()))
		return a, err
	}

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' retrieved")
	return a, nil
}

func GetArchitectureByName(archName string) (Architecture, error) {
	log.Println("INFO: Architecture retrieval requested: " + archName)
	var a Architecture
	q := "SELECT * FROM Architectures WHERE ArchitectureName IS ?"
	err := DB.QueryRow(q, archName).Scan(&a.Id, &a.ArchitectureName, &a.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archName + "': " + string(err.Error()))
		return a, err
	}

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' retrieved")
	return a, nil
}

func GetArchitectures() ([]Architecture, error) {
	log.Println("INFO: All architectures retrieval requested")
	var archs []Architecture
	q := "SELECT * FROM Architectures"
	rows, err := DB.Query(q)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architectures: " + string(err.Error()))
		return archs, err
	}

	for rows.Next() {
		var a Architecture
		err = rows.Scan(&a.Id, &a.ArchitectureName, &a.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve architecture: " + string(err.Error()))
			return archs, err
		}
		archs = append(archs, a)
	}

	log.Println("INFO: All architectures retrieved")
	return archs, nil
}
