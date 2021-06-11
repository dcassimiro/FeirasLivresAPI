package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"

	"github.com/micro/go-micro/logger"
	"github.com/unico/FeirasLivresAPI/api"
	"github.com/unico/FeirasLivresAPI/app"
	"github.com/unico/FeirasLivresAPI/store"
	config "github.com/unico/FeirasLivresAPI/utils"
)

func main() {

	config.Watch(func(c config.Config, quit chan bool) {
		ec := echo.New()

		dbWriter := sqlx.MustConnect("mysql", "root:123@/feira")
		dbReader := sqlx.MustConnect("mysql", "root:123@/feira")

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

		go ec.Start(c.GetString("server.port"))

		logger.Info("API Feiras Livres inicializado!")

	})

}
