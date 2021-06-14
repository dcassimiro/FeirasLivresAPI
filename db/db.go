package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/unico/FeirasLivresAPI/logger"
	"github.com/unico/FeirasLivresAPI/model"
)

func CreateDB() {

	url := model.Url()
	db := sqlx.MustConnect("mysql", url)

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS feira")
	if err != nil {
		logger.L.Println(time.Now(), "CreateDB:", err.Error())
		panic(err)

	}

	_, err = db.Exec("USE feira")
	if err != nil {
		logger.L.Println(time.Now(), "CreateDB:", err.Error())
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS feira(
		ID         INT(11) NOT NULL AUTO_INCREMENT
	   ,ID_FEIRA   VARCHAR(200) NOT NULL
	   ,XLONG       VARCHAR(200) NOT NULL
	   ,LAT        VARCHAR(200) NOT NULL
	   ,SETCENS    VARCHAR(200) NOT NULL
	   ,AREAP      VARCHAR(200) NOT NULL
	   ,CODDIST    VARCHAR(200) NOT NULL
	   ,DISTRITO   VARCHAR(200) NOT NULL
	   ,CODSUBPREF VARCHAR(200) NOT NULL
	   ,SUBPREFE   VARCHAR(200) NOT NULL
	   ,REGIAO5    VARCHAR(200) NOT NULL
	   ,REGIAO8    VARCHAR(200) NOT NULL
	   ,NOME_FEIRA VARCHAR(200) NOT NULL
	   ,REGISTRO   VARCHAR(200) NOT NULL
	   ,LOGRADOURO VARCHAR(200) NOT NULL
	   ,NUMERO     VARCHAR(200)
	   ,BAIRRO     VARCHAR(200)
	   ,REFERENCIA VARCHAR(200)
	   ,CREATED_AT DATETIME DEFAULT CURRENT_TIMESTAMP
	   ,UPDATED_AT DATETIME DEFAULT NULL
	   ,DELETED_AT DATETIME DEFAULT NULL
	   ,PRIMARY KEY (ID)
	 )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`)
	if err != nil {
		logger.L.Println(time.Now(), "CreateDB:", err.Error())
		panic(err)
	}

}
