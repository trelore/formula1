package drivers

type DriversResp struct {
	MRData MRData `json:"MRData"`
}
type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}
type Constructors struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}
type DriverStandings struct {
	Position     string         `json:"position"`
	PositionText string         `json:"positionText"`
	Points       string         `json:"points"`
	Wins         string         `json:"wins"`
	Driver       Driver         `json:"Driver"`
	Constructors []Constructors `json:"Constructors"`
}
type StandingsLists struct {
	Season          string            `json:"season"`
	Round           string            `json:"round"`
	DriverStandings []DriverStandings `json:"DriverStandings"`
}
type StandingsTable struct {
	Season         string           `json:"season"`
	StandingsLists []StandingsLists `json:"StandingsLists"`
}
type MRData struct {
	Xmlns          string         `json:"xmlns"`
	Series         string         `json:"series"`
	URL            string         `json:"url"`
	Limit          string         `json:"limit"`
	Offset         string         `json:"offset"`
	Total          string         `json:"total"`
	StandingsTable StandingsTable `json:"StandingsTable"`
}
