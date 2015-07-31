package raft

import (
	"log"
	"math/rand"
	"time"
)

type Role int

const (
	Follower Role = iota + 1
	Candidate
	Leader
)

type server struct {
	id int
	//Hearbit expired time (by millisecond.)
	expiredTime int
	role        Role
	nt          nodeNetwork
	//Message receive time
	msgRecvTime time.Time
	//To determine if server still alive, for kill testing.
	isAlive bool
	//id list exist in this network.
	nodeList []int
}

//New a server and given a random expired time.
func NewServer(id int, role Role, nt nodeNetwork, nodeList ...int) *server {
	rand.Seed(time.Now().UnixNano())
	expiredMiliSec := rand.Intn(5) + 1
	serv := &server{id: id, role: role, nt: nt, expiredTime: expiredMiliSec, isAlive: true, nodeList: nodeList}
	return serv
}

func (sev *server) runServerLoop() {

	for {
		switch sev.role {
		case Leader:
			sev.runLeaderLoop()
		case Candidate:
			sev.runCandidateLoop()
		case Follower:
			sev.runFollowerLoop()
		}

		//timer base on milli-second.
		time.Sleep(time.Millisecond)
	}
}

func (sev *server) sendHearbit() {
	for _, node := range sev.nodeList {
		hbMsg := message{from: sev.id, to: node, typ: Heartbit, val: "HB"}
		sev.nt.send(hbMsg)
	}
}

func (sev *server) runLeaderLoop() {
	sev.sendHearbit()

}

func (sev *server) runCandidateLoop() {
}

func (sev *server) runFollowerLoop() {

}

func (sev *server) roleChange(newRole Role) {
	log.Println("note:", sev.id, " change role from ", sev.role, " to ", newRole)
	sev.role = newRole
}
