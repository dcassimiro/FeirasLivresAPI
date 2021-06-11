package model

type Feira struct {
	ID          string `json:"id" db:"id"`
	LONG        string `json:"" db:""`
	LAT         string `json:"" db:""`
	SETCENS     string `json:"" db:""`
	AREAP       string `json:"" db:""`
	CODDIST     string `json:"" db:""`
	DISTRITO    string `json:"distrito" db:"distrito"`
	CODSUBPREF  string `json:"" db:""`
	SUBPREFE    string `json:"" db:""`
	REGIAO5T    string `json:"" db:""`
	REGIAO8T    string `json:"" db:""`
	NOMEFEIRA   string `json:"" db:""`
	REGISTROT   string `json:"" db:""`
	LOGRADOUROT string `json:"" db:""`
	NUMEROT     string `json:"" db:""`
	BAIRROT     string `json:"" db:""`
	REFERENCIAT string `json:"" db:""`
}
