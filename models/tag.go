package models

type Tag struct {
	ID    string `json:"id"`
	Type  string `json:"string" validate:"required"`
	Value string `json:"string" validate:"required"`
}
