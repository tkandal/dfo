package models

import (
	"bytes"
	"encoding/json"
	"regexp"
)

/*
 * Copyright (c) 2023-2026 Norwegian University of Science and Technology
 */

var nin *regexp.Regexp

func init() {
	nin = regexp.MustCompile(`^\d{10,11}$`)
}

// AnsattResponse is returned when fetching persons from DFO/SAP REST-API.
type AnsattResponse struct {
	Ansatt []*Ansatt `json:"ansatt"`
}

func (ar *AnsattResponse) String() string {
	return toString(ar)
}

// Ansatt is an employee from DFØ/SAP REST-API.
type Ansatt struct {
	ID                     string              `json:"id"`
	BrukerIdent            string              `json:"brukerident"`
	DFOBrukerIdent         string              `json:"dfoBrukerident"`
	EksternIdent           string              `json:"eksternIdent"`
	Tittel                 string              `json:"tittel"`
	Fornavn                string              `json:"fornavn"`
	Etternavn              string              `json:"etternavn"`
	FNR                    string              `json:"fnr"`
	AnnenID                []*AnnenID          `json:"annenId"`
	FDato                  string              `json:"fdato"`
	Kjonn                  string              `json:"kjonn"`
	SakArkivNr             string              `json:"sakarkivnr"`
	Landkode               string              `json:"landkode"`
	MedarbeiderGruppe      string              `json:"medarbeidergruppe"`
	MedarbeideUndergruppe  string              `json:"medarbeiderundergruppe"`
	Startdato              string              `json:"startdato"`
	Sluttdato              string              `json:"sluttdato"`
	Sluttarsak             string              `json:"sluttarsak"`
	StillingID             int64               `json:"stillingId"`
	Hjemmelkode            string              `json:"hjemmelKode"`
	HjemmelTekst           string              `json:"hjemmelTekst"`
	Dellonnsprosent        string              `json:"dellonnsprosent"`
	Kostnadssted           string              `json:"kostnadssted"`
	OrganisasjonID         string              `json:"organisasjonId"`
	JurBedriftsnummer      int64               `json:"jurBedriftsnummer"`
	PDO                    string              `json:"pdo"`
	Tilleggsstilling       []*Tilleggsstilling `json:"tilleggsstilling"`
	Lederflagg             bool                `json:"lederflagg"`
	Portaltilgang          bool                `json:"portaltilgang"`
	Turnustilgang          bool                `json:"turnustilgang"`
	Eksternbruker          bool                `json:"eksternbruker"`
	ReservasjonPublisering bool                `json:"reservasjonPublisering"`
	Epost                  string              `json:"epost"`
	Tjenestetelefon        string              `json:"tjenestetelefon"`
	PrivatTelefonnummer    string              `json:"privatTelefonnummer"`
	Telefonnummer          string              `json:"telefonnummer"`
	Mobilnummer            string              `json:"mobilnummer"`
	MobilPrivat            string              `json:"mobilPrivat"`
	PrivatTlfUtland        string              `json:"privatTlfUtland"`
	TelefonJobb            string              `json:"telefonJobb"`
	TelefonPrivat          string              `json:"telefonPrivat"`
	PrivatEpost            string              `json:"privatEpost"`
	PrivatPostadresse      string              `json:"privatPostadresse"`
	PrivatPostnr           string              `json:"privatPostnr"`
	PrivatPoststed         string              `json:"privatPoststed"`
	EndretDato             string              `json:"endretDato"`
	EndretKlokkeslett      string              `json:"endretKlokkeslett"`
	EndretInfotype         string              `json:"endretInfotype"`
	EndretAv               string              `json:"endretAv"`
}

// HasNIN checks of this emplyee has NIN (National Identity Number).
func (a *Ansatt) HasNIN() bool {
	return nin.Match([]byte(a.FNR))
}

func (a *Ansatt) String() string {
	return toString(a)
}

// AnnenID is an alternative identity to FNR, f.ex. passport number, national ID, etc.
type AnnenID struct {
	IDType        string `json:"idType"`
	IDBeskrivelse string `json:"idBeskrivelse"`
	IDNr          string `json:"idNr"`
	IDStartdato   string `json:"idStartdato"`
	IDSluttdato   string `json:"idSluttdato"`
	IDLand        string `json:"idLand"`
}

func (a *AnnenID) String() string {
	return toString(a)
}

// Tilleggsstilling is extra employment in addition the main employment.
type Tilleggsstilling struct {
	StillingID      int64  `json:"stillingId"`
	Startdato       string `json:"startdato"`
	Sluttdato       string `json:"sluttdato"`
	Dellonnsprosent string `json:"dellonnsprosent"`
	EkstraStilling  string `json:"ekstraStilling"`
}

func (t *Tilleggsstilling) String() string {
	return toString(t)
}

func toString(v any) string {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(v)
	return b.String()
}
