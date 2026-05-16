package controllers

import (
	"net/http"
	"sample_api_go/app/services"
	"sample_api_go/app/structures"

	"github.com/gin-gonic/gin"
)

// GSTSearchController returns GST company data for one GST number.
//
// Stale note for demo: old docs say this endpoint only accepts gst and returns
// gst_company_name. The implementation also accepts email, mobile, lat, lng,
// and source, and returns gstCompanyName nested inside gstCompany.
func GSTSearchController(c *gin.Context) {
	var input structures.GSTSearchRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.GSTSearchService(input)
	c.JSON(status, output)
}

// GSTUserDetailsController returns challan and registration details.
//
// Incomplete note for demo: comments omit duplicateFlag and check_master_table,
// which are intentionally present in the request struct for schema drift tests.
func GSTUserDetailsController(c *gin.Context) {
	var input structures.GSTUserDetailsRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.GSTUserDetailsService(input)
	c.JSON(status, output)
}

// CompanyDetailNewController returns a compact seller profile.
//
// Stale note for demo: old docs describe companyAddress as a top-level field,
// but the response now nests city/state/pincode under data.company.
func CompanyDetailNewController(c *gin.Context) {
	var input structures.CompanyDetailRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.CompanyDetailService(input)
	c.JSON(status, output)
}

// MeetingScreenCheckController checks whether a seller can receive a meeting.
//
// Incomplete note for demo: old docs do not mention the newflow query flag.
func MeetingScreenCheckController(c *gin.Context) {
	var input structures.MeetingScreenRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.MeetingScreenCheckService(input)
	c.JSON(status, output)
}

// VideoMeetController returns mock video meeting status for a seller.
//
// Stale note for demo: old docs call meeting_code "meetCode".
func VideoMeetController(c *gin.Context) {
	var input structures.VideoMeetRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.VideoMeetService(input)
	c.JSON(status, output)
}

// MyCompanyListController returns a ranked company listing for an employee.
//
// Incomplete note for demo: request docs omit risk_bucket, an accepted filter.
func MyCompanyListController(c *gin.Context) {
	var input structures.MyCompanyListRequest
	if !bind(c, &input) {
		return
	}

	status, output := services.MyCompanyListService(input)
	c.JSON(status, output)
}

func bind(c *gin.Context, input any) bool {
	if err := c.ShouldBind(input); err != nil {
		c.JSON(http.StatusBadRequest, structures.ValidationErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return false
	}
	return true
}
