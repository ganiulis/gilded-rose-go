dcr=docker run -v .:/usr/src/app gilded-rose

build:
	docker build -t gilded-rose .

cmd:
	$(dcr) $(args)

fixtures:
	$(dcr) go run main.go $(days)

test:
	$(dcr) go test -v ./...

coverage:
	$(dcr) go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

