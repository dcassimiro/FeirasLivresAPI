package feira

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/unico/FeirasLivresAPI/fverr"
	"github.com/unico/FeirasLivresAPI/logger"
	"github.com/unico/FeirasLivresAPI/model"
)

type Store interface {
	Create(ctx context.Context, feira model.FeiraRequest) (string, error)
	Update(ctx context.Context, id string, feira model.FeiraRequest) (string, error)
	ReadOne(ctx context.Context, id string) (*model.Feira, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, distrito string) ([]*model.Feira, error)
}

// NewStore cria uma nova instancia do repositorio de feira
func NewStore(writer, reader *sqlx.DB) Store {
	return &storeImpl{writer, reader}
}

type storeImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func (r *storeImpl) Create(ctx context.Context, feira model.FeiraRequest) (string, error) {

	result, err := r.writer.ExecContext(ctx, `
		INSERT INTO feira (ID_FEIRA, XLONG, LAT, SETCENS, AREAP, CODDIST, DISTRITO, CODSUBPREF, SUBPREFE, REGIAO5, REGIAO8, NOME_FEIRA, REGISTRO, LOGRADOURO, NUMERO, BAIRRO, REFERENCIA)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`, feira.ID_FEIRA, feira.XLONG, feira.LAT, feira.SETCENS, feira.AREAP, feira.CODDIST, feira.DISTRITO, feira.CODSUBPREF, feira.SUBPREFE, feira.REGIAO5, feira.REGIAO8, feira.NOME_FEIRA, feira.REGISTRO, feira.LOGRADOURO, feira.NUMERO, feira.BAIRRO, feira.REFERENCIA)
	if err != nil {
		logger.ErrorContext(ctx, "store.feira.Create", err.Error())
		logger.L.Println(time.Now(), "não foi possível criar uma nova feira:", err.Error())
		return "", fverr.New(http.StatusInternalServerError, "não foi possível criar uma nova feira", nil)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.ErrorContext(ctx, "store.feira.Create", err.Error())
		logger.L.Println(time.Now(), "store.feira.Create:", err.Error())
	}

	return strconv.FormatInt(id, 10), nil
}

func (r *storeImpl) ReadOne(ctx context.Context, id string) (*model.Feira, error) {
	feira := new(model.Feira)
	err := r.writer.GetContext(ctx, feira, `
		SELECT
			ID,
			ID_FEIRA,
			XLONG,
			LAT,
			SETCENS,
			AREAP,
			CODDIST,
			DISTRITO,
			CODSUBPREF,
			SUBPREFE,
			REGIAO5,
			REGIAO8,
			NOME_FEIRA,
			REGISTRO,
			LOGRADOURO,
			NUMERO,
			BAIRRO,
			REFERENCIA,
			CREATED_AT,
			UPDATED_AT,
			DELETED_AT
		FROM feira 
		WHERE
			ID = ?;
	`, id)
	if err != nil {
		logger.ErrorContext(ctx, "store.feira.ReadOne", err.Error())
		logger.L.Println(time.Now(), "store.feira.ReadOne:", err.Error())
		return nil, err
	}

	return feira, nil
}

func (r *storeImpl) Update(ctx context.Context, id string, feira model.FeiraRequest) (string, error) {
	_, err := r.writer.ExecContext(ctx, `
		UPDATE feira SET
		XLONG = ?,
		LAT = ?,
		SETCENS = ?,
		AREAP = ?,
		CODDIST = ?,
		DISTRITO = ?,
		CODSUBPREF = ?,
		SUBPREFE = ?,
		REGIAO5 = ?,
		REGIAO8 = ?,
		NOME_FEIRA = ?,
		REGISTRO = ?,
		LOGRADOURO = ?,
		NUMERO = ?,
		BAIRRO = ?,
		REFERENCIA = ?,
		UPDATED_AT = CURRENT_TIMESTAMP()
		WHERE ID = ? AND DELETED_AT IS NULL;
	`, feira.XLONG, feira.LAT, feira.SETCENS, feira.AREAP, feira.CODDIST, feira.DISTRITO, feira.CODSUBPREF, feira.SUBPREFE, feira.REGIAO5, feira.REGIAO8, feira.NOME_FEIRA, feira.REGISTRO, feira.LOGRADOURO, feira.NUMERO, feira.BAIRRO, feira.REFERENCIA, id)
	if err != nil {
		logger.ErrorContext(ctx, "store.feira.Update", err.Error())
		logger.L.Println(time.Now(), "não foi possível modificar a feira:", err.Error())
		return "", fverr.New(http.StatusInternalServerError, "não foi possível modificar a feira", nil)
	}

	return id, nil
}

func (r *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := r.writer.ExecContext(ctx, `UPDATE feira SET DELETED_AT = CURRENT_TIMESTAMP() WHERE id = ?;`, id)
	if err != nil {
		logger.ErrorContext(ctx, "store.feira.Delete", err.Error())
		logger.L.Println(time.Now(), "não foi possível deletar a feira:", err.Error())
		return fverr.New(http.StatusInternalServerError, "não foi possível deletar a feira", map[string]string{
			"id": id,
		})
	}

	return nil
}

func (r *storeImpl) Search(ctx context.Context, distrito string) ([]*model.Feira, error) {
	var feiras []*model.Feira
	err := r.reader.SelectContext(ctx, &feiras, `
	SELECT 
		ID,
		ID_FEIRA,
		XLONG,
		LAT,
		SETCENS,
		AREAP,
		CODDIST,
		DISTRITO,
		CODSUBPREF,
		SUBPREFE,
		REGIAO5,
		REGIAO8,
		NOME_FEIRA,
		REGISTRO,
		LOGRADOURO,
		NUMERO,
		BAIRRO,
		REFERENCIA,
		CREATED_AT,
		UPDATED_AT,
		DELETED_AT
	FROM feira where DISTRITO LIKE ? AND 
	DELETED_AT IS NULL
	ORDER BY CREATED_AT DESC;
	`, "%"+distrito+"%")

	if err != nil {
		logger.ErrorContext(ctx, "store.feira.Search", err.Error())
		logger.L.Println(time.Now(), "store.feira.Search:", err.Error())
		return nil, fverr.New(http.StatusNotFound, "não encontrei feira com esses parâmetros", map[string]interface{}{
			"search": distrito,
		})
	}

	return feiras, nil
}
