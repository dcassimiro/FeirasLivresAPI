package feira

import (
	"context"
	"time"

	"github.com/unico/FeirasLivresAPI/logger"
	"github.com/unico/FeirasLivresAPI/model"
	"github.com/unico/FeirasLivresAPI/store"
)

type App interface {
	Create(ctx context.Context, feira model.FeiraRequest) (*model.Feira, error)
	Update(ctx context.Context, id string, feira model.FeiraRequest) (*model.Feira, error)
	ReadOne(ctx context.Context, id string) (*model.Feira, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, distrito string) ([]*model.Feira, error)
}

func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores *store.Container
}

func (s *appImpl) Create(ctx context.Context, feira model.FeiraRequest) (*model.Feira, error) {

	id, err := s.stores.Feira.Create(ctx, feira)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Create.Create", err.Error())
		logger.L.Println(time.Now(), "app.feira.Create.Create", err.Error())
		return nil, err
	}

	data, err := s.stores.Feira.ReadOne(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Create.ReadOne", err.Error())
		logger.L.Println(time.Now(), "app.feira.Create.ReadOne", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *appImpl) Update(ctx context.Context, id string, feira model.FeiraRequest) (*model.Feira, error) {
	_, err := s.stores.Feira.Update(ctx, id, feira)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Update.Update", err.Error())
		logger.L.Println(time.Now(), "app.feira.Update.Update", err.Error())
		return nil, err
	}

	data, err := s.stores.Feira.ReadOne(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Update.ReadOne", err.Error())
		logger.L.Println(time.Now(), "app.feira.Update.ReadOne", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *appImpl) ReadOne(ctx context.Context, id string) (*model.Feira, error) {
	feira, err := s.stores.Feira.ReadOne(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.ReadOne.ReadOne", err.Error())
		logger.L.Println(time.Now(), "app.feira.ReadOne.ReadOne", err.Error())
		return nil, err
	}

	return feira, nil
}

func (s *appImpl) Delete(ctx context.Context, id string) error {
	err := s.stores.Feira.Delete(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Delete.Delete", "não consegui deletar a feira: "+id, err.Error())
		logger.L.Println(time.Now(), "app.feira.Delete.Delete", err.Error())
		return err
	}

	return nil
}

func (s *appImpl) Search(ctx context.Context, distrito string) ([]*model.Feira, error) {
	feiras, err := s.stores.Feira.Search(ctx, distrito)
	if err != nil {
		logger.ErrorContext(ctx, "app.feira.Search.Search", err.Error())
		logger.L.Println(time.Now(), "app.feira.Delete.Delete", err.Error())
		return nil, err
	}

	return feiras, nil
}
