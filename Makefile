test: vet
	@export TEST_MYSQL_CONNECTION_STRING="root@tcp(127.0.0.1:3306)/dream_deeply_test?charset=utf8mb4&parseTime=True"; \
	go test -v -cover -tags=integration ./pkg/product/...;

vet: 
	@go vet .
	@go vet ./cmd/...

migrate-up:
	migrate -path db/migrations -database "mysql://root:thelastword@tcp(127.0.0.1:3306)/dreamdeeply?charset=utf8mb4&parseTime=True" up

migrate-down:
	migrate -path db/migrations -database "mysql://root:thelastword@tcp(127.0.0.1:3306)/dreamdeeply?charset=utf8mb4&parseTime=True" down
