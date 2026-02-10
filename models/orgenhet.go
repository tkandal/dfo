package models

/*
 * Copyright (c) 2023-2026 Norwegian University of Science and Technology
 */

// Orgenhet is an organizational unit from DFØ/SAP REST-API.
type Orgenhet struct {
	ID                 int64     `json:"id"`
	OrgKortnavn        string    `json:"orgKortnavn"`
	Organisasjonsnavn  string    `json:"organisasjonsnavn"`
	Leder              []*Leder  `json:"leder"`
	PDO                string    `json:"pdo"`
	Type               string    `json:"type"`
	OrgKostnadssted    string    `json:"orgKostnadssted"`
	AuthWF             []*AuthWF `json:"authWf"`
	LokasjonID         string    `json:"lokasjonId"`
	DBHKode            string    `json:"dbhKode"`
	DBHBetegnelse      string    `json:"dbhBetegnelse"`
	OverordnOrgenhetID string    `json:"overordnOrgenhetId"`
}

func (oe *Orgenhet) String() string {
	return toString(oe)
}

// Leder is the organizational leader for a given organizational unit.
type Leder struct {
	LederAnsattnr    string `json:"lederAnsattnr"`
	LederBrukerident string `json:"lederBrukerident"`
}

func (l *Leder) String() string {
	return toString(l)
}

// AuthWF is ...  I have no idea.
type AuthWF struct {
	Type          string `json:"type"`
	Bruker        string `json:"bruker"`
	KnytningStart string `json:"knytningStart"`
	KnytningSlutt string `json:"knytningSlutt"`
}

func (a *AuthWF) String() string {
	return toString(a)
}
