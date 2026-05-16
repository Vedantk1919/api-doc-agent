package models

import "sample_api_go/app/structures"

func SearchGST(input structures.GSTSearchRequest) structures.GSTSearchData {
	if input.Mobile != "" || input.Email != "" {
		return structures.GSTSearchData{
			Message:   "Existing company found from contact lookup",
			GSTNumber: input.GST,
			ExistingCompanies: []structures.ExistingCompany{
				{
					GLID:                 "1245789",
					GSTCompanyName:       "Sharma Precision Tools",
					SalesAssignedTo:      "Amit Verma (98123)",
					GSTCompanyLatitude:   "28.6139",
					GSTCompanyLongitude:  "77.2090",
					DistanceFromMyLoc:    2.4,
					CompanyCustType:      "Paid",
					CurrentMeetingStatus: "Follow Up",
				},
			},
			ResultType:        1,
			RequestSource:     input.Source,
			LegacyCompanyName: "Sharma Precision Tools",
		}
	}

	return structures.GSTSearchData{
		Message:   "GST details found",
		GSTNumber: input.GST,
		GSTCompany: &structures.GSTCompany{
			GSTCompanyName:      "Sharma Precision Tools Pvt Ltd",
			GSTCompanyAddress:   "Okhla Industrial Area, New Delhi",
			GSTNumber:           input.GST,
			GSTPinCode:          "110020",
			EnterpriseType:      "ME",
			BusinessNatureType:  "Manufacturer",
			GSTCompanyLatitude:  "28.5355",
			GSTCompanyLongitude: "77.3910",
			DistanceFromMyLoc:   6.8,
			GSTData: []structures.GSTDisplayData{
				{DataKey: "GST", Value: input.GST},
				{DataKey: "Business Name", Value: "Sharma Precision Tools Pvt Ltd"},
				{DataKey: "Location", Value: "New Delhi"},
			},
		},
		ResultType:        2,
		RequestSource:     input.Source,
		LegacyCompanyName: "Sharma Precision Tools Pvt Ltd",
	}
}

func GetGSTRegistration(gst string) structures.GSTRegistrationDetails {
	return structures.GSTRegistrationDetails{
		BuildingName:           "Plot 21",
		BusinessActivityNature: "Factory / Manufacturing",
		BusinessAddressAdd:     "Okhla Industrial Area, Phase II",
		BusinessConstitution:   "Private Limited Company",
		BusinessName:           "Sharma Precision Tools Pvt Ltd",
		DateOfFiling:           "2026-04-20 12:30:00",
		FilingLastUpdationDate: "2026-04-22 09:15:00",
		GSTINNumber:            gst,
		GSTINStatus:            "ACTIVE",
		GSTInsertionDate:       "2026-04-22 09:10:00",
		Latitude:               "28.5355",
		Location:               "New Delhi",
		Longitude:              "77.3910",
		Pincode:                "110020",
		RegistrationDate:       "2019-07-01 00:00:00",
		StateName:              "Delhi",
		TaxpayerType:           "Regular",
		TradeName:              "Sharma Tools",
		AnnualTurnoverSlab:     "5 - 25 Cr",
	}
}
