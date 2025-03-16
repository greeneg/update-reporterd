package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"time"
)

type OrgUnit struct {
	Id           int    `json:"Id"`
	OrgUnitName  string `json:"orgUnitName"`
	Description  string `json:"description"`
	CreationDate string `json:"creationDate"`
}

type Role struct {
	Id           int    `json:"Id"`
	RoleName     string `json:"roleName"`
	Description  string `json:"description"`
	CreationDate string `json:"creationDate"`
}

type User struct {
	Id                     int
	UserName               string
	FullName               string
	Status                 string
	OrgUnitId              int
	RoleId                 int
	PasswordHash           string
	CreationDate           string
	LastPasswordChangeDate string
}

func convertSqliteTimestamp(t string) string {
	sqlTimestampFormat := "2006-01-02T15:04:05Z"
	timeFormat := "2006-01-02 15:04:05"
	createTime, _ := time.Parse(sqlTimestampFormat, t)
	return createTime.Format(timeFormat)
}

func getRoleStatus(role string) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("SELECT * FROM Roles WHERE RoleName IS ?")
	if err != nil {
		errPrintln("Could not prepare DB query! " + string(err.Error()))
		return false, err
	}

	rr := Role{}
	err = q.QueryRow(role).Scan(
		&rr.Id,
		&rr.RoleName,
		&rr.Description,
		&rr.CreationDate,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			warnPrintln("Encountered error when querying database: " + string(err.Error()))
			return false, err
		}
		return false, nil
	}

	t.Commit()

	return true, nil
}

func createRole(roleName string, roleDescription string) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("INSERT INTO Roles (RoleName, Description) VALUES (?, ?)")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(roleName, roleDescription)
	if err != nil {
		errPrintln("Cannot create role '" + roleName + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	return true, nil
}

func getRoleByName(roleName string) (Role, error) {
	rec, err := DB.Prepare("SELECT * FROM Roles WHERE RoleName = ?")
	if err != nil {
		errPrintln("Could not prepare the DB query!" + string(err.Error()))
		return Role{}, err
	}

	role := Role{}
	err = rec.QueryRow(roleName).Scan(
		&role.Id,
		&role.RoleName,
		&role.Description,
		&role.CreationDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			warnPrintln("No such role found in DB: " + string(err.Error()))
			return Role{}, nil
		}
		errPrintln("Cannot retrieve role from DB: " + string(err.Error()))
		return Role{}, err
	}

	role.CreationDate = convertSqliteTimestamp(role.CreationDate)
	rec.Close()

	return role, nil
}

func getAccountStatus(account string) (bool, error) {
	println("Checking for account: " + account)
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("SELECT * FROM Users WHERE UserName IS ?")
	if err != nil {
		errPrintln("Could not prepare DB query! " + string(err.Error()))
		return false, err
	}

	u := User{}
	err = q.QueryRow(account).Scan(&u.Id, &u.UserName, &u.FullName, &u.Status, &u.OrgUnitId, &u.RoleId, &u.PasswordHash, &u.CreationDate, &u.LastPasswordChangeDate)
	if err != nil {
		if err != sql.ErrNoRows {
			errPrintln("Encountered error when querying database: " + string(err.Error()))
			return false, err
		}
		return false, err
	}

	t.Commit()

	return true, nil
}

func createAccount(accountName string, accountFullName string, orgUnitId, roleId int, passwd string) (User, error) {
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction!" + string(err.Error()))
		return User{}, err
	}

	q, err := t.Prepare("INSERT INTO Users (UserName, FullName, OrgUnitId, RoleId, Status, PasswordHash) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		errPrintln("Could not prepare the DB query!" + string(err.Error()))
		return User{}, err
	}

	// take password and hash it
	hash := sha512.Sum512([]byte(passwd))
	passwdHash := hex.EncodeToString(hash[:])

	// set the status to active
	status := "active"

	// insert the user

	_, err = q.Exec(accountName, accountFullName, orgUnitId, roleId, status, passwdHash)
	if err != nil {
		errPrintln("Cannot create user '" + accountName + "': " + string(err.Error()))
		return User{}, err
	}

	t.Commit()

	user, err := getAccountByName(accountName)
	if err != nil {
		errPrintln("Could not retrieve user account: " + string(err.Error()))
		return User{}, err
	}

	return user, nil
}

