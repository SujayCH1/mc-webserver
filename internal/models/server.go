package models

type Server struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Port    string `json:"port"`
	Status  string `json:"status"`
}
