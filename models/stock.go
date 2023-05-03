package models

type Stock struct {
	Id     int32 `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
