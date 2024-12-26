run-api-dc1:
	@POSTGRES_PORT=54321 DC_NAME=DC1 DC_ADDRESS=":13231" go run ./app/api/main.go

run-api-dc2:
	@POSTGRES_PORT=54322 DC_NAME=DC2 DC_ADDRESS=":13232" go run ./app/api/main.go
