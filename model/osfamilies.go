package model

import (
	"log"
	"strconv"
)

func CreateOSFamily(o OSFamily) (bool, error) {
	log.Println("INFO: OS Family creation requested: " + o.FamilyName)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during OS Family creation: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			t.Rollback()
		}
	}()

	q, err := t.Prepare("INSERT INTO OSFamilies (FamilyName) VALUES (?)")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(o.FamilyName)
	if err != nil {
		log.Println("ERROR: Cannot create OS Family '" + o.FamilyName + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: OS Family '" + o.FamilyName + "' created")
	return true, nil
}

func DeleteOSFamily(osfId int) (bool, error) {
	osfIdStr := strconv.Itoa(osfId)
	log.Println("INFO: OS Family deletion requested: " + osfIdStr)
	t, err := DB.Begin()
	if err != nil {
		log.Println("ERROR: Could not start DB transaction!" + string(err.Error()))
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: Panic occurred during OS Family deletion: " + string(r.(error).Error()))
			t.Rollback()
		}
		if err != nil {
			t.Rollback()
		}
	}()

	q, err := DB.Prepare("DELETE FROM OSFamilies WHERE Id IS ?")
	if err != nil {
		log.Println("ERROR: Could not prepare the DB query!" + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(osfId)
	if err != nil {
		log.Println("ERROR: Cannot delete OS Family '" + osfIdStr + "': " + string(err.Error()))
		return false, err
	}

	err = t.Commit()
	if err != nil {
		log.Println("ERROR: Cannot commit the DB transaction!" + string(err.Error()))
		return false, err
	}

	log.Println("INFO: OS Family with Id '" + osfIdStr + "' has been deleted")
	return true, nil
}

func GetOSFamilyById(osfId int) (OSFamily, error) {
	osfIdStr := strconv.Itoa(osfId)
	log.Println("INFO: OS Family retrieval requested: " + osfIdStr)
	var osf OSFamily
	q := "SELECT * FROM OSFamilies WHERE Id IS ?"
	r, err := DB.Query(q, osfId)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS Family '" + osfIdStr + "': " + string(err.Error()))
		return osf, err
	}
	defer r.Close()

	err = r.Scan(&osf.Id, &osf.FamilyName, &osf.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS Family '" + osfIdStr + "': " + string(err.Error()))
		return osf, err
	}

	log.Println("INFO: OS Family '" + osf.FamilyName + "' retrieved")
	return osf, nil
}

func GetOSFamilyByName(osfName string) (OSFamily, error) {
	log.Println("INFO: OS Family retrieval requested: " + osfName)
	var osf OSFamily
	q := "SELECT * FROM OSFamilies WHERE FamilyName IS ?"
	rows, err := DB.Query(q, osfName)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS Family '" + osfName + "': " + string(err.Error()))
		return osf, err
	}
	defer rows.Close()

	err = rows.Scan(&osf.Id, &osf.FamilyName, &osf.CreationDate)
	if err != nil {
		log.Println("ERROR: Cannot retrieve OS Family '" + osfName + "': " + string(err.Error()))
		return osf, err
	}

	log.Println("INFO: OS Family '" + osfName + "' retrieved")
	return osf, nil
}

func GetOSFamilies() ([]OSFamily, error) {
	log.Println("INFO: OS Families retrieval requested")

	var osfs []OSFamily
	q := "SELECT * FROM OSFamilies"
	rows, err := DB.Query(q)
	if err != nil {
		log.Println("ERROR: Cannot retrieve os families: " + string(err.Error()))
		return osfs, err
	}
	defer rows.Close()

	for rows.Next() {
		var o OSFamily
		err = rows.Scan(&o.Id, &o.FamilyName, &o.CreationDate)
		if err != nil {
			log.Println("ERROR: Cannot retrieve OS Family: " + string(err.Error()))
			return osfs, err
		}
		osfs = append(osfs, o)
	}

	log.Println("INFO: All operating system families retrieved")
	return osfs, nil
}
