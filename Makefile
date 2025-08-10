include .env

compose-dev:
	@docker-compose -f docker-compose.dev.yml --env-file .env up

cloudflared-minio:
	@cloudflared tunnel --url localhost:$(MINIO_API_PORT) 2>&1 | grep -o 'https://[^ ]*'

run:
	@go run cmd/main.go

ngrok-http:
	@ngrok http localhost:$(HTTP_SERVER_PORT)


