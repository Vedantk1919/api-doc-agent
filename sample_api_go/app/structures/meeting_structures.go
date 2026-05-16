package structures

type MeetingScreenRequest struct {
	EmpID      string `json:"empid" form:"empid" binding:"required"`
	AuthToken  string `json:"AK" form:"AK" binding:"required"`
	VerticalID string `json:"vertical_id" form:"vertical_id"`
	GLID       string `json:"glid" form:"glid" binding:"required"`
	Flag       string `json:"flag" form:"flag"`
	ModID      string `json:"modid" form:"modid"`
	NewFlow    bool   `json:"newflow" form:"newflow"`
}

type MeetingScreenResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    MeetingScreenData `json:"data"`
}

type MeetingScreenData struct {
	EmpID                 string             `json:"empid"`
	GLID                  string             `json:"glid"`
	STSID                 string             `json:"stsid"`
	ServerDate            string             `json:"serverdate"`
	IsMeetingAllowed      bool               `json:"isMeetingAllowed"`
	Count                 int                `json:"count"`
	BounceCountStatus     string             `json:"bounce_count_status"`
	BounceCount           string             `json:"bounce_count"`
	BounceCountMessage    string             `json:"bounce_count_message"`
	CompanyCount          string             `json:"companycount"`
	CompanyCountMessage   string             `json:"companycount_message"`
	CheckTBRO             TBROWindow         `json:"checktbro"`
	CheckTBROMessage      string             `json:"checktbro_message"`
	MeetingCount          MeetingCountStatus `json:"meetingCount"`
	MeetingCountMessage   string             `json:"meetingCount_message"`
	IsOtpValidationSkip   int                `json:"isOtpValidationSkip"`
	IsMeetingDone         int                `json:"ismeetingdone"`
	FreshMeetCount        string             `json:"freshMeetCount,omitempty"`
	RestrictionReasonCode string             `json:"restrictionReasonCode,omitempty"`
}

type TBROWindow struct {
	WholeDate string `json:"WHOLE_DATE"`
	TBDate    string `json:"TB_DATE"`
	TBHour    string `json:"TB_H"`
	TBMinute  string `json:"TB_M"`
	TBROID    string `json:"TBRO_ID"`
}

type MeetingCountStatus struct {
	Signboard           string `json:"signboard"`
	Location            string `json:"location"`
	EnableGalleryOption string `json:"enable_gallery_option"`
	MeetingCount        string `json:"meetingcount"`
}

type VideoMeetRequest struct {
	Empid      string `form:"empid" json:"empid" binding:"required"`
	AK         string `form:"AK" json:"AK" binding:"required"`
	GLID       string `form:"glid" json:"glid" binding:"required"`
	MeetingID  string `form:"meeting_id" json:"meeting_id"`
	Modid      string `form:"modid" json:"modid"`
	ScreenName string `form:"screen_name" json:"screen_name"`
}

type VideoMeetResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    VideoMeetData `json:"data"`
}

type VideoMeetData struct {
	GLID             string           `json:"glid"`
	MeetingID        string           `json:"meeting_id"`
	MeetingCode      string           `json:"meeting_code"`
	JoinURL          string           `json:"join_url"`
	TodayVideoCall   VideoCallStatus  `json:"today_video_call"`
	LastVideoCall    VideoCallStatus  `json:"last_video_call"`
	ManagerContacts  []ManagerContact `json:"manager_contacts"`
	CompanyTimeZone  string           `json:"company_time_zone"`
	CanScheduleAgain bool             `json:"can_schedule_again"`
}

type VideoCallStatus struct {
	Status      string `json:"status"`
	ScheduledAt string `json:"scheduled_at"`
	DurationMin int    `json:"duration_min"`
}

type ManagerContact struct {
	Empid  string `json:"empid"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
