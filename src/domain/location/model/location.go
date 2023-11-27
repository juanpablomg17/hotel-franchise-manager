package model

type Location struct {
	Id      string `json:"id"`
	City    string `json:"city"`
	Country string `json:"country"`
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
}
