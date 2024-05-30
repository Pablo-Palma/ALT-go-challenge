package asteroid

type Asteroid struct {
	ID					string		`json:"id"`
	Name				string		`json:"name"`
	Diameter			float64		`json:"diameter"`
	DiscoveryDate		string		`json:"discovery_date"`
	Observations		string		`json:"observations, omitempty"`
	Distances			[]Distance	`json:"distances, omitempty"`
}

type Distance struct {
	Date		string `json:"date"`
	Distance	float64 `json:"distance"`
}
