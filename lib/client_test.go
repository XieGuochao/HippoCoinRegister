package lib

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"testing"
	"time"
)

var response string

func TestPing(t *testing.T) {
	client, _ := CreateClient("tcp", "localhost:9325")
	log.Println(client)

	log.Println("ping error:", client.Ping("a", &response))
	log.Println("test ping finish.")
}

func TestRegisterAddresses(t *testing.T) {
	reply := ""
	client, _ := CreateClient("tcp", "localhost:9325")
	log.Println(client)

	a1 := GetOutboundIP().String() + ":"

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalln(err)
	}
	go listener.Accept()
	defer listener.Close()

	a1 += strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	log.Println("my address:", a1)
	log.Println("register error:", client.Register(a1, &reply))

	var addresses []byte
	log.Println("addresses error:", client.Addresses(10, &addresses))
	log.Println("addresses:", addresses)

	var result []string
	json.Unmarshal(addresses, &result)
	log.Println("addresses (decoded):", result)

	time.Sleep(time.Second)
	log.Println("test register done.")
}

func TestRegisterAddressesRefresh(t *testing.T) {
	reply := ""
	client, _ := CreateClient("tcp", "localhost:9325")
	log.Println(client)

	a1 := GetOutboundIP().String() + ":"

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalln(err)
	}
	go listener.Accept()
	defer listener.Close()

	a1 += strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	log.Println("my address:", a1)
	log.Println("register error:", client.Register(a1, &reply))

	var addresses []byte
	log.Println("addresses error:", client.Addresses(10, &addresses))
	log.Println("addresses:", addresses)

	var result []string
	json.Unmarshal(addresses, &result)
	log.Println("addresses (decoded):", result)

	time.Sleep(time.Second)
	log.Println("addresses error:", client.AddressesRefresh(RefreshStruct{10, a1}, &addresses))
	log.Println("addresses:", addresses)

	json.Unmarshal(addresses, &result)
	log.Println("addresses (decoded):", result)

	log.Println("test register and refresh done.")

}

func TestOutboundIP(t *testing.T) {
	log.Println(GetOutboundIP())
}
