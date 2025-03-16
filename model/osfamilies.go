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

	t.Commit()

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

	t.Commit()

	log.Println("INFO: OS Family with Id '" + osfIdStr + "' has been deleted")
	return true, nil
}

func GetOSFamilyById(osfId int) (OSFamily, error) {
	osfIdStr := strconv.Itoa(osfId)
	log.Println("INFO: OS Family retrieval requested: " + osfIdStr)
	var osf OSFamily
	q := "SELECT FamilyName FROM OSFamilies WHERE Id IS ?"
	err := DB.QueryRow(q, osfId).Scan(&osf.FamilyName)
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
	q := "SELECT Id FROM OSFamilies WHERE FamilyName IS ?"
	err := DB.QueryRow(q, osfName).Scan(&osf.Id)
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
