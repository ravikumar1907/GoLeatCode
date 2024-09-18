package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

const (
	Follower  = "Follower"
	Candidate = "Candidate"
	Leader    = "Leader"

	HeartbeatInterval = 100 * time.Millisecond
	ElectionTimeout   = 300 * time.Millisecond
)

type RequestVoteArgs struct {
	Term         int
	CandidateID  int
	LastLogIndex int
	LastLogTerm  int
}

type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

type AppendEntriesArgs struct {
	Term     int
	LeaderID int
}

type AppendEntriesReply struct {
	Term    int
	Success bool
}

type RaftNode struct {
	ID          int
	state       string
	currentTerm int
	votedFor    int
	mutex       sync.Mutex
	voteCount   int
	heartbeat   chan bool
	address     string
}

func (n *RaftNode) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if args.Term > n.currentTerm {
		n.currentTerm = args.Term
		n.votedFor = -1
	}

	reply.Term = n.currentTerm
	reply.VoteGranted = (n.votedFor == -1 || n.votedFor == args.CandidateID) && args.Term >= n.currentTerm
	if reply.VoteGranted {
		n.votedFor = args.CandidateID
	}
	return nil
}

func (n *RaftNode) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if args.Term > n.currentTerm {
		n.currentTerm = args.Term
		n.votedFor = -1
		reply.Term = n.currentTerm
		reply.Success = false
		return nil
	}

	n.currentTerm = args.Term
	n.state = Follower
	n.heartbeat <- true // Signal that a heartbeat was received

	reply.Term = n.currentTerm
	reply.Success = true
	return nil
}

func (n *RaftNode) StartServer() {
	rpc.Register(n)
	listener, err := net.Listen("tcp", n.address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	go rpc.Accept(listener)
}

func SendRequestVote(nodeAddr string, args *RequestVoteArgs) (*RequestVoteReply, error) {
	client, err := rpc.Dial("tcp", nodeAddr)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var reply RequestVoteReply
	err = client.Call("RaftNode.RequestVote", args, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SendHeartbeat(nodeAddr string, args *AppendEntriesArgs) (*AppendEntriesReply, error) {
	client, err := rpc.Dial("tcp", nodeAddr)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var reply AppendEntriesReply
	err = client.Call("RaftNode.AppendEntries", args, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (n *RaftNode) startElection(nodes []*RaftNode) {
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
			go func(peer *RaftNode) {
				peerAddr := peer.address
				args := &RequestVoteArgs{
					Term:         n.currentTerm,
					CandidateID:  n.ID,
					LastLogIndex: 0, // Placeholder for actual last log index
					LastLogTerm:  0, // Placeholder for actual last log term
				}
				reply, err := SendRequestVote(peerAddr, args)
				if err != nil {
					fmt.Printf("Node %d failed to request vote from %s: %v\n", n.ID, peerAddr, err)
					return
				}

				if reply.Term > n.currentTerm {
					n.mutex.Lock()
					n.currentTerm = reply.Term
					n.votedFor = -1
					n.mutex.Unlock()
					return
				}

				if reply.VoteGranted {
					n.mutex.Lock()
					n.voteCount++
					if n.voteCount > len(nodes)/2 {
						n.becomeLeader(nodes)
					}
					n.mutex.Unlock()
				}
			}(peer)
		}
	}
}

func (n *RaftNode) becomeLeader(nodes []*RaftNode) {
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

func (n *RaftNode) sendHeartbeat(peer *RaftNode) {
	args := &AppendEntriesArgs{
		Term:     n.currentTerm,
		LeaderID: n.ID,
	}
	reply, err := SendHeartbeat(peer.address, args)
	if err != nil {
		fmt.Printf("Leader Node %d failed to send heartbeat to Node %d: %v\n", n.ID, peer.ID, err)
		return
	}

	if reply.Term > n.currentTerm {
		n.mutex.Lock()
		n.currentTerm = reply.Term
		n.state = Follower
		n.votedFor = -1
		n.mutex.Unlock()
	}
	fmt.Printf("Leader Node %d sent heartbeat to Node %d\n", n.ID, peer.ID)
}

func (n *RaftNode) startFollower(nodes []*RaftNode) {
	for {
		if n.state == Follower {
			select {
			case <-n.heartbeat:
				// Received heartbeat, reset the election timer
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
	nodeCount := 3
	nodes := []*RaftNode{}
	for i := 1; i <= nodeCount; i++ {
		address := fmt.Sprintf("localhost:%d", 8000+i)
		node := &RaftNode{
			ID:        i,
			state:     Follower,
			votedFor:  -1,
			heartbeat: make(chan bool),
			address:   address,
		}
		nodes = append(nodes, node)
		go node.StartServer()
	}

	// Simulate follower nodes waiting for heartbeats and starting elections
	for _, node := range nodes {
		go node.startFollower(nodes)
	}

	// Prevent the main function from exiting
	select {}
}

/*
Explanation:
RaftNode Struct:

Represents each node in the Raft cluster, with fields for state, term, vote count, and a heartbeat channel.
RPC Methods:

RequestVote: Handles vote requests from candidates.
AppendEntries: Handles heartbeats and log entries from the leader.
Election and Heartbeat Logic:

startElection: Starts a new election if the node is a candidate.
becomeLeader: Updates the node to the leader state and begins sending heartbeats.
sendHeartbeat: Sends heartbeat messages to other nodes.
startFollower: Manages the follower state, responding to heartbeats and handling election timeouts.
Main Function:

Initializes nodes, starts their servers, and simulates their operation.
*/
