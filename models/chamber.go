package model

type Chamber struct {
	Id           int
	Name         string
	AddressLine1 string
	AddressLine2 string
	PostCode     string
	Lat          float32
	Long         float32
	PhoneNumber  string
}

func (chamber Chamber) GetMockChambers() []Chamber {
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
