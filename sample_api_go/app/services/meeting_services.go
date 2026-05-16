package services

import (
	"net/http"
	"sample_api_go/app/models"
	"sample_api_go/app/structures"
)

func MeetingScreenCheckService(input structures.MeetingScreenRequest) (int, structures.MeetingScreenResponse) {
	if input.GLID == "" {
		return http.StatusBadRequest, structures.MeetingScreenResponse{
			Status:  http.StatusBadRequest,
			Message: "glid is required",
		}
	}

	data := models.GetMeetingScreen(input)
	return http.StatusOK, structures.MeetingScreenResponse{
		Status:  http.StatusOK,
		Message: data.MeetingCountMessage,
		Data:    data,
	}
}

func VideoMeetService(input structures.VideoMeetRequest) (int, structures.VideoMeetResponse) {
	if input.GLID == "" {
		return http.StatusBadRequest, structures.VideoMeetResponse{
			Status:  http.StatusBadRequest,
			Message: "glid is required",
		}
	}

	return http.StatusOK, structures.VideoMeetResponse{
		Status:  http.StatusOK,
		Message: "Video meet data",
		Data:    models.GetVideoMeet(input),
	}
}
