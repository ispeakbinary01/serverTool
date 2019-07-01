package structs

// Server ...
type Server struct {
	IP        int         `json:"id"`
	Os        string      `json:"os"`
	Softwares *[]Software `json:"software"`
	SSHs      *[]SSH      `json:"ssh"`
}
