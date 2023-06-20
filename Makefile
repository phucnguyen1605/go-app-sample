start-api:
	GO111MODULE=on; cd src/cmd/api; go run .

mocks:
	cd src/pkg; mockery --all
	cd src/repository; mockery --all
	cd src/service; mockery --all

unittest:
	GO111MODULE=on; cd src; go test --cover ./...

lint:
	cd src; golangci-lint run --disable-all -E errcheck

run:
	docker-compose up -d