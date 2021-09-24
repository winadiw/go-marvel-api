.PHONY: dependency unit-test integration-test docker-up docker-down clear swagger-build

dependency:
	@go get -v ./...

integration-test:
	@go test -v ./...

unit-test: dependency
	@go test -v -short ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

clear: docker-down

swagger-build:
	@swagger validate ./swagger.yml
	@redoc-cli bundle ./swagger.yml