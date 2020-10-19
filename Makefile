.PHONY: clean

server-build:
	go build -o bin/server register/main.go 

server: server-build
	./bin/server

server-bg: server-build
	./bin/server &

test: server-bg
	go test ./lib

clean:
	rm build/*