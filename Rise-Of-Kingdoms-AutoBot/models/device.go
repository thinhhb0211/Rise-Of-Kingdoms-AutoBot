package models

type Device struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}
