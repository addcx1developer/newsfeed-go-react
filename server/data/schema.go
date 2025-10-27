package data

import (
	"github.com/graphql-go/graphql"
)

var (
	nodeInterface  *graphql.Interface
	categoryType   *graphql.Enum
	imageType      *graphql.Object
	actorInterface *graphql.Interface
	personType     *graphql.Object
	storyType      *graphql.Object
	viewerType     *graphql.Object
	queryType      *graphql.Object
	Schema         graphql.Schema
)

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
			"EDUCATION": &graphql.EnumValueConfig{
				Value: "EDUCATION",
			},
			"NEWS": &graphql.EnumValueConfig{
				Value: "NEWS",
			},
		},
	})

	imageType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Image",
		Fields: graphql.Fields{
			"url": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Args: graphql.FieldConfigArgument{
					"width": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"height": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: imageURLResolver,
			},
			"altText": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	actorInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Actor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"profilePicture": &graphql.Field{
				Type: imageType,
			},
			"joined": &graphql.Field{
				Type: graphql.String,
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *Person:
				return personType
			default:
				return nil
			}
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
			"profilePicture": &graphql.Field{
				Type: imageType,
			},
			"joined": &graphql.Field{
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
			"thumbnail": &graphql.Field{
				Type: imageType,
			},
			"summary": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: categoryType,
			},
			"poster": &graphql.Field{
				Type:    graphql.NewNonNull(actorInterface),
				Resolve: storyPosterResolver,
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Interfaces: []*graphql.Interface{nodeInterface},
	})

	viewerType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Viewer",
		Fields: graphql.Fields{
			"actor": &graphql.Field{
				Type: actorInterface,
			},
			"contacts": &graphql.Field{
				Type: graphql.NewList(actorInterface),
				Args: graphql.FieldConfigArgument{
					"search": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: contactsResolver,
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"viewer": &graphql.Field{
				Type:    viewerType,
				Resolve: viewerResolver,
			},
			"node": &graphql.Field{
				Type: nodeInterface,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: nodeResolver,
			},
			"topStory": &graphql.Field{
				Type: storyType,
				Args: graphql.FieldConfigArgument{
					"category": &graphql.ArgumentConfig{
						Type: categoryType,
					},
				},
				Resolve: topStoryResolver,
			},
			"topStories": &graphql.Field{
				Type:    graphql.NewList(storyType),
				Resolve: topStoriesResolver,
			},
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Types: []graphql.Type{
			personType,
			storyType,
		},
	})

	if err != nil {
		panic(err)
	}
}
