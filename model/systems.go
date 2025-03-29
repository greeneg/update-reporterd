package model

import (
	"log"
	"strconv"
)

func CreateSystem(s System) (bool, error) {
	log.Println("INFO: System creation requested: " + s.Fqdn)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("INSERT INTO Systems (FQDN, OSFamilyId, OsId, ArchId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(s.Fqdn, s.OSFamilyId, s.OSId, s.ArchitectureId)
	if err != nil {
		log.Println("ERROR: Cannot create system '" + s.Fqdn + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

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

	q, err := t.Prepare("DELETE FROM Systems WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(sysId)
	if err != nil {
		log.Println("ERROR: Cannot delete system '" + sysIdStr + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	log.Println("INFO: System '" + sysIdStr + "' deleted")
	return true, nil
}

func GetSystems() ([]System, error) {
	log.Println("INFO: System list requested")
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return nil, err
	}

	q, err := t.Prepare("SELECT * FROM Systems")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}

	rows, err := q.Query()
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
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return System{}, err
	}

	q, err := t.Prepare("SELECT * FROM Systems WHERE Id = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return System{}, err
	}

	rows, err := q.Query(sysId)
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
		return System{}, nil
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
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return System{}, err
	}

	q, err := t.Prepare("SELECT * FROM Systems WHERE FQDN = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return System{}, err
	}

	rows, err := q.Query(sysName)
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
		return System{}, nil
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
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return nil, err
	}

	q, err := t.Prepare("SELECT * FROM Systems WHERE ArchId = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}

	rows, err := q.Query(archId)
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
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return nil, err
	}

	q, err := t.Prepare("SELECT * FROM Systems WHERE OSFamilyId = ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return nil, err
	}

	rows, err := q.Query(osfId)
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
