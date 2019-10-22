package main

import (
	"log"
)

type Node struct {
	id    string
	edges []string
}

func (n Node) Id() string {
	return n.id
}

func (n Node) Edges() []string {
	return n.edges
}

type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

func (g Graph) AddNode(id string, edges []string) bool {
	if _, pr := g.nodes[id]; pr {
		return false
	}
	g.nodes[id] = &Node{
		id:    id,
		edges: edges,
	}
	return true
}
func (g Graph) AddAloneNode(id string) bool {
	if _, pr := g.nodes[id]; pr {
		return false
	}
	g.nodes[id] = &Node{
		id:    id,
		edges: make([]string, 0),
	}
	return true
}
func IsPresent(edges []string, n string) bool {
	for _, v := range edges {
		if v == n {
			return true
		}
	}
	return false
}

func (g Graph) AddEdge(from string, to string) bool {
	if found := IsPresent(g.nodes[from].edges, to); found {
		return false
	}
	g.nodes[from].edges = append(g.nodes[from].edges, to)
	return true
}

func (g Graph) Log() {
	for _, v := range g.nodes {
		log.Printf("Node %s:", v.id)
		for _, c := range v.edges {
			log.Printf(" %s", c)
		}
		log.Printf("")
	}
}
