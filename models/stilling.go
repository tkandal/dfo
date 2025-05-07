package models

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Stilling is a position from DFØ/SAP API.
type Stilling struct {
	ID              int64                `json:"id,omitempty"`
	Stillingsnavn   string               `json:"stillingsnavn,omitempty"`
	Stillingskode   int64                `json:"stillingskode,omitempty"`
	Stillingstittel string               `json:"stillingstittel,omitempty"`
	OrganisasjonID  int64                `json:"organisasjon_id,omitempty"`
	Yrkeskode       string               `json:"yrkeskode,omitempty"`
	Yrkeskodetekst  string               `json:"yrkeskodetekst,omitempty"`
	Innehaver       []*Innehaver         `json:"innehaver,omitempty"`
	Stillingskat    []*StillingsKategori `json:"stillingskat,omitempty"`
}

func (s *Stilling) String() string {
	return toString(s)
}

// Innehaver is a person that has a given position.
type Innehaver struct {
	InnehaverAnsattnr  string    `json:"innehaverAnsattnr,omitempty"`
	InnehaverStartdato string `json:"innehaverStartdato,omitempty"`
	InnehaverSluttdato string `json:"innehaverSluttdato,omitempty"`
}

func (i *Innehaver) String() string {
	return toString(i)
}

// StillingsKategori is a position category from DFØ/SAP API.
type StillingsKategori struct {
	StillingskatID        int64     `json:"stillingskatId,omitempty"`
	StillingskatStartdato string `json:"stillingskatStartdato,omitempty"`
	StillingskatSluttdato string `json:"stillingskatSluttdato,omitempty"`
	StillingskatBetegn    string    `json:"stillingskatBetegn,omitempty"`
}

func (s *StillingsKategori) String() string {
	return toString(s)
}
