package models

type Pokemon struct {
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Name   string `json:"name"`
	Id     int    `json:"id"`
}
