package services

import (
	"net/http"
	"regexp"
	"strings"

	"sample_api_go/app/models"
	"sample_api_go/app/structures"
)

var gstPattern = regexp.MustCompile(`^[0-9]{2}[A-Z]{5}[0-9]{4}[A-Z][1-9A-Z]Z[0-9A-Z]$`)

func GSTSearchService(input structures.GSTSearchRequest) (int, structures.GSTSearchResponse) {
	input.GST = strings.ToUpper(strings.TrimSpace(input.GST))

	if input.GST == "" && input.Email == "" && input.Mobile == "" {
		return http.StatusBadRequest, structures.GSTSearchResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
		}
	}
	if input.GST != "" && !gstPattern.MatchString(input.GST) {
		return http.StatusBadRequest, structures.GSTSearchResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid GST format",
		}
	}

	result := models.SearchGST(input)
	return http.StatusOK, structures.GSTSearchResponse{
		Status:  http.StatusOK,
		Message: "GST Search Result",
		Data:    result,
	}
}

func GSTUserDetailsService(input structures.GSTUserDetailsRequest) (int, structures.GSTUserDetailsResponse) {
	input.GST = strings.ToUpper(strings.TrimSpace(input.GST))
	if input.GST == "" {
		return http.StatusBadRequest, structures.GSTUserDetailsResponse{
			Status:  http.StatusBadRequest,
			Message: "GST is required",
		}
	}
	if !gstPattern.MatchString(input.GST) {
		return http.StatusBadRequest, structures.GSTUserDetailsResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid GST format",
		}
	}

	registration := models.GetGSTRegistration(input.GST)
	return http.StatusOK, structures.GSTUserDetailsResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: structures.GSTUserDetailsData{
			Message:     "Data Found",
			Result:      registration,
			IsDuplicate: input.DuplicateFlag == "1" && input.GST == "07AAECS1234F1Z5",
			RefreshedBy: refreshMode(input),
		},
	}
}

func refreshMode(input structures.GSTUserDetailsRequest) string {
	if input.RealTime == "1" || input.FcpRealTime == "1" {
		return "realtime-masterindia"
	}
	if input.CheckMaster == "1" {
		return "master-table"
	}
	return "cache"
}
