package main

import (
	"github.com/graphql-go/graphql"
)

type Node interface {
	GetID() string
	GetType() string
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p *Person) GetID() string {
	return p.ID
}

func (p *Person) GetType() string {
	return "Person"
}

type Story struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

func (s *Story) GetID() string {
	return s.ID
}

func (s *Story) GetType() string {
	return "Story"
}

var nodes = []Node{
	&Person{
		ID:   "the-viewer",
		Name: "A. D. Veloper",
	},
	&Story{
		ID:       "2",
		Title:    "Local Yak Named Yak of the Year",
		Category: "ALL",
	},
}

func nodeResolver(id string) Node {
	for _, n := range nodes {
		if n.GetID() == id {
			return n
		}
	}
	return nil
}

func nodeInterfaceResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	return nodeResolver(id), nil
}
