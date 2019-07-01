package structs

// Server ...
type Server struct {
	IP       int        `json:"id"`
	Os       string     `json:"os"`
	Software []Software `json:"software"`
	SSH      []SSH      `json:"ssh"`
}
