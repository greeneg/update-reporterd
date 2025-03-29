package main

type UpdateReporter struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

type credentials struct {
	Program UpdateReporter `json:"updatereporter"`
}

type osReleaseStruct struct {
	Id      string
	Version string
}

type UpdateStruct struct {
	Updates     []Update `json:"updates"`
	UpdateCount int      `json:"updateCount"`
	FQDN        string   `json:"fqdn"`
	OsFamily    string   `json:"osFamily"`
	OsId        string   `json:"osId"`
	OsVersion   string   `json:"osVersion"`
	HostArch    string   `json:"hostArchitecture"`
}

type Update struct {
	Kind       string `json:"kind" xml:"kind,attr"`
	Name       string `json:"name" xml:"name,attr"`
	Version    string `json:"version" xml:"edition,attr"`
	Arch       string `json:"arch" xml:"arch,attr"`
	OldVersion string `json:"oldVersion" xml:"edition-old,attr"`
	Summary    string `json:"summary" xml:"summary"`
}

type UpdateList struct {
	Updates []Update `xml:"update"`
}

type UpdateStatus struct {
	UpdateList UpdateList `xml:"update-list"`
}

type Stream struct {
	UpdateStatus UpdateStatus `xml:"update-status"`
}
