package data

import (
	"fmt"
	"net/url"

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

type Image struct {
	URL     string `json:"url"`
	AltText string `json:"altText"`
}

type Story struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Thumbnail *Image `json:"thumbnail"`
	Summary   string `json:"summary"`
	Category  string `json:"category"`
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
		ID:    "2",
		Title: "Local Yak Named Yak of the Year",
		Thumbnail: &Image{
			URL:     "/assets/yak.png",
			AltText: "Portrait of Max the Yak by a local artist",
		},
		Summary:  "The annual Yak of the Year awards ceremony took place last night, and this year's winner is none other than Max, a beloved yak from the small town of Millville. Max, who is known for his friendly personality and hardworking nature, beat out stiff competition from other yaks in the region to take home the coveted title.\n \nAccording to the judges, Max stood out due to his exceptional contributions to the community. He has been used as a pack animal to help transport goods to and from the town's market, and has also been a reliable source of milk and wool for local farmers. In addition, Max has become something of a local celebrity, often posing for photos with tourists and participating in community events.",
		Category: "ALL",
	},
	&Story{
		ID:       "3",
		Title:    "Why did the chicken cross the road? To get to the other side!",
		Summary:  "Chickens are curious animals and will often explore their surroundings, including crossing roads if the opportunity arises. It is important to note that chickens are intelligent and adaptable animals, and the specific reasons for any given chicken crossing the road may vary depending on the individual and its circumstances.",
		Category: "EDUCATION",
	},
	&Story{
		ID:       "4",
		Title:    "New Hedgehog Species Discovered",
		Summary:  "Breaking news! Scientists have just announced the discovery of a new species of hedgehog, and you won't believe what makes this species unique.\n \n     According to the researchers, the new hedgehogs, which have been named 'sparklehogs,' are distinguished by their ability to produce rainbow-colored sparks from their spikes when they are feeling threatened.\n     \n     But that's not all! The sparklehogs have also been observed using their sparkling spikes to communicate with one another, creating dazzling light shows in the process.\n     \n     'We've never seen anything like it,' said lead researcher Dr. Maria Hernandez. 'These hedgehogs are truly one of a kind.'",
		Category: "NEWS",
	},
	&Story{
		ID:       "story6",
		Title:    "What is my cat trying to tell me?",
		Summary:  "I just had the most surreal conversation with my cat. I was talking to her about my day and she just stared at me with this really intense look in her eyes and meowed. I have no idea what she was trying to say, but I have a feeling she was trying to impart some deep wisdom on me. Or maybe she just wanted more treats. Either way, it was a very interesting conversation. #catconvo #felinewisdom",
		Category: "ALL",
	},
}

type Viewer struct {
	Actor *Person
}

type ImageArgs struct {
	URL    string
	Width  *int
	Height *int
}

func buildImageURL(args ImageArgs) string {
	u, err := url.Parse(args.URL)
	if err != nil {
		return args.URL
	}

	q := u.Query()
	if args.Width != nil {
		q.Set("width", fmt.Sprintf("%d", *args.Width))
	}
	if args.Height != nil {
		q.Set("height", fmt.Sprintf("%d", *args.Height))
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func imageURLResolver(p graphql.ResolveParams) (interface{}, error) {
	image, ok := p.Source.(*Image)
	if !ok || image == nil {
		return nil, nil
	}

	var widthPtr, heightPtr *int
	if w, ok := p.Args["width"].(int); ok {
		widthPtr = &w
	}
	if h, ok := p.Args["height"].(int); ok {
		heightPtr = &h
	}

	return buildImageURL(ImageArgs{
		URL:    image.URL,
		Width:  widthPtr,
		Height: heightPtr,
	}), nil
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

func topStoriesResolver(p graphql.ResolveParams) (interface{}, error) {
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
