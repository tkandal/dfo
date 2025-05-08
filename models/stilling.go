package models

/*
 * Copyright (c) 2023-2025 Norwegian University of Science and Technology
 */

// StillingResponse is returned when fetching positions from DFO/SAP REST-API.
type StillingResponse struct {
	Stilling []*Stilling `json:"stilling,omitempty"`
}

func (sr *StillingResponse) String() string {
	return toString(sr)
}

// Stilling is a position from DFØ/SAP REST-API.
type Stilling struct {
	ID                int64                `json:"id,omitempty"`
	Stillingsnavn     string               `json:"stillingsnavn,omitempty"`
	Stillingskode     int64                `json:"stillingskode,omitempty"`
	KodeKnytningStart int64                `json:"kodeKnytningStart,omitempty"`
	KodeKnytningSlutt int64                `json:"kodeKnytningSlutt,omitempty"`
	Stillingstittel   string               `json:"stillingstittel,omitempty"`
	OrganisasjonID    int64                `json:"organisasjon_id,omitempty"`
	OrgKnytningStart  int64                `json:"orgKnytningStart,omitempty"`
	OrgKnytningSlutt  int64                `json:"orgKnytningSlutt,omitempty"`
	Yrkeskode         string               `json:"yrkeskode,omitempty"`
	Yrkeskodetekst    string               `json:"yrkeskodetekst,omitempty"`
	Bistilling        bool                 `json:"bistilling"`
	Innehaver         []*Innehaver         `json:"innehaver,omitempty"`
	Stillingskat      []*StillingsKategori `json:"stillingskat,omitempty"`
	Finansiering      []Finansiering       `json:"finansiering,omitempty"`
}

func (s *Stilling) String() string {
	return toString(s)
}

// Innehaver is a person that has a given position.
type Innehaver struct {
	InnehaverAnsattnr  string `json:"innehaverAnsattnr,omitempty"`
	InnehaverStartdato string `json:"innehaverStartdato,omitempty"`
	InnehaverSluttdato string `json:"innehaverSluttdato,omitempty"`
}

func (i *Innehaver) String() string {
	return toString(i)
}

// StillingsKategori is a position category from DFØ/SAP REST-API.
type StillingsKategori struct {
	StillingskatID        int64  `json:"stillingskatId,omitempty"`
	StillingskatStartdato string `json:"stillingskatStartdato,omitempty"`
	StillingskatSluttdato string `json:"stillingskatSluttdato,omitempty"`
	StillingskatBetegn    string `json:"stillingskatBetegn,omitempty"`
}

func (s *StillingsKategori) String() string {
	return toString(s)
}

// Finansiering contains information about the funding of a position.
type Finansiering struct {
	FinansieringstypeID     int64  `json:"finansieringstypeId,omitempty"`
	Finansieringskode       string `json:"finansieringskode,omitempty"`
	FinansieringStartdato   string `json:"finansieringStartdato,omitempty"`
	FinansieringSluttdato   string `json:"finansieringSluttdato,omitempty"`
	Finansieringsbetegnelse string `json:"finansieringsbetegnelse,omitempty"`
	Finansieringsprosent    string `json:"finansieringsprosent,omitempty"`
}

func (f *Finansiering) String() string {
	return toString(f)
}
