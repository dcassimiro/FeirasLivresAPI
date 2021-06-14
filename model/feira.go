package model

import "time"

type FeiraRequest struct {
	ID_FEIRA   string  `json:"id_feira" validate:"required"`
	XLONG      string  `json:"long" validate:"required"`
	LAT        string  `json:"lat" validate:"required"`
	SETCENS    string  `json:"setcens" validate:"required"`
	AREAP      string  `json:"areap" validate:"required"`
	CODDIST    string  `json:"coddist" validate:"required"`
	DISTRITO   string  `json:"distrito" validate:"required"`
	CODSUBPREF string  `json:"codsubpref" validate:"required"`
	SUBPREFE   string  `json:"subprefe" validate:"required"`
	REGIAO5    string  `json:"regiao5" validate:"required"`
	REGIAO8    string  `json:"regiao8" validate:"required"`
	NOME_FEIRA string  `json:"nome_feira" validate:"required"`
	REGISTRO   string  `json:"registro" validate:"required"`
	LOGRADOURO *string `json:"logradouro"`
	NUMERO     *string `json:"numero"`
	BAIRRO     *string `json:"bairro"`
	REFERENCIA *string `json:"referencia"`
}

type Feira struct {
	ID         string     `json:"id" db:"ID"`
	ID_FEIRA   string     `json:"id_feira" db:"ID_FEIRA"`
	XLONG      string     `json:"long" db:"XLONG"`
	LAT        string     `json:"lat" db:"LAT"`
	SETCENS    string     `json:"setcens" db:"SETCENS"`
	AREAP      string     `json:"areap" db:"AREAP"`
	CODDIST    string     `json:"coddist" db:"CODDIST"`
	DISTRITO   string     `json:"distrito" db:"DISTRITO"`
	CODSUBPREF string     `json:"codsubpref" db:"CODSUBPREF"`
	SUBPREFE   string     `json:"subprefe" db:"SUBPREFE"`
	REGIAO5    string     `json:"regiao5" db:"REGIAO5"`
	REGIAO8    string     `json:"regiao8" db:"REGIAO8"`
	NOME_FEIRA string     `json:"nome_feira" db:"NOME_FEIRA"`
	REGISTRO   string     `json:"registro" db:"REGISTRO"`
	LOGRADOURO *string    `json:"logradouro" db:"LOGRADOURO"`
	NUMERO     *string    `json:"numero" db:"NUMERO"`
	BAIRRO     *string    `json:"bairro" db:"BAIRRO"`
	REFERENCIA *string    `json:"referencia" db:"REFERENCIA"`
	CREATED_AT time.Time  `json:"created_at" db:"CREATED_AT"`
	UPDATED_AT *time.Time `json:"updated_at" db:"UPDATED_AT"`
	DELETED_AT *time.Time `json:"deleted_at" db:"DELETED_AT"`
}
