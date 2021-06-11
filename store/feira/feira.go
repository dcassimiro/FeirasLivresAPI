package feira

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/unico/FeirasLivresAPI/model"
)

type Store interface {
	Create(ctx context.Context, feira model.Feira) (string, error)
	// Update(ctx context.Context, id string, feira model.Feira) (string, error)
	ReadOne(ctx context.Context, id string) (*model.Feira, error)
	// ReadAll(ctx context.Context) ([]*model.Feira, error)
	// Delete(ctx context.Context, id string) error
}

// NewStore cria uma nova instancia do repositorio de feira
func NewStore(writer, reader *sqlx.DB) Store {
	return &storeImpl{writer, reader}
}

type storeImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func (r *storeImpl) Create(ctx context.Context, feira model.Feira) (string, error) {

	result, err := r.writer.ExecContext(ctx, `
		INSERT INTO feiras (distrito)
		VALUES (?);
	`, feira.DISTRITO)
	if err != nil {
		fmt.Println(ctx, "store.feira.Create", err.Error())
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(ctx, "store.feira.Create", err.Error())
	}

	return strconv.FormatInt(id, 10), nil
}

func (r *storeImpl) ReadOne(ctx context.Context, id string) (*model.Feira, error) {
	feira := new(model.Feira)
	err := r.writer.GetContext(ctx, feira, `
		SELECT
			id,
			distrito,
		FROM feira 
		WHERE
			id = ? AND
	`, id)
	if err != nil {
		fmt.Println(ctx, "store.feira.ReadOne", err.Error())
		return nil, err
	}

	return feira, nil
}
