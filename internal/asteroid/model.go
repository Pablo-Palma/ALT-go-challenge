package asteroid

import "go.mongodb.org/mongo-driver/bson/primitive"

type Asteroid struct {
	ID					primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Name				string				`json:"name"`
	Diameter			float64				`json:"diameter"`
	DiscoveryDate		string				`json:"discovery_date"`
	Observations		string				`json:"observations, omitempty"`
	Distances			[]Distance			`json:"distances, omitempty"`
}

type Distance struct {
	Date		string `json:"date"`
	Distance	float64 `json:"distance"`
}
