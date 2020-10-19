# HippoCoinRegister

The register server for HippoCoin using go RPC.

## Instructions

- `make server`: compile and run the server on port 9325.
- `make server-build`: compile the server to bin/server.
- `make server-bg`: compile and run the server on port 9325 on background.
- `make clean`: remove the compiled program.
- `lsof -i:9325`: find the program running on port 9325.

## Use the client library

Go to `client.go` and `client_test.go` for more information.