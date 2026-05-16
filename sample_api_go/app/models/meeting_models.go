package models

import (
	"sample_api_go/app/structures"
	"time"
)

func GetMeetingScreen(input structures.MeetingScreenRequest) structures.MeetingScreenData {
	allowed := input.Flag != "followup-limit"
	reason := ""
	count := 2
	message := "Meeting Allowed"
	if !allowed {
		reason = "MAX_28_DAY_FOLLOWUP"
		count = 4
		message = "Maximum limit reached with the seller in the last 28 Days"
	}

	return structures.MeetingScreenData{
		EmpID:               input.EmpID,
		GLID:                input.GLID,
		STSID:               "STS10091",
		ServerDate:          time.Now().Format("2006,01,02"),
		IsMeetingAllowed:    allowed,
		Count:               count,
		BounceCountStatus:   "200",
		BounceCount:         "1",
		BounceCountMessage:  "1 Bounce in 12 Month",
		CompanyCount:        "1",
		CompanyCountMessage: "Company is assigned to " + input.EmpID,
		CheckTBRO: structures.TBROWindow{
			WholeDate: "2026-05-17 11:30:00",
			TBDate:    "17-May-26",
			TBHour:    "11",
			TBMinute:  "30",
			TBROID:    "98123",
		},
		CheckTBROMessage:      "TBRO List with Date and Time",
		MeetingCount:          structures.MeetingCountStatus{Signboard: "1", Location: "1", EnableGalleryOption: "1", MeetingCount: "0"},
		MeetingCountMessage:   message,
		IsOtpValidationSkip:   1,
		IsMeetingDone:         0,
		FreshMeetCount:        "1",
		RestrictionReasonCode: reason,
	}
}

func GetVideoMeet(input structures.VideoMeetRequest) structures.VideoMeetData {
	meetingID := input.MeetingID
	if meetingID == "" {
		meetingID = "VM-20260516-001"
	}

	return structures.VideoMeetData{
		GLID:        input.GLID,
		MeetingID:   meetingID,
		MeetingCode: "NSD-8421",
		JoinURL:     "https://meet.example.com/nsd/NSD-8421",
		TodayVideoCall: structures.VideoCallStatus{
			Status:      "scheduled",
			ScheduledAt: "2026-05-16T15:30:00+05:30",
			DurationMin: 30,
		},
		LastVideoCall: structures.VideoCallStatus{
			Status:      "completed",
			ScheduledAt: "2026-05-12T11:00:00+05:30",
			DurationMin: 24,
		},
		ManagerContacts: []structures.ManagerContact{
			{Empid: "98123", Name: "Amit Verma", Mobile: "9810012345", Email: "amit.verma@example.com"},
			{Empid: "98124", Name: "Neha Rao", Mobile: "9810099999", Email: "neha.rao@example.com"},
		},
		CompanyTimeZone:  "Asia/Kolkata",
		CanScheduleAgain: true,
	}
}
