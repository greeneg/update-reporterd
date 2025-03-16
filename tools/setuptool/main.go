package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pborman/getopt/v2"
	"golang.org/x/term"
)

var VERSION string = "0.0.1"

var appFullPath, _ = os.Executable()
var app = path.Base(appFullPath)

// setup the global flags

var (
	dbFile             string
	account            string
	fullName           string
	orgUnitName        string
	orgUnitDescription string
	role               string
	roleDescription    string
	optHelp            = getopt.BoolLong("help", 'h', "This help message")
	optVersion         = getopt.BoolLong("version", 'v', "Show the version")
)

func showHelp() {
	println(app + " - Setup tool for Update Reporter Daemon")
	dividerLine := strings.Repeat("=", 43)
	println(dividerLine)
	println("Add and configure roles or accounts for the Update Reporter Daemon\n")
	println("OPTIONS:")
	println("   -d|--database-file FILENAME_PATH       REQUIRED: The full or relative path")
	println("                                          to the database file")
	println("   -a|--account ACCOUNT_NAME              OPTIONAL: The account to create")
	println("   -o|--orgunit ORGUNIT_NAME              OPTIONAL: The orgunit to create")
	println("   -O|--orgunit-description ORGUNIT_DESC  CONDITIONALLY OPTIONAL: If the")
	println("                                          orgunit flag is set, this is required.")
	println("                                          This should be the description for")
	println("                                          the orgunit to be registered with the")
	println("                                          system.")
	println("   -r|--role ROLE_NAME                    OPTIONAL: The role to create")
	println("   -f|--fullname QUOTED_FULLNAME          CONDITIONALLY OPTIONAL: If the")
	println("                                          account flag is set, this is required.")
	println("                                          This should be the full name or")
	println("                                          description for the account to be")
	println("                                          registered with the system.")
	println("   -D|--role-description ROLE_DESCRIPTION CONDITIONALLY OPTIONAL: If the")
	println("                                          role flag is set, this is required.")
	println("                                          This should be the description for")
	println("                                          the role to be registered with the")
	println("                                          system.")
	println("")
	println("Author: Gary L. Greene, Jr. <greeneg@tolharadys.net>")
	println("License: Apache Public License, v2")
	showVersion()
}

func showVersion() {
	println("version: " + VERSION)
}

func init() {
	getopt.FlagLong(&dbFile, "database-file", 'd', "The full path to the database file")
	getopt.FlagLong(&account, "account", 'a', "The account to add to the system")
	getopt.FlagLong(&fullName, "fullname", 'f', "The full name to associate with the account")
	getopt.FlagLong(&orgUnitName, "orgunit", 'o', "The orgunit to add to the system")
	getopt.FlagLong(&orgUnitDescription, "orgunit-description", 'O', "The description of the orgunit to process")
	getopt.FlagLong(&role, "role", 'r', "The role to add to the system")
	getopt.FlagLong(&roleDescription, "role-description", 'D', "The description of the role to process")
}

