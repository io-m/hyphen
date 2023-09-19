## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/main.go

db-test-up:
	docker-compose -f infra/docker-compose.tst.yml up

db-test-down:
	docker-compose -f infra/docker-compose.tst.yml down

db-dev-up:
	docker-compose -f infra/docker-compose.dev.yml up

db-dev-down:
	docker-compose -f infra/docker-compose.dev.yml down

db-prod-up:
	docker-compose -f infra/docker-compose.prod.yml up

db-prod-down:
	docker-compose -f infra/docker-compose.prod.yml down
