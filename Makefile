up:
	docker-compose up --build

.PHONY: test

test:
	docker-compose exec app go test ./tests/...

tidy:
	docker-compose exec app -c "go mod tidy"

down:
	docker-compose down