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

	mockgen -source=./store/company/company.go -destination=./mocks/company_store_mock.go -package=mocks -mock_names=Store=MockCompanyStore
	mockgen -source=./store/health/health.go -destination=./mocks/health_mock.go -package=mocks -mock_names=Store=MockHealthStore
	mockgen -source=./app/company/company.go -destination=./mocks/company_app_mock.go -package=mocks -mock_names=App=MockCompanyApp
