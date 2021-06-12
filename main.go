package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/unico/FeirasLivresAPI/api"
	"github.com/unico/FeirasLivresAPI/app"
	"github.com/unico/FeirasLivresAPI/store"
	config "github.com/unico/FeirasLivresAPI/utils"
)

func main() {

	config.Watch(func(c config.Config, quit chan bool) {
		ec := echo.New()
		//ec.Validator = validator.New()

		dbWriter := sqlx.MustConnect("mysql", "root:123@/feira?charset=utf8mb4,utf8\\u0026readTimeout=30s\\u0026writeTimeout=30s&parseTime=true")
		dbReader := sqlx.MustConnect("mysql", "root:123@/feira?charset=utf8mb4,utf8\\u0026readTimeout=30s\\u0026writeTimeout=30s&parseTime=true")

		// criação dos stores com a injeção do banco de escrita e leitura
		stores := store.New(store.Options{
			Writer: dbWriter,
			Reader: dbReader,
		})

		// criação dos serviços
		apps := app.New(app.Options{
			Stores: stores,
		})

		// registros dos handlers
		api.Register(api.Options{
			Group: ec.Group(""),
			Apps:  apps,
		})

		// função para fechar as conexões
		go func() {
			<-quit
			dbReader.Close()
			dbWriter.Close()
			ec.Close()
		}()

		go ec.Start(":7000")

		fmt.Println("Feiras Livres API Inicializado!")

	})

}
