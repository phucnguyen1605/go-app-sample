start-api:
	GO111MODULE=on; cd src/cmd/api; go run .

mocks:
	cd src/pkg; mockery --all
	cd src/repository; mockery --all
	cd src/service; mockery --all

unittest:
	GO111MODULE=on; cd src; go test --cover ./...

build:
	GO111MODULE=on; cd src; go build -v ./...

lint:
	cd src; golangci-lint run

run:
	docker-compose up -d