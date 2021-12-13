package circuits

type CircuitsResp struct {
	MRData MRData `json:"MRData"`
}
type Location struct {
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Locality string `json:"locality"`
	Country  string `json:"country"`
}
type Circuits struct {
	CircuitID   string   `json:"circuitId"`
	URL         string   `json:"url"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"Location"`
}
type CircuitTable struct {
	Season   string     `json:"season"`
	Circuits []Circuits `json:"Circuits"`
}
type MRData struct {
	Xmlns        string       `json:"xmlns"`
	Series       string       `json:"series"`
	URL          string       `json:"url"`
	Limit        string       `json:"limit"`
	Offset       string       `json:"offset"`
	Total        string       `json:"total"`
	CircuitTable CircuitTable `json:"CircuitTable"`
}
