package feira_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/unico/FeirasLivresAPI/app/feira"
	"github.com/unico/FeirasLivresAPI/fverr"
	"github.com/unico/FeirasLivresAPI/mocks"
	"github.com/unico/FeirasLivresAPI/model"
	"github.com/unico/FeirasLivresAPI/store"

	"github.com/unico/FeirasLivresAPI/test"
)

var (
	defaultDate  = time.Now()
	defaultError = fverr.New(http.StatusInternalServerError, "ocorreu um erro", nil)
)

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Feira

		InputFeira  model.FeiraRequest
		PrepareMock func(mockStore *mocks.MockFeiraStore)
	}{
		"deve retornar sucesso": {
			ExpectedData: &model.Feira{
				ID:         "default-id",
				ID_FEIRA:   "1",
				XLONG:      "-12312313",
				LAT:        "-2312313213",
				SETCENS:    "355030810000027",
				AREAP:      "3550308005005",
				CODDIST:    "10",
				DISTRITO:   "BRAS",
				CODSUBPREF: "10",
				SUBPREFE:   "MOCCA",
				REGIAO5:    "leste",
				REGIAO8:    "leste 2",
				NOME_FEIRA: "CONCORDIA",
				REGISTRO:   "4003-7",
				CREATED_AT: defaultDate,
			},

			InputFeira: model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"}).
					Times(1).
					Return("default-id", nil)

				mockStore.EXPECT().ReadOne(gomock.Any(), "default-id").
					Times(1).
					Return(&model.Feira{
						ID:         "default-id",
						ID_FEIRA:   "1",
						XLONG:      "-12312313",
						LAT:        "-2312313213",
						SETCENS:    "355030810000027",
						AREAP:      "3550308005005",
						CODDIST:    "10",
						DISTRITO:   "BRAS",
						CODSUBPREF: "10",
						SUBPREFE:   "MOCCA",
						REGIAO5:    "leste",
						REGIAO8:    "leste 2",
						NOME_FEIRA: "CONCORDIA",
						REGISTRO:   "4003-7",
						CREATED_AT: defaultDate,
					}, nil)
			},
		},
		"deve retornar erro na criação": {
			ExpectedErr: defaultError,

			InputFeira: model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"}).
					Times(1).
					Return("default-id", defaultError)
			},
		},
		"deve retornar erro na leitura": {
			ExpectedErr: defaultError,

			InputFeira: model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"},
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.FeiraRequest{ID_FEIRA: "1", XLONG: "-12312313", LAT: "-2312313213", SETCENS: "355030810000027", AREAP: "3550308005005", CODDIST: "10", DISTRITO: "BRAS", CODSUBPREF: "10", SUBPREFE: "MOCCA", REGIAO5: "leste", REGIAO8: "leste 2", NOME_FEIRA: "CONCORDIA", REGISTRO: "4003-7"}).
					Times(1).
					Return("default-id", nil)

				mockStore.EXPECT().ReadOne(gomock.Any(), "default-id").
					Times(1).
					Return(nil, defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockStore := mocks.NewMockFeiraStore(ctrl)

			cs.PrepareMock(mockStore)

			app := feira.NewApp(&store.Container{Feira: mockStore})

			company, err := app.Create(ctx, cs.InputFeira)

			if diff := cmp.Diff(company, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestReadOne(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Feira

		InputID     string
		PrepareMock func(mockStore *mocks.MockFeiraStore)
	}{
		"deve retornar sucesso": {
			ExpectedData: &model.Feira{
				ID:         "default-id",
				ID_FEIRA:   "1",
				XLONG:      "-12312313",
				LAT:        "-2312313213",
				SETCENS:    "355030810000027",
				AREAP:      "3550308005005",
				CODDIST:    "10",
				DISTRITO:   "BRAS",
				CODSUBPREF: "10",
				SUBPREFE:   "MOCCA",
				REGIAO5:    "leste",
				REGIAO8:    "leste 2",
				NOME_FEIRA: "CONCORDIA",
				REGISTRO:   "4003-7",
				CREATED_AT: defaultDate,
			},

			InputID: "default-id",
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().ReadOne(gomock.Any(), "default-id").
					Times(1).
					Return(&model.Feira{
						ID:         "default-id",
						ID_FEIRA:   "1",
						XLONG:      "-12312313",
						LAT:        "-2312313213",
						SETCENS:    "355030810000027",
						AREAP:      "3550308005005",
						CODDIST:    "10",
						DISTRITO:   "BRAS",
						CODSUBPREF: "10",
						SUBPREFE:   "MOCCA",
						REGIAO5:    "leste",
						REGIAO8:    "leste 2",
						NOME_FEIRA: "CONCORDIA",
						REGISTRO:   "4003-7",
						CREATED_AT: defaultDate,
					}, nil)
			},
		},
		"deve retornar um erro": {
			ExpectedErr: defaultError,

			InputID: "default-id",
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().ReadOne(gomock.Any(), "default-id").
					Times(1).
					Return(nil, defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockStore := mocks.NewMockFeiraStore(ctrl)

			cs.PrepareMock(mockStore)

			app := feira.NewApp(&store.Container{Feira: mockStore})
			feira, err := app.ReadOne(ctx, cs.InputID)
			if diff := cmp.Diff(feira, cs.ExpectedData); diff != "" {
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
		ExpectedErr  error
		ExpectedData *model.Feira

		InputID     string
		PrepareMock func(mockStore *mocks.MockFeiraStore)
	}{
		"deve retornar sucesso": {
			ExpectedData: &model.Feira{
				ID:         "default-id",
				ID_FEIRA:   "1",
				XLONG:      "-12312313",
				LAT:        "-2312313213",
				SETCENS:    "355030810000027",
				AREAP:      "3550308005005",
				CODDIST:    "10",
				DISTRITO:   "BRAS",
				CODSUBPREF: "10",
				SUBPREFE:   "MOCCA",
				REGIAO5:    "leste",
				REGIAO8:    "leste 2",
				NOME_FEIRA: "CONCORDIA",
				REGISTRO:   "4003-7",
				CREATED_AT: defaultDate,
			},

			InputID: "default-id",
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().Delete(gomock.Any(), "default-id").
					Times(1).
					Return(nil)
			},
		},
		"deve retornar um erro": {
			ExpectedErr: defaultError,

			InputID: "default-id",
			PrepareMock: func(mockStore *mocks.MockFeiraStore) {
				mockStore.EXPECT().Delete(gomock.Any(), "default-id").
					Times(1).
					Return(defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockStore := mocks.NewMockFeiraStore(ctrl)

			cs.PrepareMock(mockStore)

			app := feira.NewApp(&store.Container{Feira: mockStore})

			err := app.Delete(ctx, cs.InputID)
			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
