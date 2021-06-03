test:
	go test ./... -race -count=1

lint:
	golangci-lint run --deadline=5m -v

build:
	go build ./cmd/api/

install_moq:
	go get github.com/matryer/moq

up:
	docker-compose -f docker/docker-compose.yml up -d --build

down:
	docker-compose -f docker/docker-compose.yml down
