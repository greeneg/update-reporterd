package model

import (
	"log"
	"strconv"
)

// CreateArchitecture creates a new architecture
// param a Architecture
// return bool, error
func CreateArchitecture(a Architecture) (bool, error) {
	log.Println("INFO: Architecture creation requested: " + a.ArchitectureName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during architecture creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			t.Rollback()
		}
	}()

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

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' created")
	return true, nil
}

// DeleteArchitecture deletes an architecture
// param archId int
// return bool, error
func DeleteArchitecture(archId int) (bool, error) {
	archIdStr := strconv.Itoa(archId)
	log.Println("INFO: Architecture deletion requested: " + archIdStr)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during architecture deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			t.Rollback()
		}
	}()

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

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: Architecture with Id '" + archIdStr + "' has been deleted")
	return true, nil
}

// GetArchitectureById gets an architecture by Id
// param archId int
// return Architecture, error
func GetArchitectureById(archId int) (Architecture, error) {
	archIdStr := strconv.Itoa(archId)
	log.Println("INFO: Architecture retrieval requested: " + archIdStr)
	var a Architecture
	q := "SELECT * FROM Architectures WHERE Id IS ?"
	r, err := DB.Query(q, archId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archIdStr + "': " + string(err.Error()))
		return a, err
	}
	defer r.Close()

	err = r.Scan(&a.Id, &a.ArchitectureName, &a.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archIdStr + "': " + string(err.Error()))
		return a, err
	}
	if a.Id == 0 {
		log.Println("ERROR: Cannot retrieve architecture '" + archIdStr + "': Architecture not found")
		return a, err
	}

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' retrieved")
	return a, nil
}

// GetArchitectureByName gets an architecture by name
// param archName string
// return Architecture, error
func GetArchitectureByName(archName string) (Architecture, error) {
	log.Println("INFO: Architecture retrieval requested: " + archName)
	var a Architecture
	q := "SELECT * FROM Architectures WHERE ArchName IS ?"
	r, err := DB.Query(q, archName)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archName + "': " + string(err.Error()))
		return a, err
	}
	defer r.Close()

	err = r.Scan(&a.Id, &a.ArchitectureName, &a.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architecture '" + archName + "': " + string(err.Error()))
		return a, err
	}
	if a.Id == 0 {
		log.Println("ERROR: Cannot retrieve architecture '" + archName + "': Architecture not found")
		return a, err
	}
	if a.ArchitectureName == "" {
		log.Println("ERROR: Cannot retrieve architecture '" + archName + "': Architecture not found")
		return a, err
	}

	log.Println("INFO: Architecture '" + a.ArchitectureName + "' retrieved")
	return a, nil
}

// GetArchitectures gets all architectures
// param none
// return []Architecture, error
func GetArchitectures() ([]Architecture, error) {
	log.Println("INFO: All architectures retrieval requested")
	var archs []Architecture
	q := "SELECT * FROM Architectures"
	rows, err := DB.Query(q)
	if err != nil {
		log.Println("ERROR: Cannot retrieve architectures: " + string(err.Error()))
		return archs, err
	}
	defer rows.Close()

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
