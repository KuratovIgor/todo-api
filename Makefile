run:
	docker-compose up -d --build
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up


swag:
	swag init -g cmd/main.go