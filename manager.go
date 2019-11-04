package main

import "log"

type Manager struct {
	register chan *User
	relation chan *Relation
	debug    chan int
}

func NewManager() *Manager {
	return &Manager{
		make(chan *User),
		make(chan *Relation),
		make(chan int),
	}
}
func (m Manager) manage(users map[string]*User, g *Graph) {
	for {
		select {
		case newUser := <-m.register:
			if correct := g.AddAloneNode(newUser.Id); correct {
				users[newUser.Id] = newUser
			} else {
				log.Printf("Error user %s is already on the system", newUser.Id)
			}
		case newRel := <-m.relation:
			if correct := g.AddEdge(newRel.Scanned.Id, newRel.Scans.Id); !correct {
				log.Printf("Error relation %s -> %s already on the system", newRel.Scanned.Id, newRel.Scans.Id)
			}
		case op := <-m.debug:
			if op == 0 {
				return
			} else if op == 1 {
				g.Log()
			} else if op == 2 {
				g.ToJson("tmp/graph.json")
			} else if op == 3 {
				g.FromJson("tmp/graph.json")
			} else {
				log.Printf("Unknown operation %d", op)
			}
		}
	}
}
