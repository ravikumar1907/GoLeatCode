package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// const (
// 	Follower  = "Follower"
// 	Candidate = "Candidate"
// 	Leader    = "Leader"

// 	// HeartbeatInterval is the interval at which the leader sends heartbeats
// 	HeartbeatInterval = 100 * time.Millisecond

// 	// ElectionTimeout is the base duration after which a follower starts an election
// 	ElectionTimeout = 300 * time.Millisecond
// )

type Node struct {
	ID          int
	state       string
	currentTerm int
	votedFor    int
	mutex       sync.Mutex
	voteCount   int
	heartbeat   chan bool // Channel to receive heartbeats
}

// startElection initiates the election process
func (n *Node) startElection(nodes []*Node) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.state == Leader {
		return // Already a leader
	}

	// Become a candidate and start a new election
	n.state = Candidate
	n.currentTerm++
	n.votedFor = n.ID
	n.voteCount = 1 // Vote for itself

	fmt.Printf("Node %d started election for term %d\n", n.ID, n.currentTerm)

	// Request votes from other nodes
	for _, peer := range nodes {
		if peer.ID != n.ID {
			go func(peer *Node) {
				if n.requestVote(peer) {
					n.mutex.Lock()
					defer n.mutex.Unlock()
					n.voteCount++
					if n.voteCount > len(nodes)/2 {
						n.becomeLeader(nodes)
					}
				}
			}(peer)
		}
	}
}

// requestVote simulates requesting a vote from another node
func (n *Node) requestVote(peer *Node) bool {
	peer.mutex.Lock()
	defer peer.mutex.Unlock()

	// Check if the peer can grant a vote
	if peer.currentTerm < n.currentTerm {
		// Update peer's term to the current term and grant vote
		peer.currentTerm = n.currentTerm
		peer.votedFor = n.ID
		fmt.Printf("Node %d granted vote to Node %d for term %d\n", peer.ID, n.ID, n.currentTerm)
		return true
	} else if peer.currentTerm == n.currentTerm && (peer.votedFor == -1 || peer.votedFor == n.ID) {
		// If the term is the same, grant vote only if it hasn't voted yet or voted for this candidate
		peer.votedFor = n.ID
		fmt.Printf("Node %d granted vote to Node %d for term %d\n", peer.ID, n.ID, n.currentTerm)
		return true
	} else {
		// Vote not granted
		fmt.Printf("Node %d denied vote to Node %d for term %d\n", peer.ID, n.ID, n.currentTerm)
		return false
	}
}

// becomeLeader transitions the candidate to the leader role and starts heartbeats
func (n *Node) becomeLeader(nodes []*Node) {
	fmt.Printf("Node %d became the Leader\n", n.ID)
	n.state = Leader

	// Start sending heartbeats to all followers
	go func() {
		for n.state == Leader {
			for _, peer := range nodes {
				if peer.ID != n.ID {
					n.sendHeartbeat(peer)
				}
			}
			time.Sleep(HeartbeatInterval)
		}
	}()
}

// sendHeartbeat sends a heartbeat to a follower node
func (n *Node) sendHeartbeat(peer *Node) {
	peer.heartbeat <- true // Send a heartbeat to the follower
	fmt.Printf("Leader Node %d sent heartbeat to Node %d\n", n.ID, peer.ID)
}

// startFollower simulates a follower node waiting for a leader heartbeat
func (n *Node) startFollower(nodes []*Node) {
	for {
		if n.state == Follower {
			select {
			case <-n.heartbeat:
				// Received heartbeat, reset the timer
				fmt.Printf("Node %d (Follower) received heartbeat\n", n.ID)
			case <-time.After(ElectionTimeout + time.Duration(rand.Intn(150))*time.Millisecond):
				// Timeout waiting for heartbeat, start election
				fmt.Printf("Node %d (Follower) did not receive heartbeat, starting election\n", n.ID)
				n.startElection(nodes)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Create a cluster of nodes
	nodes := []*Node{}
	for i := 1; i <= 3; i++ {
		nodes = append(nodes, &Node{
			ID:        i,
			state:     Follower,
			votedFor:  -1, // No vote initially
			heartbeat: make(chan bool),
		})
	}

	// Simulate follower nodes waiting for heartbeats and starting elections
	for _, node := range nodes {
		go node.startFollower(nodes)
	}

	// Prevent the main function from exiting
	select {}
}
