package feira

import (
	"context"
	"fmt"

	"github.com/unico/FeirasLivresAPI/model"
	"github.com/unico/FeirasLivresAPI/store"
)

type App interface {
	Create(ctx context.Context, feira model.Feira) (*model.Feira, error)
}

func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores *store.Container
}

func (s *appImpl) Create(ctx context.Context, feira model.Feira) (*model.Feira, error) {

	id, err := s.stores.Feira.Create(ctx, feira)
	if err != nil {
		fmt.Println(ctx, "app.feira.Create.Create", err.Error())
		return nil, err
	}

	data, err := s.stores.Feira.ReadOne(ctx, id)
	if err != nil {
		fmt.Println(ctx, "app.feira.Create.ReadOne", err.Error())
		return nil, err
	}

	return data, nil
}
