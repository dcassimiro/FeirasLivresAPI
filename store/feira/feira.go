package feira

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
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
		INSERT INTO feira (idfeira, xlong, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5t, regiao8t, nomefeira, registrot, logradourot, numerot, bairrot, referenciat)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`, feira.IDFEIRA, feira.XLONG, feira.LAT, feira.SETCENS, feira.AREAP, feira.CODDIST, feira.DISTRITO, feira.CODSUBPREF, feira.SUBPREFE, feira.REGIAO5T, feira.REGIAO8T, feira.NOMEFEIRA, feira.REGISTROT, feira.LOGRADOUROT, feira.NUMEROT, feira.BAIRROT, feira.REFERENCIAT)
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
			idfeira,
			xlong,
			lat,
			setcens,
			areap,
			coddist,
			distrito,
			codsubpref,
			subprefe,
			regiao5t,
			regiao8t,
			nomefeira,
			registrot,
			logradourot,
			numerot,
			bairrot,
			referenciat,
			createdat,
			updatedat,
			deletedat
		FROM feira 
		WHERE
			id = ?;
	`, id)
	if err != nil {
		fmt.Println(ctx, "store.feira.ReadOne", err.Error())
		return nil, err
	}

	return feira, nil
}

func (r *storeImpl) Update(ctx context.Context, id string, feira model.FeiraRequest) (string, error) {
	_, err := r.writer.ExecContext(ctx, `
		UPDATE feira 
		SET 
		xlong = ?,
		lat = ?,
		setcens = ?,
		areap = ?,
		coddist = ?,
		distrito = ?,
		codsubpref = ?,
		subprefe = ?,
		regiao5t = ?,
		regiao8t = ?,
		nomefeira = ?,
		registrot = ?,
		logradourot = ?,
		numerot = ?,
		bairrot = ?,
		referenciat = ? ,
		updatedat = CURRENT_TIMESTAMP()
		WHERE id = ? AND deletedat IS NULL;
	`, feira.XLONG, feira.LAT, feira.SETCENS, feira.AREAP, feira.CODDIST, feira.DISTRITO, feira.CODSUBPREF, feira.SUBPREFE, feira.REGIAO5T, feira.REGIAO8T, feira.NOMEFEIRA, feira.REGISTROT, feira.LOGRADOUROT, feira.NUMEROT, feira.BAIRROT, feira.REFERENCIAT, id)
	if err != nil {
		fmt.Println(ctx, "store.company.Update", err.Error())
		return "", err
	}

	return id, nil
}

func (r *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := r.writer.ExecContext(ctx, `UPDATE feira SET deleted_at = CURRENT_TIMESTAMP() WHERE id = ?;`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *storeImpl) Search(ctx context.Context, distrito string) ([]*model.Feira, error) {
	var feiras []*model.Feira
	err := r.reader.SelectContext(ctx, &feiras, `
	SELECT 
		id,
		idfeira,
		xlong,
		lat,
		setcens,
		areap,
		coddist,
		distrito,
		codsubpref,
		subprefe,
		regiao5t,
		regiao8t,
		nomefeira,
		registrot,
		logradourot,
		numerot,
		bairrot,
		referenciat,
		createdat,
		updatedat,
		deletedat
	FROM feira where distrito LIKE ? AND 
	deletedat IS NULL
	ORDER BY createdat DESC;
	`, "%"+distrito+"%")

	if err != nil {
		return nil, err
	}

	return feiras, nil
}