func main() {
	getopt.Parse()

	if *optHelp {
		showHelp()
		os.Exit(0)
	}

	if *optVersion {
		showVersion()
		os.Exit(0)
	}

	println("Starting setuptool... \n")

	// first, setup our DB connection
	if dbFile != "" {
		println("Database file: " + dbFile)
		ConnectDatabase(dbFile)
		println("Database connection completed")
	} else {
		errPrintln("Database file must be defined.")
		showHelp()
		os.Exit(1)
	}

	// do we need to process an account the user passed in?
	if account != "" {
		println("Account: " + account)
		if fullName != "" {
			println("Account fullname: " + fullName)
		} else {
			errPrintln("Account must have a full name")
			showHelp()
			os.Exit(1)
		}
	}

	// How about processing roles?
	if role != "" {
		println("Role: " + role)
		if roleDescription != "" {
			println("Role description: " + roleDescription)
		} else {
			errPrintln("Role must have a role description")
			showHelp()
			os.Exit(1)
		}
	}

	var orgunit OrgUnit
	// We're working on the built-in orgunit, 'SYSTEM'
	//
	// first, does the SYSTEM orgunit already exist?
	biOrgUnitState, err := getOrgUnitStatus("SYSTEM")
	if err != nil && err != sql.ErrNoRows {
		errPrintln("Encountered error when checking orgunit status: " + string(err.Error()))
		os.Exit(1)
	}
	if !biOrgUnitState {
		println("Creating orgunit 'SYSTEM'")
		status, err := createOrgUnit("SYSTEM", "Built-in system orgunit")
		if err != nil {
			errPrintln("Encountered error when creating orgunit: " + string(err.Error()))
			os.Exit(1)
		}
		if status {
			orgunit, err = getOrgUnitByName("SYSTEM")
			if err != nil {
				errPrintln("Encountered error when retrieving orgunit 'SYSTEM'")
				os.Exit(1)
			}
			orgunitStr, err := json.Marshal(orgunit)
			if err != nil {
				errPrintln("Encountered error when converting struct to JSON: " + string(err.Error()))
				os.Exit(1)
			}
			println("orgunit 'SYSTEM' created: " + string(orgunitStr))
		}
	} else {
		println("Built-in orgunit 'SYSTEM' already exists. Continuing")
		orgunit, _ = getOrgUnitByName("SYSTEM")
	}

	var roleRecord Role
	// We're working on the built-in account, 'admin' and role 'administrators'
	//
	// first, does the administrators role already exist?
	biRoleState, err := getRoleStatus("administrators")
	if err != nil && err != sql.ErrNoRows {
		errPrintln("Encountered error when checking role status: " + string(err.Error()))
		os.Exit(1)
	}
	if !biRoleState {
		println("Creating role 'administrators'")
		status, err := createRole("administrators", "Accounts that have full administrative rights to the system")
		if err != nil {
			errPrintln("Encountered error when creating role: " + string(err.Error()))
			os.Exit(1)
		}
		if status {
			roleRecord, err := getRoleByName("administrators")
			if err != nil {
				errPrintln("Encountered error when retrieving role 'administrators'")
				os.Exit(1)
			}
			roleRecordStr, err := json.Marshal(roleRecord)
			if err != nil {
				errPrintln("Encountered error when converting struct to JSON: " + string(err.Error()))
				os.Exit(1)
			}
			println("role 'administrators' created: " + string(roleRecordStr))
		}
	} else {
		println("Built-in role 'administrators' already exists. Continuing")
		roleRecord, _ = getRoleByName("administrators")
	}

	// now handle our admin user
	biAdminAccountState, err := getAccountStatus("admin")
	if err != nil && err != sql.ErrNoRows {
		errPrintln("Encountered error when checking account status: " + string(err.Error()))
		os.Exit(1)
	}
	if !biAdminAccountState {
		fmt.Print("Enter new password: ")
		input, _ := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Print("\nRe-enter passphrase: ")
		input2, _ := term.ReadPassword(int(os.Stdin.Fd()))
		println("")
		if strings.Compare(string(input), string(input2)) != 0 {
			errPrintln("Password does not match. Exiting")
			os.Exit(1)
		}

		accountRecord, err := createAccount("admin", "System Administrator", orgunit.Id, roleRecord.Id, string(input))
		if err != nil {
			errPrintln("Encountered error when creating account: " + string(err.Error()))
			os.Exit(1)
		}
		accountRecordStr, err := json.Marshal(accountRecord)
		if err != nil {
			errPrintln("Encountered error when converting struct to JSON: " + string(err.Error()))
			os.Exit(1)
		}
		println("account 'admin' created: " + string(accountRecordStr))
	} else {
		println("Built-in account 'admin' already exists. Continuing\n")
		userRecord, _ := getAccountByName("admin")
		println("Account: " + userRecord.UserName)
		println("Full Name: " + userRecord.FullName)
		println("Status: " + userRecord.Status)
		orgUnitName, _ := getOrgUnitById(userRecord.OrgUnitId)
		println("OrgUnit: " + orgUnitName)
		roleName, _ := getRoleById(userRecord.RoleId)
		println("Role: " + roleName)
	}
}
