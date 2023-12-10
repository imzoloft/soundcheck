package model

type Response struct {
	Error      string `json:"error"`
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}
