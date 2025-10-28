package data

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

type Node interface {
	GetID() string
	GetType() string
}

type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	ProfilePicture *Image    `json:"profilePicture"`
	Location       *Location `json:"location"`
	Joined         string    `json:"joined"`
}

func (p *Person) GetID() string {
	return p.ID
}

func (p *Person) GetType() string {
	return "Person"
}

type Organization struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ProfilePicture   *Image `json:"profilePicture"`
	OrganizationKind string `json:"organizationKind"`
	Joined           string `json:"joined"`
}

func (o *Organization) GetID() string {
	return o.ID
}

func (o *Organization) GetType() string {
	return "Organization"
}

type Image struct {
	URL     string `json:"url"`
	AltText string `json:"altText"`
}

type Story struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Thumbnail *Image     `json:"thumbnail"`
	Summary   string     `json:"summary"`
	Category  string     `json:"category"`
	AuthorID  string     `json:"authorID"`
	Comments  []*Comment `json:"comments"`
	CreatedAt string     `json:"createdAt"`
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
		ProfilePicture: &Image{
			URL: "/assets/a.png",
		},
		Location: &Location{
			ID:   "viewer-location",
			Name: "Wheresoever You Are",
		},
		Joined: "2025-10-27T00:00:00.000Z",
	},
	&Person{
		ID:   "1",
		Name: "Chris P. Bacon",
		ProfilePicture: &Image{
			URL: "/assets/pig.png",
		},
		Location: &Location{
			ID:   "10",
			Name: "Poultropolis",
		},
		Joined: "2025-10-27T00:00:00.000Z",
	},
	&Person{
		ID:   "8",
		Name: "Suzie Queue",
		ProfilePicture: &Image{
			URL: "/assets/blogger.png",
		},
		Location: &Location{
			ID:   "11",
			Name: "Venice",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "9",
		Name: "Katrina Starer",
		ProfilePicture: &Image{
			URL: "/assets/cat_avatar.png",
		},
		Location: &Location{
			ID:   "12",
			Name: "New York",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "13",
		Name: "Alexei Seligsteinwitz",
		ProfilePicture: &Image{
			URL: "/assets/alexei.png",
		},
		Location: &Location{
			ID:   "14",
			Name: "Mystery",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "15",
		Name: "Jennifer Letuchyberg",
		ProfilePicture: &Image{
			URL: "/assets/j.png",
		},
		Location: &Location{
			ID:   "16",
			Name: "Zig Zag",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "17",
		Name: "Paige Talaberg",
		ProfilePicture: &Image{
			URL: "/assets/p.png",
		},
		Location: &Location{
			ID:   "18",
			Name: "Boring",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "19",
		Name: "Catalina Chaiman",
		ProfilePicture: &Image{
			URL: "/assets/c.png",
		},
		Location: &Location{
			ID:   "20",
			Name: "Weed",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "21",
		Name: "Huy Li",
		ProfilePicture: &Image{
			URL: "/assets/h.png",
		},
		Location: &Location{
			ID:   "22",
			Name: "Cave Junction",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Person{
		ID:   "30",
		Name: "Brock Boc",
		ProfilePicture: &Image{
			URL: "/assets/chicken_lover.png",
		},
		Location: &Location{
			ID:   "31",
			Name: "Paradoxopolus",
		},
		Joined: "2025-10-28T00:00:00.000Z",
	},
	&Organization{
		ID:   "6",
		Name: "Gazelle Gazette",
		ProfilePicture: &Image{
			URL: "/assets/gazelle.png",
		},
		OrganizationKind: "JOURNALISTIC",
		Joined:           "2025-10-28T00:00:00.000Z",
	},
	&Organization{
		ID:   "7",
		Name: "Baller Bovine Board",
		ProfilePicture: &Image{
			URL:     "/assets/bovine.png",
			AltText: "Blue-ribboned insignia of the Baller Bovine Board",
		},
		OrganizationKind: "NONPROFIT",
		Joined:           "2025-10-27T00:00:00.000Z",
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
		AuthorID: "7",
		Comments: []*Comment{
			{
				ID:   "comment1",
				Text: "So proud of our local yak",
			},
			{
				ID:   "comment2",
				Text: "I've been waiting my whole life for this moment",
			},
			{
				ID:   "comment3",
				Text: "What's a yak???",
			},
			{
				ID:   "comment4",
				Text: "We used to keep yaks in the old country",
			},
			{
				ID:   "comment5",
				Text: "It's a yak attack, this award is whack",
			},
			{
				ID:   "comment6",
				Text: "There are better yaks in every pasture of this country",
			},
			{
				ID:   "comment7",
				Text: "Yak yak yak yak yak yak yak yak",
			},
			{
				ID:   "comment8",
				Text: "He's a good yak, he saved my child from drowning",
			},
		},
		CreatedAt: "2025-10-27T00:00:00.000Z",
	},
	&Story{
		ID:    "3",
		Title: "Why did the chicken cross the road? To get to the other side!",
		Thumbnail: &Image{
			URL: "/assets/chicken.png",
		},
		Summary:  "Chickens are curious animals and will often explore their surroundings, including crossing roads if the opportunity arises. It is important to note that chickens are intelligent and adaptable animals, and the specific reasons for any given chicken crossing the road may vary depending on the individual and its circumstances.",
		Category: "EDUCATION",
		AuthorID: "1",
		Comments: []*Comment{
			{
				ID:   "comment9",
				Text: "I never knew!",
			},
			{
				ID:   "comment10",
				Text: "This is a very deep joke.",
			},
		},
		CreatedAt: "2025-10-27T00:00:00.000Z",
	},
	&Story{
		ID:    "4",
		Title: "New Hedgehog Species Discovered",
		Thumbnail: &Image{
			URL: "/assets/hedgehog.png",
		},
		Summary:  "Breaking news! Scientists have just announced the discovery of a new species of hedgehog, and you won't believe what makes this species unique.\n \n     According to the researchers, the new hedgehogs, which have been named 'sparklehogs,' are distinguished by their ability to produce rainbow-colored sparks from their spikes when they are feeling threatened.\n     \n     But that's not all! The sparklehogs have also been observed using their sparkling spikes to communicate with one another, creating dazzling light shows in the process.\n     \n     'We've never seen anything like it,' said lead researcher Dr. Maria Hernandez. 'These hedgehogs are truly one of a kind.'",
		Category: "NEWS",
		AuthorID: "6",
		Comments: []*Comment{
			{
				ID:   "comment11",
				Text: "Aren't hedges dark though???",
			},
		},
		CreatedAt: "2025-10-27T00:00:00.000Z",
	},
	&Story{
		ID:    "5",
		Title: "Onion Soup Recipe",
		Thumbnail: &Image{
			URL: "/assets/recipe.png",
		},
		Summary:  "I am so excited to share with you my all-time favorite recipe for French onion soup. I can't even begin to tell you how many times I've made this dish for my family and friends, and it never fails to impress.\n\n As a self-proclaimed wine mom, I always love finding new and creative ways to incorporate my favorite vintages into my cooking. And let me tell you, the dry white wine in this recipe really takes the flavor of the onions to the next level. Trust me, it's a game changer.\n \n But don't just take my word for it – give this recipe a try for yourself and see how it becomes a new staple in your household. Not only is it delicious, but it's also the perfect comfort food for those cold winter nights.\n \n So grab your wine glasses and let's get cooking!",
		Category: "COOKING",
		AuthorID: "8",
		Comments: []*Comment{
			{
				ID:   "comment12",
				Text: "I tried it with passionfruit instead of onions, it's a great substitution!",
			},
		},
		CreatedAt: "2025-10-28T00:00:00.000Z",
	},
	&Story{
		ID:    "99",
		Title: "Study: The egg came first, but only after the chicken",
		Thumbnail: &Image{
			URL: "/assets/puzzled_egg.png",
		},
		Summary:  `In a shocking new study, scientists have finally determined the age-old question of whether the chicken or the egg came first. And it turns out, the answer is both!\n      According to the research, the egg actually came first — but only after the chicken had already laid it.\n      \n      "We were amazed by the results," said lead researcher Dr. Janet Hennessy. "It seems that the chicken somehow managed to lay an egg before it even existed. It\'s a real chicken-and-egg paradox."\n      \n      The study, which involved observing hundreds of chickens on a farm, found that the birds would lay eggs and then, a short time later, a fully-formed chicken would emerge from the shell.\n      \n      "We always thought that the egg came first and the chicken was born from it," said Hennessy. "But it turns out, the chicken was there all along, just waiting to hatch."\n      \n      The findings have caused quite a stir in the scientific community, with many experts calling for further research to be done on the mysterious life cycle of the chicken.\n      \n      "It\'s a groundbreaking discovery that will change the way we think about the chicken and the egg," said Hennessy. "Who knows what other secrets these amazing creatures may be hiding?"`,
		Category: "NEWS",
		AuthorID: "30",
		Comments: []*Comment{
			{
				ID:   "comment13",
				Text: "Wait...",
			},
		},
		CreatedAt: "2025-10-28T00:00:00.000Z",
	},
	&Story{
		ID:    "story6",
		Title: "What is my cat trying to tell me?",
		Thumbnail: &Image{
			URL: "/assets/cat.png",
		},
		Summary:   "I just had the most surreal conversation with my cat. I was talking to her about my day and she just stared at me with this really intense look in her eyes and meowed. I have no idea what she was trying to say, but I have a feeling she was trying to impart some deep wisdom on me. Or maybe she just wanted more treats. Either way, it was a very interesting conversation. #catconvo #felinewisdom",
		Category:  "ALL",
		AuthorID:  "9",
		CreatedAt: "2025-10-27T00:00:00.000Z",
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

type Comment struct {
	ID   string `json:"id"`
	Text string `json:"text"`
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

func storyPosterResolver(p graphql.ResolveParams) (interface{}, error) {
	story, ok := p.Source.(*Story)
	if !ok || story == nil {
		return nil, nil
	}

	return nodeResolver(graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": story.AuthorID,
		},
	})
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
		if story, ok := node.(*Story); ok {
			stories = append(stories, story)
		}
	}

	if len(stories) > 3 {
		stories = stories[:3]
	}

	if stories == nil {
		stories = []*Story{}
	}

	return stories, nil
}

func newsfeedStoriesResolver(p graphql.ResolveParams) (interface{}, error) {
	first, _ := p.Args["first"].(int)
	afterStr, _ := p.Args["after"].(string)
	category, _ := p.Args["category"].(string)

	count := first
	if count == 0 {
		count = len(nodes)
	}

	after := 0
	if afterStr != "" {
		if n, err := strconv.Atoi(afterStr); err == nil {
			after = n
		}
	}

	var filtered []*Story
	for _, n := range nodes {
		story, ok := n.(*Story)
		if !ok {
			continue
		}
		if category == "" || category == "ALL" || story.Category == category {
			filtered = append(filtered, story)
		}
	}

	next := after + count
	if next > len(filtered) {
		next = len(filtered)
	}

	edges := make([]map[string]interface{}, next-after)
	for i, story := range filtered[after:next] {
		edges[i] = map[string]interface{}{
			"node":   story,
			"cursor": story.ID,
		}
	}

	pageInfo := map[string]interface{}{
		"hasNextPage": next < len(filtered),
		"endCursor":   strconv.Itoa(next),
	}

	return map[string]interface{}{
		"edges":    edges,
		"pageInfo": pageInfo,
	}, nil
}

func contactsResolver(p graphql.ResolveParams) (interface{}, error) {
	searchArg, _ := p.Args["search"].(string)
	searchArg = strings.ToLower(searchArg)

	var persons []*Person

	for _, node := range nodes {
		person, ok := node.(*Person)
		if !ok || person.ID == "the-viewer" {
			continue
		}
		persons = append(persons, person)
	}

	if searchArg == "" {
		return persons, nil
	}

	var filtered []*Person
	for _, person := range persons {
		if strings.Contains(strings.ToLower(person.Name), searchArg) {
			filtered = append(filtered, person)
		}
	}

	return filtered, nil
}

func storyCommentsResolver(p graphql.ResolveParams) (interface{}, error) {
	story, ok := p.Source.(*Story)
	if !ok || story == nil {
		return nil, nil
	}

	first, firstOK := p.Args["first"].(int)
	afterStr, afterOK := p.Args["after"].(string)

	after := 0
	if afterOK {
		if a, err := strconv.Atoi(afterStr); err == nil {
			after = a
		}
	}

	count := len(story.Comments)
	if firstOK {
		count = first
	}

	next := after + count
	if next > len(story.Comments) {
		next = len(story.Comments)
	}

	edges := make([]map[string]interface{}, next-after)
	for i, c := range story.Comments[after:next] {
		edges[i] = map[string]interface{}{
			"node":   c,
			"cursor": strconv.Itoa(after + i + 1),
		}
	}

	pageInfo := map[string]interface{}{
		"hasNextPage": next < len(story.Comments),
		"endCursor":   strconv.Itoa(next),
	}

	return map[string]interface{}{
		"edges":    edges,
		"pageInfo": pageInfo,
	}, nil
}
