.PHONY:make

make: make run

run:
	go run main.go

build: 
	go build -o ./api ./main.go