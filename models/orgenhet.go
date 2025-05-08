package models

/*
 * Copyright (c) 2023-2025 Norwegian University of Science and Technology
 */

// Orgenhet is an organizational unit from DFØ/SAP API.
type Orgenhet struct {
	ID                 int64     `json:"id,omitempty"`
	OrgKortnavn        string    `json:"orgKortnavn,omitempty"`
	Organisasjonsnavn  string    `json:"organisasjonsnavn,omitempty"`
	Leder              []*Leder  `json:"leder,omitempty"`
	PDO                string    `json:"pdo,omitempty"`
	Type               string    `json:"type,omitempty"`
	LokasjonID         string    `json:"lokasjonId,omitempty"`
	OrgKostnadssted    string    `json:"orgKostnadssted,omitempty"`
	AuthWF             []*AuthWF `json:"authWf,omitempty"`
	OverordnOrgenhetID string    `json:"overordnOrgenhetId,omitempty"`
}

func (oe *Orgenhet) String() string {
	return toString(oe)
}

// Leder is the organizational leader for a given organizational unit.
type Leder struct {
	LederAnsattnr    string `json:"lederAnsattnr,omitempty"`
	LederBrukerident string `json:"lederBrukerident,omitempty"`
}

func (l *Leder) String() string {
	return toString(l)
}

// AuthWF is ...  I have no idea.
type AuthWF struct {
	Type          string `json:"type,omitempty"`
	Bruker        string `json:"bruker,omitempty"`
	KnytningStart string `json:"knytningStart,omitempty"`
	KnytningSlutt string `json:"knytningSlutt,omitempty"`
}

func (a *AuthWF) String() string {
	return toString(a)
}
