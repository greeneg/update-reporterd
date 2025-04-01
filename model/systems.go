package model

import (
	"errors"
	"log"
	"strconv"
)

func CreateSystem(s ProposedSystem) (bool, error) {
	log.Println("INFO: System creation requested: " + s.Fqdn)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during system creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			log.Println("ERROR: Error occurred during system creation: " + string(err.Error()))
			t.Rollback()
		}
	}()

	q, err := t.Prepare("INSERT INTO Systems (FQDN, OsFamilyId, OsId, ArchId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(s.Fqdn, s.OSFamilyId, s.OSId, s.ArchitectureId)
	if err != nil {
		log.Println("ERROR: Cannot create system '" + s.Fqdn + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: System '" + s.Fqdn + "' created")
	return true, nil
}

func DeleteSystem(sysId int) (bool, error) {
	// convert sysId to string
	sysIdStr := strconv.Itoa(sysId)
	log.Println("INFO: System deletion requested: " + sysIdStr)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during system deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			t.Rollback()
		}
	}()

	q, err := t.Prepare("DELETE FROM Systems WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	res, err := q.Exec(sysId)
	if err != nil {
		log.Println("ERROR: Cannot delete system '" + sysIdStr + "': " + string(err.Error()))
		return false, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("ERROR: Cannot get the number of affected rows: " + string(err.Error()))
		return false, err
	}
	if rows == 0 {
		log.Println("ERROR: No system found with id " + sysIdStr)
		return false, &UnknownSystemById{Err: errors.New("No system found with id " + sysIdStr)}
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: System '" + sysIdStr + "' deleted")
	return true, nil
}

func GetSystems() ([]System, error) {
	log.Println("INFO: System list requested")

	stmt, err := DB.Prepare("SELECT * FROM Systems")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: Cannot query systems: " + string(err.Error()))
		return nil, err
	}
	defer rows.Close()

	var systems []System
	for rows.Next() {
		var s System
		err = rows.Scan(&s.Id, &s.Fqdn, &s.OSFamilyId, &s.OSId, &s.ArchitectureId, &s.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot assign system to system object: " + string(err.Error()))
			return nil, err
		}
		systems = append(systems, s)
	}
	if err = rows.Err(); err != nil {
		log.Println("ERROR: Cannot iterate over systems: " + string(err.Error()))
		return nil, err
	}

	if len(systems) == 0 {
		log.Println("INFO: No systems found")
		return nil, nil
	}

	log.Println("INFO: " + strconv.Itoa(len(systems)) + " systems found")
	return systems, nil
}

func GetSystemById(sysId int) (System, error) {
	log.Println("INFO: System requested: " + strconv.Itoa(sysId))
	stmt, err := DB.Prepare("SELECT * FROM Systems WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return System{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sysId)
	if err != nil {
		log.Println("ERROR: Cannot query system: " + string(err.Error()))
		return System{}, err
	}
	defer rows.Close()

	var s System
	if rows.Next() {
		err = rows.Scan(&s.Id, &s.Fqdn, &s.OSFamilyId, &s.OSId, &s.ArchitectureId, &s.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot assign system to system object: " + string(err.Error()))
			return System{}, err
		}
	} else {
		log.Println("ERROR: No system found with id " + strconv.Itoa(sysId))
		return System{}, &UnknownSystemById{Err: errors.New("No system found with id " + strconv.Itoa(sysId))}
	}

	if err = rows.Err(); err != nil {
		log.Println("ERROR: Cannot iterate over system: " + string(err.Error()))
		return System{}, err
	}
	log.Println("INFO: System '" + s.Fqdn + "' found")
	return s, nil
}

func GetSystemByName(sysName string) (System, error) {
	log.Println("INFO: System requested: " + sysName)
	stmt, err := DB.Prepare("SELECT * FROM Systems WHERE FQDN = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return System{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sysName)
	if err != nil {
		log.Println("ERROR: Cannot query system: " + string(err.Error()))
		return System{}, err
	}
	defer rows.Close()

	var s System
	if rows.Next() {
		err = rows.Scan(&s.Id, &s.Fqdn, &s.OSFamilyId, &s.OSId, &s.ArchitectureId, &s.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot assign system to system object: " + string(err.Error()))
			return System{}, err
		}
	} else {
		log.Println("ERROR: No system found with name " + sysName)
		return System{}, &UnknownSystemByName{}
	}

	if err = rows.Err(); err != nil {
		log.Println("ERROR: Cannot iterate over system: " + string(err.Error()))
		return System{}, err
	}
	log.Println("INFO: System '" + s.Fqdn + "' found")
	return s, nil
}

func GetSystemsByArchId(archId int) ([]System, error) {
	log.Println("INFO: System list requested by architecture id: " + strconv.Itoa(archId))
	stmt, err := DB.Prepare("SELECT * FROM Systems WHERE ArchId = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(archId)
	if err != nil {
		log.Println("ERROR: Cannot query systems: " + string(err.Error()))
		return nil, err
	}
	defer rows.Close()

	var systems []System
	for rows.Next() {
		var s System
		err = rows.Scan(&s.Id, &s.Fqdn, &s.OSFamilyId, &s.OSId, &s.ArchitectureId, &s.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot assign system to system object: " + string(err.Error()))
			return nil, err
		}
		systems = append(systems, s)
	}

	if err = rows.Err(); err != nil {
		log.Println("ERROR: Cannot iterate over systems: " + string(err.Error()))
		return nil, err
	}

	if len(systems) == 0 {
		log.Println("INFO: No systems found")
		return nil, nil
	}
	log.Println("INFO: " + strconv.Itoa(len(systems)) + " systems found")
	return systems, nil
}

func GetSystemsByOsFamilyId(osfId int) ([]System, error) {
	log.Println("INFO: System list requested by OS family id: " + strconv.Itoa(osfId))
	stmt, err := DB.Prepare("SELECT * FROM Systems WHERE OSFamilyId = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(osfId)
	if err != nil {
		log.Println("ERROR: Cannot query systems: " + string(err.Error()))
		return nil, err
	}
	defer rows.Close()

	var systems []System
	for rows.Next() {
		var s System
		err = rows.Scan(&s.Id, &s.Fqdn, &s.OSFamilyId, &s.OSId, &s.ArchitectureId, &s.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot assign system to system object: " + string(err.Error()))
			return nil, err
		}
		systems = append(systems, s)
	}
	if err = rows.Err(); err != nil {
		log.Println("ERROR: Cannot iterate over systems: " + string(err.Error()))
		return nil, err
	}

	if len(systems) == 0 {
		log.Println("INFO: No systems found")
		return nil, nil
	}
	log.Println("INFO: " + strconv.Itoa(len(systems)) + " systems found")
	return systems, nil
}
