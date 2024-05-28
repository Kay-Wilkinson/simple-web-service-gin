up:
	docker-compose up --build

upd:
	docker-compose up --build -d

.PHONY: test

test:
	docker-compose exec app go test ./tests/...

tidy:
	docker-compose exec app -c "go mod tidy"

down:
	docker-compose down