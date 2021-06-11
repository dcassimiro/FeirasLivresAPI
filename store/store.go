package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/unico/FeirasLivresAPI/store/feira"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Feira feira.Store
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		Feira: feira.NewStore(opts.Writer, opts.Reader),
	}
}
