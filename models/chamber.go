package model

type Chamber struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 string  `json:"addressLine2"`
	PostCode     string  `json:"postCode"`
	Lat          float32 `json:"lat"`
	Long         float32 `json:"long"`
	PhoneNumber  string  `json:"phoneNumber"`
}

func GetMockChambers() []Chamber {
	var chambers []Chamber = []Chamber{
		{
			Id:           1,
			Name:         "Test Chaber 1",
			AddressLine1: "Test Line 1",
			AddressLine2: "Test Line 2",
			PostCode:     "AB1 2CD",
			Lat:          50.1,
			Long:         45.2,
			PhoneNumber:  "0800908765",
		},
		{
			Id:           2,
			Name:         "Test Chaber 1",
			AddressLine1: "Test Line 1",
			AddressLine2: "Test Line 2",
			PostCode:     "CD2 2EF",
			Lat:          51.1,
			Long:         34.1,
			PhoneNumber:  "0800909778",
		},
	}

	return chambers
}
