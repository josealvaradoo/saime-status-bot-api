package model

const (
	Online  = "online"
	Offline = "offline"
)

type Saime struct {
	Status string `json:"status,omitempty"`
}
