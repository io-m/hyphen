## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/api/main.go

db-dev-up:
	docker-compose -f docker-compose.dev.yml up

db-dev-down:
	docker-compose -f docker-compose.dev.yml down

db-prod-up:
	docker-compose -f docker-compose.prod.yml up

db-prod-down:
	docker-compose -f docker-compose.prod.yml down
