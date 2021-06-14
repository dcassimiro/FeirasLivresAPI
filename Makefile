.PHONY: run watch mock  test

VERSION = $(shell git branch --show-current)

# comandos para execução

run:
	VERSION=$(VERSION) go run main.go

run-watch:
	VERSION=$(VERSION) nodemon --exec go run main.go --signal SIGTERM

# comandos para teste

test:
	go test -coverprofile=coverage.out ./app/... ./api/... ./model/... ./store/...

mock: 
	rm -rf ./mocks

	mockgen -source=./store/feira/feira.go -destination=./mocks/feira_store_mock.go -package=mocks -mock_names=Store=MockFeiraStore
	mockgen -source=./app/feira/feira.go -destination=./mocks/feira_app_mock.go -package=mocks -mock_names=App=MockFeiraApp
