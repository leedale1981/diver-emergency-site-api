package model

type DiveSite struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 string  `json:"addressLine2"`
	PostCode     string  `json:"postCode"`
	Lat          float32 `json:"lat"`
	Long         float32 `json:"long"`
	PhoneNumber  string  `json:"phoneNumber"`
}
