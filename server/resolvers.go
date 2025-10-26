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
	&Story{
		ID:       "3",
		Title:    "Why did the chicken cross the road? To get to the other side!",
		Category: "EDUCATION",
	},
	&Story{
		ID:       "4",
		Title:    "New Hedgehog Species Discovered",
		Category: "NEWS",
	},
	&Story{
		ID:       "story6",
		Title:    "What is my cat trying to tell me?",
		Category: "ALL",
	},
}

type Viewer struct {
	Actor *Person
}

func viewerResolver(p graphql.ResolveParams) (interface{}, error) {
	for _, n := range nodes {
		if person, ok := n.(*Person); ok && person.GetID() == "the-viewer" {
			return &Viewer{Actor: person}, nil
		}
	}
	return nil, nil
}

func nodeResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)

	for _, n := range nodes {
		if n.GetID() == id {
			return n, nil
		}
	}

	return nil, nil
}

func topStoryResolver(p graphql.ResolveParams) (interface{}, error) {
	categoryArg, ok := p.Args["category"].(string)

	for _, node := range nodes {
		story, isStory := node.(*Story)
		if !isStory {
			continue
		}

		if ok && categoryArg != "ALL" {
			if story.Category == categoryArg {
				return story, nil
			}
		} else {
			return story, nil
		}
	}

	return nil, nil
}

func TopStoriesResolver(p graphql.ResolveParams) (interface{}, error) {
	var stories []*Story

	for _, node := range nodes {
		story, ok := node.(*Story)
		if !ok {
			continue
		}
		stories = append(stories, story)
	}

	if len(stories) > 3 {
		stories = stories[:3]
	}

	return stories, nil
}
