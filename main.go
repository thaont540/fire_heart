package main

import (
	"fire_heart/controllers/ms365"
	"fire_heart/models/db"
	"fire_heart/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	_ "time"
)

func main() {
	db.Connection()
	utils.InitCache()
	//router := routers.InitRouter()
	router := gin.Default()
	router.HTMLRender = createMyRender()

	authenticateController := new(ms365.AuthenticateController)
	calendarController := new (ms365.CalendarController)

	router.GET("/", authenticateController.Index)
	router.GET("/calendars", calendarController.Index)
	router.POST("/calendars", calendarController.Store)
	router.GET("/calendars/:calendarId", calendarController.Show)
	router.GET("/callback", authenticateController.CallBack)

	_ = router.Run()
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/index.html")
	r.AddFromFiles("callback", "templates/callback.html")
	r.AddFromFiles("calendars", "templates/calendars.html")
	r.AddFromFiles("calendar", "templates/calendar.html")
	return r
}
