package musicutil

import (
	"github.com/graphql-go/graphql"
)

var (
	Illmatic Disc
	CharityStartsAtHome Disc

	DiscData   map[int]Disc

	discType   *graphql.Object

	MusicSchema graphql.Schema
)

type Disc struct {
	Id     string
	Title  string
	Artist string
	Year   int
}

func init() {
	Illmatic = Disc{
		Title:  "Illmatic",
		Artist: "Nas",
		Year:   1994,
		Id:     "1",
	}
	CharityStartsAtHome = Disc{
		Title:  "CharityStartsAtHome",
		Artist: "Phonte",
		Year:   2011,
		Id:     "2",
	}

	DiscData = map[int]Disc{
		1:  Illmatic,
		2:  CharityStartsAtHome,
	}

	discType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Disc",
		Description: "A set of songs from one or many artists.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The Identifier of the disc.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Id, nil
					}
					return nil, nil
				},
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "The Title of the album.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Title, nil
					}
					return nil, nil
				},
			},
			"artist": &graphql.Field{
				Type:        graphql.String,
				Description: "The Artist of the album.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Artist, nil
					}
					return nil, nil
				},
			},
			"year": &graphql.Field{
				Type:        graphql.Int,
				Description: "The release year.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disc, ok := p.Source.(Disc); ok {
						return disc.Year, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
				  "discs": &graphql.Field{
				Type: graphql.NewList(discType), // we return a list of discType, defined above
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetAllDiscs(), nil // <-- every time this query is called, return the result of GetAllDiscs()
				},
			},
		},
	})

	MusicSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
    		// mutation will be added later
	})
}

//GetAllDiscs return all the dummy data
func GetAllDiscs() []Disc {
	discs := []Disc{}
	for _, disc := range DiscData {
		discs = append(discs, disc)
	}
	return discs
}