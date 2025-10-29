package data

import (
	"github.com/graphql-go/graphql"
)

var (
	nodeInterface                    *graphql.Interface
	categoryType                     *graphql.Enum
	organizationKindType             *graphql.Enum
	imageType                        *graphql.Object
	locationType                     *graphql.Object
	actorInterface                   *graphql.Interface
	personType                       *graphql.Object
	organizationType                 *graphql.Object
	commentType                      *graphql.Object
	pageInfoType                     *graphql.Object
	storyType                        *graphql.Object
	viewerType                       *graphql.Object
	queryType                        *graphql.Object
	storyMutationResponseType        *graphql.Object
	storyCommentMutationResponseType *graphql.Object
	mutationType                     *graphql.Object
	Schema                           graphql.Schema
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
			case *Organization:
				return organizationType
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
			"COOKING": &graphql.EnumValueConfig{
				Value: "COOKING",
			},
		},
	})

	organizationKindType = graphql.NewEnum(graphql.EnumConfig{
		Name: "OrganizationKind",
		Values: graphql.EnumValueConfigMap{
			"NONPROFIT": &graphql.EnumValueConfig{
				Value: "NONPROFIT",
			},
			"JOURNALISTIC": &graphql.EnumValueConfig{
				Value: "JOURNALISTIC",
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

	locationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
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
			case *Organization:
				return organizationType
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
			"location": &graphql.Field{
				Type: locationType,
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

	organizationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Organization",
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
			"organizationKind": &graphql.Field{
				Type: organizationKindType,
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

	commentType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	pageInfoType = graphql.NewObject(graphql.ObjectConfig{
		Name: "PageInfo",
		Fields: graphql.Fields{
			"startCursor": &graphql.Field{
				Type: graphql.String,
			},
			"endCursor": &graphql.Field{
				Type: graphql.String,
			},
			"lastCursor": &graphql.Field{
				Type: graphql.String,
			},
			"hasNextPage": &graphql.Field{
				Type: graphql.Boolean,
			},
			"hasPreviousPage": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

	createConnectionType := func(name string, nodeType *graphql.Object) (connectionType, edgeType *graphql.Object) {
		edgeType = graphql.NewObject(graphql.ObjectConfig{
			Name: name + "Edge",
			Fields: graphql.Fields{
				"node": &graphql.Field{
					Type: nodeType,
				},
				"cursor": &graphql.Field{
					Type: graphql.String,
				},
			},
		})

		connectionType = graphql.NewObject(graphql.ObjectConfig{
			Name: name + "Connection",
			Fields: graphql.Fields{
				"edges": &graphql.Field{
					Type: graphql.NewList(edgeType),
				},
				"pageInfo": &graphql.Field{
					Type: pageInfoType,
				},
			},
		})

		return connectionType, edgeType
	}

	commentsConnectionType, commentsConnectionEdgeType := createConnectionType("comments", commentType)

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
			"likeCount": &graphql.Field{
				Type: graphql.Int,
			},
			"doesViewerLike": &graphql.Field{
				Type: graphql.Boolean,
			},
			"comments": &graphql.Field{
				Type: commentsConnectionType,
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"after": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: storyCommentsResolver,
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Interfaces: []*graphql.Interface{nodeInterface},
	})

	storiesConnectionType, _ := createConnectionType("stories", storyType)

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
			"newsfeedStories": &graphql.Field{
				Type: storiesConnectionType,
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"after": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: categoryType,
					},
				},
				Resolve: newsfeedStoriesResolver,
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

	storyMutationResponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "StoryMutationResponse",
		Fields: graphql.Fields{
			"story": &graphql.Field{
				Type: storyType,
			},
		},
	})

	storyCommentMutationResponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "StoryCommentMutationResponse",
		Fields: graphql.Fields{
			"story": &graphql.Field{
				Type: storyType,
			},
			"commentEdge": &graphql.Field{
				Type: commentsConnectionEdgeType,
			},
		},
	})

	mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"likeStory": &graphql.Field{
				Type: storyMutationResponseType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"doesLike": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: resolveLikeStoryMutation,
			},
			"postStoryComment": &graphql.Field{
				Type: storyCommentMutationResponseType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvePostStoryCommentMutation,
			},
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
		Types: []graphql.Type{
			personType,
			organizationType,
			storyType,
		},
	})

	if err != nil {
		panic(err)
	}
}
