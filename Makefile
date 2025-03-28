fs:
	@go build -o build/fs cmd/fs/main.go

duckdb:
	@go build -o build/duckdb cmd/duckdb/main.go