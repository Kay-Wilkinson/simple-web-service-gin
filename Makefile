up:
	docker-compose up --build

down:
	docker-compose down

test:
	docker-compose exec app -c "go test ./tests/..."

tidy:
	docker-compose exec app -c "go mod tidy"