package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// ConsistentHash represents a consistent hashing ring.
type ConsistentHash struct {
	// Hash ring stores points on the ring.
	ring []int
	// Maps points to the server associated with that point.
	hashMap map[int]string
	// Number of virtual nodes per server.
	replicas int
	// Mutex for thread safety.
	mu sync.RWMutex
}

// NewConsistentHash initializes a new consistent hashing ring.
func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		ring:     []int{},
		hashMap:  make(map[int]string),
		replicas: replicas,
	}
}

// hash computes the hash of a given key.
func (ch *ConsistentHash) hash(key string) int {
	// Use CRC32 hashing function for simplicity.
	return int(crc32.ChecksumIEEE([]byte(key)))
}

// AddServer adds a new server to the hash ring.
func (ch *ConsistentHash) AddServer(server string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	// Add virtual nodes for the server.
	for i := 0; i < ch.replicas; i++ {
		virtualNodeKey := server + "#" + strconv.Itoa(i)
		hash := ch.hash(virtualNodeKey)
		ch.ring = append(ch.ring, hash)
		ch.hashMap[hash] = server
	}

	// Sort the ring after adding a new server.
	sort.Ints(ch.ring)
}

// RemoveServer removes a server from the hash ring.
func (ch *ConsistentHash) RemoveServer(server string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	// Remove all virtual nodes of the server.
	for i := 0; i < ch.replicas; i++ {
		virtualNodeKey := server + "#" + strconv.Itoa(i)
		hash := ch.hash(virtualNodeKey)

		// Remove hash from the ring.
		index := sort.Search(len(ch.ring), func(i int) bool { return ch.ring[i] >= hash })
		if index < len(ch.ring) && ch.ring[index] == hash {
			ch.ring = append(ch.ring[:index], ch.ring[index+1:]...)
			delete(ch.hashMap, hash)
		}
	}
}

// GetServer returns the server responsible for a given key.
func (ch *ConsistentHash) GetServer(key string) string {
	ch.mu.RLock()
	defer ch.mu.RUnlock()

	if len(ch.ring) == 0 {
		return ""
	}

	// Compute the hash of the key.
	keyHash := ch.hash(key)

	// Binary search to find the appropriate server.
	index := sort.Search(len(ch.ring), func(i int) bool { return ch.ring[i] >= keyHash })

	// Wrap around if needed.
	if index == len(ch.ring) {
		index = 0
	}

	return ch.hashMap[ch.ring[index]]
}

func main() {
	// Create a consistent hash ring with 3 virtual nodes per server.
	ch := NewConsistentHash(3)

	// Add servers to the hash ring.
	ch.AddServer("Server1")
	ch.AddServer("Server2")
	ch.AddServer("Server3")
	// ch.AddServer("Server4")
	// ch.AddServer("Server5")
	// ch.AddServer("Server6")
	// ch.AddServer("Server7")
	// ch.AddServer("Server8")

	// Get the server for some keys.
	fmt.Println("Key1 is mapped to", ch.GetServer("Key1"))
	fmt.Println("Key2 is mapped to", ch.GetServer("Key2"))
	fmt.Println("Key3 is mapped to", ch.GetServer("Key3"))
	fmt.Println("Key4 is mapped to", ch.GetServer("Key4"))
	fmt.Println("Key5 is mapped to", ch.GetServer("Key5"))
	fmt.Println("Key6 is mapped to", ch.GetServer("Key6"))
	fmt.Println("Key7 is mapped to", ch.GetServer("Key7"))
	fmt.Println("Key8 is mapped to", ch.GetServer("Key8"))

	// Remove a server.
	ch.RemoveServer("Server2")

	// Check the mapping again.
	fmt.Println("Key1 is now mapped to", ch.GetServer("Key1"))
	fmt.Println("Key2 is now mapped to", ch.GetServer("Key2"))
	fmt.Println("Key3 is now mapped to", ch.GetServer("Key3"))
	fmt.Println("Key4 is now mapped to", ch.GetServer("Key4"))
	fmt.Println("Key5 is now mapped to", ch.GetServer("Key5"))
	fmt.Println("Key6 is now mapped to", ch.GetServer("Key6"))
	fmt.Println("Key7 is now mapped to", ch.GetServer("Key7"))
	fmt.Println("Key8 is now mapped to", ch.GetServer("Key8"))
}
