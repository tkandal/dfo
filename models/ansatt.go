package models

import (
	"bytes"
	"encoding/json"
	"time"
)

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// AnsattResponse is the struct that returned when fetching persons from DFO/SAP API.
type AnsattResponse struct {
	Ansatt []*Ansatt `json:"ansatt"`
}

func (ar *AnsattResponse) String() string {
	return toString(ar)
}

// Ansatt is an employee from DFØ/SAP API.
type Ansatt struct {
	ID                     string              `json:"id,omitempty"`
	BrukerIdent            string              `json:"brukerident,omitempty"`
	DFOBrukerIdent         string              `json:"dfoBrukerident,omitempty"`
	EksternIdent           string              `json:"eksternIdent,omitempty"`
	Fornavn                string              `json:"fornavn,omitempty"`
	Etternavn              string              `json:"etternavn,omitempty"`
	Tittel                 string              `json:"tittel,omitempty"`
	FNR                    string              `json:"fnr,omitempty"`
	AnnenID                []*AnnenID          `json:"annen_id,omitempty"`
	FDato                  string              `json:"fdato,omitempty"`
	Kjonn                  string              `json:"kjonn,omitempty"`
	SakArkivNr             string              `json:"sakarkivnr,omitempty"`
	Landkode               string              `json:"landkode,omitempty"`
	MedarbeiderGruppe      string              `json:"medarbeidergruppe,omitempty"`
	MedarbeideUndergruppe  string              `json:"medarbeiderundergruppe,omitempty"`
	Startdato              string              `json:"startdato,omitempty"`
	Sluttdato              string              `json:"sluttdato,omitempty"`
	Sluttarsak             string              `json:"sluttarsak,omitempty"`
	StillingID             int64               `json:"stillingId,omitempty"`
	Dellonnsprosent        string              `json:"dellonnsprosent,omitempty"`
	Bevilgning             string              `json:"bevilgning,omitempty"`
	Kostnadssted           string              `json:"kostnadssted,omitempty"`
	OrganisasjonID         int64               `json:"organisasjonId,omitempty"`
	JurBedriftsnummer      int64               `json:"jurBedriftsnummer,omitempty"`
	PDO                    string              `json:"pdo,omitempty"`
	Tilleggsstilling       []*Tilleggsstilling `json:"tilleggsstilling,omitempty"`
	Lederflagg             bool                `json:"lederflagg"`
	Portaltilgang          bool                `json:"portaltilgang"`
	Turnustilgang          bool                `json:"turnustilgang,omitempty"`
	Eksternbruker          bool                `json:"eksternbruker"`
	Epost                  string              `json:"epost,omitempty"`
	Tjenestetelefon        string              `json:"tjenestetelefon,omitempty"`
	PrivatTelefonnummer    string              `json:"privatTelefonnummer,omitempty"`
	Telefonnummer          string              `json:"telefonnummer,omitempty"`
	Mobilnummer            string              `json:"mobilnummer,omitempty"`
	MobilPrivat            string              `json:"mobilPrivat,omitempty"`
	PrivatTlfUtland        string              `json:"privatTlfUtland,omitempty"`
	TelefonJobb            string              `json:"telefonJobb,omitempty"`
	TelefonPrivat          string              `json:"telefonPrivat,omitempty"`
	PrivatPostadresse      string              `json:"privatPostadresse,omitempty"`
	PrivatPostnr           string              `json:"privatPostnr,omitempty"`
	PrivatPoststed         string              `json:"privatPoststed,omitempty"`
	EndretDato             string              `json:"endretDato,omitempty"`
	EndretAv               string              `json:"endretAv,omitempty"`
	ReservasjonPublisering bool                `json:"reservasjonPublisering"`
	Hjemmelkode            string              `json:"hjemmelKode,omitempty"`
	HjemmelTekst           string              `json:"hjemmelTekst,omitempty"`
	EndretInfotype         string              `json:"endretInfotype,omitempty"`
	EndretKlokkeslett      string              `json:"endretKlokkeslett,omitempty"`
	PrivatEpost            string              `json:"privatEpost,omitempty"`
}

func (a *Ansatt) String() string {
	return toString(a)
}

// AnnenID is an alternative identity to FNR, f.ex. passport number.
type AnnenID struct {
	IDType        string `json:"idType,omitempty"`
	IDBeskrivelse string `json:"idBeskrivelse,omitempty"`
	IDNr          string `json:"idNr,omitempty"`
	IDStartdato   string `json:"idStartdato,omitempty"`
	IDSluttdato   string `json:"idSluttdato,omitempty"`
	IDLand        string `json:"idLand,omitempty"`
}

func (a *AnnenID) String() string {
	return toString(a)
}

// Tilleggsstilling is an extra employment in addition the main employment.
type Tilleggsstilling struct {
	StillingID      int64     `json:"stillingId,omitempty"`
	Startdato       time.Time `json:"startdato,omitempty"`
	Sluttdato       time.Time `json:"sluttdato,omitempty"`
	Dellonnsprosent float64   `json:"dellonnsprosent,omitempty"`
	EkstraStilling  string    `json:"ekstraStilling,omitempty"`
}

func (t *Tilleggsstilling) String() string {
	return toString(t)
}

func toString(v interface{}) string {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(v)
	return b.String()
}
