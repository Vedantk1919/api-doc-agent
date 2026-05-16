package models

import (
	"sample_api_go/app/structures"
	"time"
)

func GetCompanyDetail(input structures.CompanyDetailRequest) structures.CompanyDetailData {
	return structures.CompanyDetailData{
		Company: structures.CompanyProfile{
			GLID:           input.GLID,
			CompanyName:    "Sharma Precision Tools Pvt Ltd",
			City:           "New Delhi",
			State:          "Delhi",
			Pincode:        "110020",
			CustType:       "Paid",
			EnterpriseType: "Manufacturer",
			WebsiteURL:     "https://example.com/sharma-tools",
			PrimaryGST:     "07AAECS1234F1Z5",
		},
		Contacts: []structures.CompanyContact{
			{Type: "mobile", Value: "9810012345", Masked: "98100*****", Verified: true, IsPrimary: true},
			{Type: "email", Value: "sales@sharmatools.example", Masked: "sa***@sharmatools.example", Verified: true, IsPrimary: true},
		},
		Verification: structures.VerificationState{
			GSTVerified:    true,
			EmailVerified:  true,
			MobileVerified: true,
			CINStatus:      "VERIFIED",
			LastCheckedAt:  "2026-05-16T10:20:00+05:30",
		},
		Assignment: structures.EmployeeAssignment{
			Empid:    "98123",
			Name:     "Amit Verma",
			Vertical: "Industrial Supplies",
			Role:     "NSD Sales",
		},
		ServiceSummary: structures.ServiceSummary{
			PaidServiceCount: 3,
			LatestService:    "MDC Premium",
			WorkOrderCount:   2,
		},
		RelatedGSTNumbers: []string{"07AAECS1234F1Z5", "09AAECS1234F1Z3"},
	}
}

func GetCompanyList(input structures.MyCompanyListRequest) []structures.CompanyListItem {
	now := time.Now().Add(2 * time.Hour).Unix()
	list := []structures.CompanyListItem{
		{
			PageNo:                   0,
			CompanyName:              "Sharma Precision Tools Pvt Ltd",
			CompanyID:                "STS10091",
			CompanyGLID:              "1245789",
			CompanyCity:              "New Delhi",
			CompanyState:             "Delhi",
			CompanyLastMeetDate:      "14-May-26",
			CompanyFCPStatus:         "Paid",
			CompanyGSTVerifyStatus:   "Verified",
			CompanyLatitude:          "28.5355",
			CompanyLongitude:         "77.3910",
			CompanyPin:               "110020",
			ContactName:              "Rohit Sharma",
			WebsiteURL:               "https://example.com/sharma-tools",
			ReminderTimestamp:        now,
			CompanyReminderDays:      "+1D",
			CompanyLastMeetDays:      "2D",
			CompanyLastCallDays:      "1D",
			HLName:                   "Payment Attempt",
			HLFilter:                 1,
			PriorityKeyStr:           "1-0-0-1-1-0-1-1",
			Rank:                     12,
			HighTurnOver:             1,
			EffectiveConnectFlag:     1,
			VerifiedVideoConnectFlag: 1,
			LastMetDisposition:       "Interested",
			LastMetSalesType:         "Follow Up",
			NewFreshFilter:           "1",
			AnnualTurnoverDisplay:    "5 - 25 Cr",
			CompanyMobileList: []structures.MobileDetails{
				{Display: "98100***** - Primary", Value: "9810012345", EncryptValue: "demo-encrypted-9810012345"},
			},
			DeprecatedCompanyBucketName: "Prime",
		},
		{
			PageNo:                   0,
			CompanyName:              "Khan Packaging Works",
			CompanyID:                "STS10092",
			CompanyGLID:              "1245790",
			CompanyCity:              "Noida",
			CompanyState:             "Uttar Pradesh",
			CompanyLastMeetDate:      "",
			CompanyFCPStatus:         "Free",
			CompanyGSTVerifyStatus:   "No GST",
			CompanyLatitude:          "28.5821",
			CompanyLongitude:         "77.3267",
			CompanyPin:               "201301",
			ContactName:              "Sameer Khan",
			ReminderTimestamp:        now + 86400,
			CompanyReminderDays:      "+2D",
			CompanyLastMeetDays:      "",
			CompanyLastCallDays:      "12D",
			HLName:                   "Schedule Demo",
			HLFilter:                 1,
			PriorityKeyStr:           "0-1-1-0-1-0-0-0",
			Rank:                     44,
			HighTurnOver:             0,
			EffectiveConnectFlag:     0,
			VerifiedVideoConnectFlag: 0,
			LastMetDisposition:       "Meeting Not Done",
			LastMetSalesType:         "Fresh",
			NewFreshFilter:           "0",
			AnnualTurnoverDisplay:    "0 - 40 L",
			CompanyMobileList: []structures.MobileDetails{
				{Display: "98765***** - Primary", Value: "9876500000", EncryptValue: "demo-encrypted-9876500000"},
			},
			DeprecatedCompanyBucketName: "Rest HL",
		},
	}

	if input.RiskBucket == "prime" {
		return list[:1]
	}
	return list
}
