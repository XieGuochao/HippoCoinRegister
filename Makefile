.PHONY: clean

server-build:
	go build -o bin/hippo-register hippo-register/main.go 

server: server-build
	./bin/hippo-register

server-bg: server-build
	./bin/hippo-register &

test: server-bg
	go test ./lib

clean:
	rm build/*