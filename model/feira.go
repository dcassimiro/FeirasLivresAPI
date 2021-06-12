package model

import "time"

type FeiraRequest struct {
	IDFEIRA     string `json:"idfeira"`
	XLONG       string `json:"long"`
	LAT         string `json:"lat"`
	SETCENS     string `json:"setcens"`
	AREAP       string `json:"areap"`
	CODDIST     string `json:"coddist"`
	DISTRITO    string `json:"distrito"`
	CODSUBPREF  string `json:"codsubpref"`
	SUBPREFE    string `json:"subprefe"`
	REGIAO5T    string `json:"regiao5t"`
	REGIAO8T    string `json:"regiao8t"`
	NOMEFEIRA   string `json:"nomefeira"`
	REGISTROT   string `json:"registrot"`
	LOGRADOUROT string `json:"logradourot"`
	NUMEROT     string `json:"numerot"`
	BAIRROT     string `json:"bairrot"`
	REFERENCIAT string `json:"referenciat"`
}

type Feira struct {
	Id          string     `json:"id"`
	IDFEIRA     string     `json:"idfeira"`
	XLONG       string     `json:"long"`
	LAT         string     `json:"lat"`
	SETCENS     string     `json:"setcens"`
	AREAP       string     `json:"areap"`
	CODDIST     string     `json:"coddist"`
	DISTRITO    string     `json:"distrito"`
	CODSUBPREF  string     `json:"codsubpref"`
	SUBPREFE    string     `json:"subprefe"`
	REGIAO5T    string     `json:"regiao5t"`
	REGIAO8T    string     `json:"regiao8t"`
	NOMEFEIRA   string     `json:"nomefeira"`
	REGISTROT   string     `json:"registrot"`
	LOGRADOUROT string     `json:"logradourot"`
	NUMEROT     string     `json:"numerot"`
	BAIRROT     string     `json:"bairrot"`
	REFERENCIAT string     `json:"referenciat"`
	CreatedAt   time.Time  `json:"createdat"`
	UpdatedAt   *time.Time `json:"updatedat"`
	DeletedAt   *time.Time `json:"deletedat"`
}
