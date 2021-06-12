package app

import (
	"github.com/unico/FeirasLivresAPI/app/feira"
	"github.com/unico/FeirasLivresAPI/store"
)

type Container struct {
	Feira feira.App
}

type Options struct {
	Stores *store.Container
}

func New(opts Options) *Container {

	container := &Container{
		Feira: feira.NewApp(opts.Stores),
	}

	return container
}
