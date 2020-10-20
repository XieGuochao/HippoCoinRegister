package main

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/XieGuochao/HippoCoinRegister/lib"
)

const (
	// HippoAddressServiceName ...
	HippoAddressServiceName = "github.com/XieGuochao/HippoCoinRegister"

	// TTL ...
	// The TTL is 10 seconds
	TTL = 10

	// MaxQuery ...
	MaxQuery = 30

	// MaxCycle ...
	MaxCycle = 5

	// NumPerCycle ...
	NumPerCycle = 20

	// Threshold ...
	Threshold = 0.2
)

func expired(t, now int64) bool {
	return now-t > TTL
}

// clearCycle
func clearCycle(ctx context.Context, c *sync.Map) {
	cleared := 1
	count := 1
	t := time.Now().Unix()

	for float64(cleared)/float64(count) > Threshold {
		cleared = 1
		count = 1
		c.Range(func(key, value interface{}) bool {
			select {
			case <-ctx.Done():
				return false
			default:
				{
					if expired(value.(int64), t) {
						c.Delete(key)
						cleared++
					}
					count++
					if count > NumPerCycle {
						return false
					}
				}
			}
			return true
		})
	}

}

// clear the outdated cache.
func clearCache(ctx context.Context, c *sync.Map) {
	for i := 0; i < MaxCycle; i++ {
		go clearCycle(ctx, c)
	}
}

func main() {
	log.Println("register server starts...")
	lib.RegisterHippoAddress(new(lib.ServiceStruct))
	listener, err := net.Listen("tcp", ":9325")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// clear cache for every 10 seconds

	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
			clearCache(ctx, &lib.Cache)
			<-ctx.Done()
			time.Sleep(time.Second)
			cancel()
		}
	}()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
