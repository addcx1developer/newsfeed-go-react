package main

import "github.com/graphql-go/graphql"

var nodeInterface *graphql.Interface
var categoryType *graphql.Enum
var actorInterface *graphql.Interface
var personType *graphql.Object
var storyType *graphql.Object
var queryType *graphql.Object

func init() {
	nodeInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Node",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *Person:
				return personType
			case *Story:
				return storyType
			default:
				return nil
			}
		},
	})

	categoryType = graphql.NewEnum(graphql.EnumConfig{
		Name: "Category",
		Values: graphql.EnumValueConfigMap{
			"ALL": &graphql.EnumValueConfig{
				Value: "ALL",
			},
		},
	})

	actorInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Actor",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	personType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			nodeInterface,
			actorInterface,
		},
	})

	storyType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Story",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"category": &graphql.Field{
				Type: graphql.NewNonNull(categoryType),
			},
		},
		Interfaces: []*graphql.Interface{nodeInterface},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": &graphql.Field{
				Type: nodeInterface,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: nodeInterfaceResolver,
			},
		},
	})
}
