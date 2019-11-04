package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
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
	log.Print("loggiiiing")
	for _, v := range g.nodes {
		log.Printf("Node %s:", v.id)
		for _, c := range v.edges {
			log.Printf(" %s", c)
		}
		log.Printf("")
	}
}

func (g Graph) ToJson(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Print("Error writting to Json")
		log.Print("Error is: ", err.Error())
		return
	}
	defer file.Close()
	_, _ = file.WriteString("{\n\"nodes\": [\n")
	i := 0
	for _, v := range g.nodes {
		if i != 0 {
			_, _ = file.WriteString(",\n")
		}
		_, _ = file.WriteString("{\n\"id\": \"")
		_, _ = file.WriteString(v.id)
		_, _ = file.WriteString("\",\n\"label\":\"")
		_, _ = file.WriteString(v.id)
		_, _ = file.WriteString("\",\n\"size\":1\n}")
		i = i + 1
	}
	_, _ = file.WriteString("],\n\"edges\": [\n")
	i = 0
	for _, from := range g.nodes {
		for _, to := range from.edges {
			if i != 0 {
				_, _ = file.WriteString(",\n")
			}
			_, _ = file.WriteString("{\n\"id\":\"")
			_, _ = file.WriteString(strings.Join([]string{from.id, to}, ""))
			_, _ = file.WriteString("\",\n\"source\":\"")
			_, _ = file.WriteString(from.id)
			_, _ = file.WriteString("\",\n\"target\":\"")
			_, _ = file.WriteString(to)
			_, _ = file.WriteString("\"\n}")

			i = i + 1
		}
	}
	_, _ = file.WriteString("\n]\n}")
}

func (g Graph) FromJson(filename string) {
	var newG struct {
		Nodes []struct {
			Id    string `json:"id"`
			Label string `json:"label"`
			Size  int    `json:"size"`
		} `json:"nodes"`

		Edges []struct {
			Id     string `json:"id"`
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"edges"`
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		log.Print("cannot open json")
		log.Print("Error is: ", err.Error())
		return
	}
	d := json.NewDecoder(file)
	if err := d.Decode(&newG); err != nil {
		log.Print("Error parsing graph ")
		log.Print(err.Error())
		return
	}
	//delete current G
	for _, del := range g.nodes {
		st := del.id
		del = nil
		delete(g.nodes, st)
	}

	//fill from newG nodes
	for _, n := range newG.Nodes {
		g.AddAloneNode(n.Id)
	}

	//fill from newG relations
	for _, rel := range newG.Edges {
		g.AddEdge(rel.Source, rel.Target)
	}
}
