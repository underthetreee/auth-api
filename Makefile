.SILENT:

build:
	go build -o bin/app cmd/*.go

run: build
	./bin/app

.PHONY: compose-up
compose-up:
	docker-compose up -d

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans