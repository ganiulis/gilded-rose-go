dcr=docker run -v .:/usr/src/app gilded-rose

build:
	docker build -t gilded-rose .

go:
	$(dcr) go

fixture:
	$(dcr) go run texttest_fixture.go $(days)

test:
	$(dcr) go test -v ./gildedrose/...

coverage:
	$(dcr) go test -v ./gildedrose/... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

