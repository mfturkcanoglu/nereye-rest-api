package model

type Address struct {
	Country     string
	City        string
	County      string
	Distric     string
	FullAddress string
	Latitude    float64
	Longitude   float64
	DefaultModel
}
