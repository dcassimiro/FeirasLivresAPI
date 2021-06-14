package feira_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	"github.com/unico/FeirasLivresAPI/fverr"
	"github.com/unico/FeirasLivresAPI/model"
	"github.com/unico/FeirasLivresAPI/store/feira"
	"github.com/unico/FeirasLivresAPI/test"
)

var defaultDate = time.Now()

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputPartner model.FeiraRequest
		PrepareMock  func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			InputPartner: model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO feira (ID_FEIRA, XLONG, LAT, SETCENS, AREAP, CODDIST, DISTRITO, CODSUBPREF, SUBPREFE, REGIAO5, REGIAO8, NOME_FEIRA, REGISTRO, LOGRADOURO, NUMERO, BAIRRO, REFERENCIA)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
				`).
					WithArgs("1", "-12312313", "-2312313213", "355030810000027", "3550308005005", "10", "BRAS", "10", "MOCCA", "leste", "leste 2", "CONCORDIA", "4003-7", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro com a mensagem: 'não foi possível criar uma nova feira'": {
			ExpectedErr: fverr.New(http.StatusInternalServerError, "não foi possível criar uma nova feira", nil),

			InputPartner: model.FeiraRequest{ID_FEIRA: "1"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO feira (ID_FEIRA, XLONG, LAT, SETCENS, AREAP, CODDIST, DISTRITO, CODSUBPREF, SUBPREFE, REGIAO5, REGIAO8, NOME_FEIRA, REGISTRO, LOGRADOURO, NUMERO, BAIRRO, REFERENCIA)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
				`).
					WithArgs("1").
					WillReturnError(errors.New("não foi possível criar uma nova feira"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := feira.NewStore(db, nil)
			ctx := context.Background()

			id, err := store.Create(ctx, cs.InputPartner)

			if err == nil && id == "" {
				t.Error(id)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	cases := map[string]struct {
		ExpectedData string
		ExpectedErr  error

		InputID     string
		InputFeira  model.FeiraRequest
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			ExpectedData: "update-id",

			InputID:    "update-id",
			InputFeira: model.FeiraRequest{XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
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
				`).
					WithArgs("-12312313", "-2312313213", "355030810000027", "3550308005005", "10", "BRAS", "10", "MOCCA", "leste", "leste 2", "CONCORDIA", "4003-7", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "update-id").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro com a mensagem: 'não foi possível modificar a feira'": {
			ExpectedErr: fverr.New(http.StatusInternalServerError, "não foi possível modificar a feira", nil),

			InputID:    "update-id",
			InputFeira: model.FeiraRequest{XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
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
				`).
					WithArgs("update-id").
					WillReturnError(errors.New("não foi possível modificar a feira"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := feira.NewStore(db, nil)
			ctx := context.Background()

			id, err := store.Update(ctx, cs.InputID, cs.InputFeira)

			if diff := cmp.Diff(id, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputID     string
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			InputID: "delete-id",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					UPDATE feira SET DELETED_AT = CURRENT_TIMESTAMP() WHERE id = ?;
				`).
					WithArgs("delete-id").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"deve retornar erro com a mensagem: 'não foi possível deletar a feira'": {
			ExpectedErr: fverr.New(http.StatusInternalServerError, "não foi possível deletar a feira", map[string]string{"id": "delete-id"}),

			InputID: "delete-id",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					UPDATE feira SET DELETED_AT = CURRENT_TIMESTAMP() WHERE id = ?;
				`).
					WithArgs("delete-id").
					WillReturnError(errors.New("não foi possível deletar a feira"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := feira.NewStore(db, nil)
			ctx := context.Background()

			err := store.Delete(ctx, cs.InputID)

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