func getAccountByName(accountName string) (User, error) {
	rec, err := DB.Prepare("SELECT Id,UserName,FullName,Status,OrgUnitId,RoleId,CreationDate FROM Users WHERE UserName = ?")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return User{}, err
	}

	user := User{}
	err = rec.QueryRow(accountName).Scan(
		&user.Id,
		&user.UserName,
		&user.FullName,
		&user.Status,
		&user.OrgUnitId,
		&user.RoleId,
		&user.CreationDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			warnPrintln("No such account found in DB named " + accountName + ": " + string(err.Error()))
			return User{}, nil
		}
		errPrintln("Cannot retrieve account from DB named " + accountName + ": " + string(err.Error()))
		return User{}, err
	}

	user.CreationDate = convertSqliteTimestamp(user.CreationDate)

	return user, nil
}

func getOrgUnitByName(orgUnitName string) (OrgUnit, error) {
	rec, err := DB.Prepare("SELECT * FROM OrgUnits WHERE OrgUnitName = ?")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return OrgUnit{}, err
	}

	orgunit := OrgUnit{}
	err = rec.QueryRow(orgUnitName).Scan(
		&orgunit.Id,
		&orgunit.OrgUnitName,
		&orgunit.Description,
		&orgunit.CreationDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			warnPrintln("No such role found in DB: " + string(err.Error()))
			return OrgUnit{}, nil
		}
		errPrintln("Cannot retrieve role from DB: " + string(err.Error()))
		return OrgUnit{}, err
	}

	orgunit.CreationDate = convertSqliteTimestamp(orgunit.CreationDate)

	return orgunit, nil
}

func createOrgUnit(orgUnitName string, orgUnitDescription string) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := t.Prepare("INSERT INTO OrgUnits (OrgUnitName, Description) VALUES (?, ?)")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return false, err
	}

	_, err = q.Exec(orgUnitName, orgUnitDescription)
	if err != nil {
		errPrintln("Cannot create org unit '" + orgUnitName + "': " + string(err.Error()))
		return false, err
	}

	t.Commit()

	return true, nil
}

func getOrgUnitStatus(orgUnit string) (bool, error) {
	t, err := DB.Begin()
	if err != nil {
		errPrintln("Could not start DB transaction: " + string(err.Error()))
		return false, err
	}

	q, err := DB.Prepare("SELECT * FROM OrgUnits WHERE OrgUnitName IS ?")
	if err != nil {
		errPrintln("Could not prepare DB query! " + string(err.Error()))
		return false, err
	}

	rr := OrgUnit{}
	err = q.QueryRow(orgUnit).Scan(
		&rr.Id,
		&rr.OrgUnitName,
		&rr.Description,
		&rr.CreationDate,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			warnPrintln("Encountered error when querying database: " + string(err.Error()))
			return false, err
		}
		return false, nil
	}

	t.Commit()

	return true, nil
}

func getRoleById(roleId int) (string, error) {
	rec, err := DB.Prepare("SELECT RoleName FROM Roles WHERE Id = ?")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return "", err
	}

	var roleName string
	err = rec.QueryRow(roleId).Scan(&roleName)
	if err != nil {
		if err == sql.ErrNoRows {
			warnPrintln("No such role found in DB: " + string(err.Error()))
			return "", nil
		}
		errPrintln("Cannot retrieve role from DB: " + string(err.Error()))
		return "", err
	}

	return roleName, nil
}

func getOrgUnitById(orgUnitId int) (string, error) {
	rec, err := DB.Prepare("SELECT OrgUnitName FROM OrgUnits WHERE Id = ?")
	if err != nil {
		errPrintln("Could not prepare the DB query: " + string(err.Error()))
		return "", err
	}

	var orgUnitName string
	err = rec.QueryRow(orgUnitId).Scan(&orgUnitName)
	if err != nil {
		if err == sql.ErrNoRows {
			warnPrintln("No such role found in DB: " + string(err.Error()))
			return "", nil
		}
		errPrintln("Cannot retrieve role from DB: " + string(err.Error()))
		return "", err
	}

	return orgUnitName, nil
}
