include .env.local

build:
	mkdir bin
	cp .env.local ./bin/.env.local
	go build -o ./bin/$(APPNAME) ./cmd/server

clean-bin:
	rm -rf bin

docker:
	docker build -t $(APPNAME) --build-arg ENV=$(ENV) --build-arg PORT=$(PORT) --build-arg APPNAME="$(APPNAME)" .

run:
	ENV=".env.local" go run ./cmd/server