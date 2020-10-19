package lib

import (
	"log"
	"net/rpc"
)

// Client ...
type Client struct {
	c       *rpc.Client
	address string
}

// CreateClient ...
func CreateClient(protocol, address string) (client *Client, err error) {
	client = new(Client)
	client.c, err = rpc.Dial(protocol, address)
	if err != nil {
		log.Println("create client error:", err)
		return nil, err
	}
	return client, err
}

// Ping ...
func (client *Client) Ping(request string, reply *string) error {
	return client.c.Call(HippoAddressServiceName+".Ping", request, reply)
}

// Register ...
func (client *Client) Register(address string, reply *string) error {
	return client.c.Call(HippoAddressServiceName+".Register", address, reply)
}

// Addresses ...
// Get up to 30 addresses.
func (client *Client) Addresses(number int, reply *[]byte) error {
	return client.c.Call(HippoAddressServiceName+".Addresses", number, reply)
}

// Close ...
func (client *Client) Close() {
	client.c.Close()
}
