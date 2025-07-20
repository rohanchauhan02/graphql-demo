# Makefile for the Internal Transfer Project

.PHONY: postgres setup start build test clean

# Start a local PostgreSQL instance using Docker
postgres:
	docker pull postgres:latest
	docker run --name postgres \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=root \
		-e POSTGRES_DB=internal_transfer_local \
		-p 5432:5432 \
		-d postgres:latest

# Set up Go modules and configuration files, and start PostgreSQL
setup: postgres
	go mod tidy
	cp configs/app.config.sample.yml configs/app.config.local.yml

# Run the application
start:
	go run main.go
