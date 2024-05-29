package asteroid

type Asteroid struct {
	ID					string	`json:"id"`
	Name				string	`json:"name"`
	Diameter			string	`json:"diameter"`
	DiscoveryDate		string	`json:"discovery_date!"`
	ObservationCount	string	`json:"observation_count"`
	DistanceFromEarth	string	`json:"distance_from_earth"`
}
