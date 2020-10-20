package lib

import (
	"log"
	"net"
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

// AddressesRefresh ...
// Get up to 30 addresses.
func (client *Client) AddressesRefresh(refresh RefreshStruct, reply *[]byte) error {
	return client.c.Call(HippoAddressServiceName+".AddressesRefresh", refresh, reply)
}

// Close ...
func (client *Client) Close() {
	client.c.Close()
}

// GetOutboundIP ...
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
