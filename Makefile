BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...
	go test ./... -coverprofile coverage.out

test-output: 
	go tool cover -html=coverage.out

engine:
	go build -o ${BINARY} app/*.go

docker:
	docker-compose up --build -d

unittest:
	docker-compose up --build -d
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

tidy:
	go mod tidy
	go mod vendor

run:
	docker-compose up --build -d
	go run app/main.go

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

seed:
	go run database/seeder/main.go

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint