package services

import (
	"net/http"
	"sample_api_go/app/models"
	"sample_api_go/app/structures"
)

func CompanyDetailService(input structures.CompanyDetailRequest) (int, structures.CompanyDetailResponse) {
	if input.GLID == "" {
		return http.StatusBadRequest, structures.CompanyDetailResponse{
			Status:  http.StatusBadRequest,
			Message: "glid is required",
		}
	}

	return http.StatusOK, structures.CompanyDetailResponse{
		Status:  http.StatusOK,
		Message: "Company Detail",
		Data:    models.GetCompanyDetail(input),
	}
}

func MyCompanyListService(input structures.MyCompanyListRequest) (int, structures.MyCompanyListResponse) {
	list := models.GetCompanyList(input)
	return http.StatusOK, structures.MyCompanyListResponse{
		Status:  http.StatusOK,
		Message: "My Company Listing Data",
		Data: structures.MyCompanyListData{
			TotalCompanyCount: len(list),
			AllFilters:        "all,untouched,hotlead,reminder,prime,twox,active,gst,unmet,nearme,catalog",
			HLMessage:         "",
			Message:           "My Company Listing Data",
			MyCompanyData:     list,
		},
	}
}
