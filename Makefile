test: vet
	@export TEST_MYSQL_CONNECTION_STRING="root@tcp(127.0.0.1:3306)/dream_deeply_test?charset=utf8mb4&parseTime=True"; \
	go test -v -cover -tags=integration ./internal/product/...;

vet: 
	@go vet .
	@go vet ./cmd/...