package structures

type ValidationErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type GSTSearchRequest struct {
	GST    string `form:"gst" json:"gst"`
	Email  string `form:"email" json:"email"`
	Mobile string `form:"mobile" json:"mobile"`
	Lat    string `form:"lat" json:"lat"`
	Lng    string `form:"lng" json:"lng"`
	Empid  string `form:"empid" json:"empid" binding:"required"`
	AK     string `form:"AK" json:"AK" binding:"required"`
	Modid  string `form:"MODID" json:"MODID"`
	Source string `form:"source" json:"source"`
}

type GSTSearchResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    GSTSearchData `json:"data"`
}

type GSTSearchData struct {
	Message           string            `json:"message"`
	GSTNumber         string            `json:"gstNumber"`
	GSTCompany        *GSTCompany       `json:"gstCompany"`
	ExistingCompanies []ExistingCompany `json:"existingCompanies"`
	ResultType        int               `json:"resultType"`
	RequestSource     string            `json:"requestSource,omitempty"`
	// Deprecated: use gstCompany.gstCompanyName.
	LegacyCompanyName string `json:"legacyCompanyName,omitempty"`
}

type GSTCompany struct {
	GSTCompanyName      string           `json:"gstCompanyName"`
	GSTCompanyAddress   string           `json:"gstCompanyAddress"`
	GSTNumber           string           `json:"gstNumber"`
	GSTPinCode          string           `json:"gstPinCode"`
	EnterpriseType      string           `json:"enterpriseType"`
	BusinessNatureType  string           `json:"businessNatureType"`
	GSTCompanyLatitude  string           `json:"gstCompanyLatitude"`
	GSTCompanyLongitude string           `json:"gstCompanyLongitude"`
	DistanceFromMyLoc   float64          `json:"distanceFromMyLocation"`
	GSTData             []GSTDisplayData `json:"gstData"`
}

type GSTDisplayData struct {
	DataKey string `json:"dataKey"`
	Value   string `json:"value"`
}

type ExistingCompany struct {
	GLID                 string  `json:"glid"`
	GSTCompanyName       string  `json:"gstCompanyName"`
	SalesAssignedTo      string  `json:"salesAssignedTo"`
	GSTCompanyLatitude   string  `json:"gstCompanyLatitude"`
	GSTCompanyLongitude  string  `json:"gstCompanyLongitude"`
	DistanceFromMyLoc    float64 `json:"distanceFromMyLocation"`
	CompanyCustType      string  `json:"companyCustType"`
	CurrentMeetingStatus string  `json:"currentMeetingStatus"`
}

type GSTUserDetailsRequest struct {
	Empid         string `form:"empid" json:"empid" binding:"required"`
	AuthToken     string `form:"AK" json:"AK" binding:"required"`
	Modid         string `form:"modid" json:"modid"`
	Screen        string `form:"screen_name" json:"screen_name"`
	Response      string `form:"response" json:"response"`
	Mode          string `form:"mode" json:"mode"`
	GLID          string `form:"glid" json:"glid"`
	RealTime      string `form:"realtime" json:"realtime"`
	FcpRealTime   string `form:"fcprealtime" json:"fcprealtime"`
	GST           string `form:"gst" json:"gst" binding:"required"`
	Display       string `form:"display" json:"display"`
	DuplicateFlag string `form:"duplicateFlag" json:"duplicateFlag"`
	CheckMaster   string `form:"check_master_table" json:"check_master_table"`
	Platform      string `form:"platform" json:"platform"`
	OS            string `form:"os" json:"os"`
}

type GSTUserDetailsResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    GSTUserDetailsData `json:"data"`
}

type GSTUserDetailsData struct {
	Message     string                 `json:"message"`
	Result      GSTRegistrationDetails `json:"result"`
	IsDuplicate bool                   `json:"isDuplicate"`
	RefreshedBy string                 `json:"refreshedBy"`
}

type GSTRegistrationDetails struct {
	BuildingName           string `json:"building_name"`
	BusinessActivityNature string `json:"business_activity_nature"`
	BusinessAddressAdd     string `json:"bussiness_address_add"`
	BusinessConstitution   string `json:"business_constitution"`
	BusinessName           string `json:"business_name"`
	DateOfFiling           string `json:"date_of_filing"`
	FilingLastUpdationDate string `json:"filing_last_updation_date"`
	GSTINNumber            string `json:"gstin_number"`
	GSTINStatus            string `json:"gstin_status"`
	GSTInsertionDate       string `json:"gst_insertion_date"`
	Latitude               string `json:"lattitude"`
	Location               string `json:"location"`
	Longitude              string `json:"longitude"`
	Pincode                string `json:"pincode"`
	RegistrationDate       string `json:"registration_date"`
	StateName              string `json:"state_name"`
	TaxpayerType           string `json:"taxpayer_type"`
	TradeName              string `json:"trade_name"`
	AnnualTurnoverSlab     string `json:"annual_turnover_slab"`
}
