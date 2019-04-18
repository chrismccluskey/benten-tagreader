package models

type File struct {
	ID   string `json:"id"`
	Name string `json:"string" validate:"required"`
	Size int64  `json:"size" validate:"required"`
}
