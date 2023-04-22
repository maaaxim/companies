SHELL = /bin/sh

docker_up:
	PWD=$(PWD) user=$(id -u) group=$(id -g) docker-compose --file build/dev/docker-compose-local.yml up

run_main:
	go run main.go

# автогенерация, например для моков
generate:
	go generate ./...

tests:
	go test -v ./...

tests_failed:
	 go test -v ./... | grep FAIL

lint:
	golangci-lint run