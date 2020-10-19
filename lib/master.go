package lib

import (
	"context"
	"encoding/json"
	"log"
	"net/rpc"
	"sync"
	"time"
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

// cache
var cache sync.Map

// Addresses ...
type Addresses = []string

// HippoAddressServiceInterface ...
type HippoAddressServiceInterface = interface {
	Ping(request string, reply *string) error
}

// RegisterHippoAddress ...
func RegisterHippoAddress(svc HippoAddressServiceInterface) error {
	return rpc.RegisterName(HippoAddressServiceName, svc)
}

// ServiceStruct ...
type ServiceStruct struct {
}

// Ping ...
func (s *ServiceStruct) Ping(request string, reply *string) error {
	log.Println("Ping")
	return nil
}

// Register ...
func (s *ServiceStruct) Register(address string, reply *string) error {
	log.Println("Register address:", address)
	cache.Store(address, time.Now().Unix())
	return nil
}

func expired(t, now int64) bool {
	return now-t > TTL
}

// Addresses ...
func (s *ServiceStruct) Addresses(number int, reply *[]byte) error {
	log.Println("Query addresses:", number)
	if number > MaxQuery || number < 0 {
		number = MaxQuery
	}

	addresses := new(Addresses)
	*addresses = make([]string, MaxQuery)
	now := time.Now().Unix()
	count := 0
	cache.Range(func(key interface{}, value interface{}) bool {
		// Check invalid.
		if expired(value.(int64), now) {
			cache.Delete(key)
		} else {
			(*addresses)[count] = key.(string)
			count++
		}
		return count < number
	})
	(*addresses) = (*addresses)[:count]
	log.Println(addresses)

	b, _ := json.Marshal(*addresses)
	*reply = b
	return nil
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
