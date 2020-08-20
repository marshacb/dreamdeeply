test: vet
	@export TEST_MYSQL_CONNECTION_STRING="root@tcp(127.0.0.1:3306)/dreamdeeplyhair_test?charset=utf8mb4&parseTime=True"; \
	go test -v -cover -coverprofile=test_dreamdeeplyhair.out -tags=integration ./internal/product/...;

vet: 
	@go vet .
	@ go vet ./cmd/...