build:
	go build -o bin/main main.go

dev: 
	./bin/air

run:
	go run main.go

test: 
	go test -v ./tests/${module}...