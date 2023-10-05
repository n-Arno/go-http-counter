all: build

build:
	GOOS=linux GOARCH=amd64 go build -o go-http-counter-linux-amd64

test:
	go build && ./go-http-counter

clean:
	- rm -f go-http-counter*
