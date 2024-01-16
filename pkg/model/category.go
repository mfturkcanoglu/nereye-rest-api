package model

type Category struct {
	Category string
	PhotoUrl string
	DefaultModel
}

type CategoryGet struct {
	Id       string  `json:"id"`
	Category string  `json:"category"`
	PhotoUrl *string `json:"photo_url"`
}

type CategoryCreate struct {
	Category string  `json:"category"`
	PhotoUrl *string `json:"photo_url"`
}
