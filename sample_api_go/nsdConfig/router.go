package nsdConfig

import (
	"sample_api_go/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/go/api/nsd")
	{
		v1.GET("/v1/fsa/GstSearch", controllers.GSTSearchController)
		v1.POST("/v1/fsa/GstSearch", controllers.GSTSearchController)

		v1.GET("/v1/GstDetails/gstuserdetails", controllers.GSTUserDetailsController)
		v1.POST("/v1/GstDetails/gstuserdetails", controllers.GSTUserDetailsController)

		v1.GET("/v1/userlist/companydetailnew", controllers.CompanyDetailNewController)
		v1.GET("/v1/NSDDashboard/MeetingScreenCheck", controllers.MeetingScreenCheckController)
		v1.GET("/v1/tele/getVideoMeetData", controllers.VideoMeetController)
		v1.GET("/v1/userlisting/mycompanylist", controllers.MyCompanyListController)
	}
}
