package model

type RestaurantPhoto struct {
	Name       *string
	Url        string
	ParentID   string
	PhotoOrder *int16
	DefaultModel
}

type RestaurantPhotoCreate struct {
	Name       *string `json:"name"`
	Url        string  `json:"url"`
	ParentID   string  `json:"parent_id"`
	PhotoOrder *int16  `json:"photo_order"`
}

type RestaurantPhotoGet struct {
	Name       *string `json:"name"`
	Url        string  `json:"url"`
	ParentID   string  `json:"parent_id"`
	PhotoOrder *int16  `json:"photo_order"`
}
