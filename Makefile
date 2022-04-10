run:
	go run main.go

build:
	go build -o chuck main.go

all:
	go fmt main.go
	go build -o chuck main.go && \
	./chuck