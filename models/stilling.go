package models

/*
 * Copyright (c) 2023-2026 Norwegian University of Science and Technology
 */

// StillingResponse is returned when fetching positions from DFO/SAP REST-API.
type StillingResponse struct {
	Stilling []*Stilling `json:"stilling"`
}

func (sr *StillingResponse) String() string {
	return toString(sr)
}

// Stilling is a position from DFØ/SAP REST-API.
type Stilling struct {
	ID                int64                `json:"id"`
	Stillingsnavn     string               `json:"stillingsnavn"`
	Stillingskode     int64                `json:"stillingskode"`
	KodeKnytningStart int64                `json:"kodeKnytningStart"`
	KodeKnytningSlutt int64                `json:"kodeKnytningSlutt"`
	Stillingstittel   string               `json:"stillingstittel"`
	OrganisasjonID    int64                `json:"organisasjon_id"`
	OrgKnytningStart  int64                `json:"orgKnytningStart"`
	OrgKnytningSlutt  int64                `json:"orgKnytningSlutt"`
	Yrkeskode         string               `json:"yrkeskode"`
	Yrkeskodetekst    string               `json:"yrkeskodetekst"`
	Bistilling        bool                 `json:"bistilling"`
	Innehaver         []*Innehaver         `json:"innehaver"`
	Stillingskat      []*StillingsKategori `json:"stillingskat"`
	Finansiering      []Finansiering       `json:"finansiering"`
}

func (s *Stilling) String() string {
	return toString(s)
}

// Innehaver is a person that has a given position.
type Innehaver struct {
	InnehaverAnsattnr  string `json:"innehaverAnsattnr"`
	InnehaverStartdato string `json:"innehaverStartdato"`
	InnehaverSluttdato string `json:"innehaverSluttdato"`
}

func (i *Innehaver) String() string {
	return toString(i)
}

// StillingsKategori is a position category from DFØ/SAP REST-API.
type StillingsKategori struct {
	StillingskatID        int64  `json:"stillingskatId"`
	StillingskatStartdato string `json:"stillingskatStartdato"`
	StillingskatSluttdato string `json:"stillingskatSluttdato"`
	StillingskatBetegn    string `json:"stillingskatBetegn"`
}

func (s *StillingsKategori) String() string {
	return toString(s)
}

// Finansiering contains information about the funding of a position.
type Finansiering struct {
	FinansieringstypeID     int64  `json:"finansieringstypeId"`
	Finansieringskode       string `json:"finansieringskode"`
	FinansieringStartdato   string `json:"finansieringStartdato"`
	FinansieringSluttdato   string `json:"finansieringSluttdato"`
	Finansieringsbetegnelse string `json:"finansieringsbetegnelse"`
	Finansieringsprosent    string `json:"finansieringsprosent"`
}

func (f *Finansiering) String() string {
	return toString(f)
}
