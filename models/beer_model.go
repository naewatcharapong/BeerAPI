package models

type BeerModel struct {
	ID     uint
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Images string `json:"images"`
}
