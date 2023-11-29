install-tools:
	if [ ! $$(which go) ]; then \
		echo "goLang not found."; \
		echo "Try installing go..."; \
		exit 1; \
	fi
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/axw/gocov/gocov@latest
	go get golang.org/x/tools/cmd/goimports
	go install github.com/AlekSi/gocov-xml@latest
	if [ ! $$( which migrate ) ]; then \
		echo "The 'migrate' command was not found in your path. You most likely need to add \$$HOME/go/bin to your PATH."; \
		exit 1; \
	fi

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

test: tidy
	gocov test ./... | gocov report 

coverage: 
	gocov test ./...  | gocov-xml > coverage.cobertura.xml

build:
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api/api.go


package:
	docker  build -t $(tag) . 

run:database
	go mod tidy
	if [ ! -f '.env' ]; then \
		cp .env.example .env; \
	fi
	go run ./cmd/api/api.go 

create-migration: ## usage: make name=new create-migration
	migrate create -ext sql -dir ./db/migrations -seq $(name)

database:
	docker-compose up -d

gen:
	go mod tidy
	go generate ./...
