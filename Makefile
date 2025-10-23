.PHONY: server frontend

server:
	cd ./cmd/server && go run .

frontend:
	cd ./frontend && npm install && npm run dev