package structures

type CompanyDetailRequest struct {
	AK              string `form:"AK" json:"AK" binding:"required"`
	Empid           string `form:"empid" json:"empid" binding:"required"`
	Modid           string `form:"MODID" json:"MODID"`
	ScreenName      string `form:"screen_name" json:"screen_name"`
	OS              string `form:"os" json:"os"`
	Platform        string `form:"platform" json:"platform"`
	ApiEmpid        string `form:"apiempid" json:"apiempid"`
	Native          string `form:"native" json:"native"`
	All             string `form:"all" json:"all"`
	Keyword         string `form:"keyword" json:"keyword"`
	GLID            string `form:"glid" json:"glid" binding:"required"`
	IncludeInactive string `form:"includeInactive" json:"includeInactive"`
}

type CompanyDetailResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    CompanyDetailData `json:"data"`
}

type CompanyDetailData struct {
	Company           CompanyProfile     `json:"company"`
	Contacts          []CompanyContact   `json:"contacts"`
	Verification      VerificationState  `json:"verification"`
	Assignment        EmployeeAssignment `json:"assignment"`
	ServiceSummary    ServiceSummary     `json:"serviceSummary"`
	RelatedGSTNumbers []string           `json:"relatedGstNumbers"`
}

type CompanyProfile struct {
	GLID           string `json:"glid"`
	CompanyName    string `json:"companyName"`
	City           string `json:"city"`
	State          string `json:"state"`
	Pincode        string `json:"pincode"`
	CustType       string `json:"custType"`
	EnterpriseType string `json:"enterpriseType"`
	WebsiteURL     string `json:"websiteUrl"`
	PrimaryGST     string `json:"primaryGst"`
}

type CompanyContact struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Masked    string `json:"masked"`
	Verified  bool   `json:"verified"`
	IsPrimary bool   `json:"isPrimary"`
}

type VerificationState struct {
	GSTVerified    bool   `json:"gstVerified"`
	EmailVerified  bool   `json:"emailVerified"`
	MobileVerified bool   `json:"mobileVerified"`
	CINStatus      string `json:"cinStatus"`
	LastCheckedAt  string `json:"lastCheckedAt"`
}

type EmployeeAssignment struct {
	Empid    string `json:"empid"`
	Name     string `json:"name"`
	Vertical string `json:"vertical"`
	Role     string `json:"role"`
}

type ServiceSummary struct {
	PaidServiceCount int    `json:"paidServiceCount"`
	LatestService    string `json:"latestService"`
	WorkOrderCount   int    `json:"workOrderCount"`
}

type MyCompanyListRequest struct {
	Empid         string `form:"empid" json:"empid" binding:"required"`
	AK            string `form:"AK" json:"AK" binding:"required"`
	SelectedEmpid string `form:"selected_empid" json:"selected_empid"`
	Records       string `form:"records" json:"records"`
	BuildVersion  string `form:"build_version" json:"build_version"`
	VersionCode   string `form:"version_code" json:"version_code"`
	Platform      string `form:"platform" json:"platform"`
	VerticalID    string `form:"vertical_id" json:"vertical_id"`
	OS            string `form:"os" json:"os"`
	Display       string `form:"display" json:"display"`
	Modid         string `form:"modid" json:"modid"`
	Flag          string `form:"flag" json:"flag"`
	RiskBucket    string `form:"risk_bucket" json:"risk_bucket"`
}

type MyCompanyListResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    MyCompanyListData `json:"data"`
}

type MyCompanyListData struct {
	TotalCompanyCount int               `json:"totalCompanyCount"`
	AllFilters        string            `json:"allFilters"`
	HLMessage         string            `json:"hlMessage"`
	Message           string            `json:"message"`
	MyCompanyData     []CompanyListItem `json:"myCompanyData"`
}

type CompanyListItem struct {
	PageNo                   int             `json:"pageNo"`
	CompanyName              string          `json:"companyName"`
	CompanyID                string          `json:"companyId"`
	CompanyGLID              string          `json:"companyGlid"`
	CompanyCity              string          `json:"companyCity"`
	CompanyState             string          `json:"companyState"`
	CompanyLastMeetDate      string          `json:"companyLastMeetDate"`
	CompanyFCPStatus         string          `json:"companyFcpStatus"`
	CompanyGSTVerifyStatus   string          `json:"companyGstVerifyStatus"`
	CompanyLatitude          string          `json:"companyGlusrLatitude"`
	CompanyLongitude         string          `json:"companyGlusrLongitude"`
	CompanyPin               string          `json:"companyPin"`
	CompanyMobileList        []MobileDetails `json:"companyMobileList"`
	ContactName              string          `json:"contactName"`
	WebsiteURL               string          `json:"websiteUrl"`
	ReminderTimestamp        int64           `json:"reminderTimestamp"`
	CompanyReminderDays      string          `json:"companyReminderDays"`
	CompanyLastMeetDays      string          `json:"companyLastMeetDays"`
	CompanyLastCallDays      string          `json:"companyLastcallDays"`
	HLName                   string          `json:"hlName"`
	HLFilter                 int             `json:"hlFilter"`
	PriorityKeyStr           string          `json:"priorityKeyStr"`
	Rank                     int             `json:"rank"`
	HighTurnOver             int             `json:"HighTurnOver"`
	EffectiveConnectFlag     int             `json:"EffectiveConnectFlag"`
	VerifiedVideoConnectFlag int             `json:"VerifiedVideoConnectFlag"`
	LastMetDisposition       string          `json:"lastmetDisposition"`
	LastMetSalesType         string          `json:"lastmetSalesType"`
	NewFreshFilter           string          `json:"newFreshFilter"`
	AnnualTurnoverDisplay    string          `json:"annual_turnover_val"`
	// Deprecated: use priorityKeyStr and rank instead.
	DeprecatedCompanyBucketName string `json:"companyBucketName,omitempty"`
}

type MobileDetails struct {
	Display      string `json:"display"`
	Value        string `json:"value"`
	EncryptValue string `json:"encryptValue"`
}
