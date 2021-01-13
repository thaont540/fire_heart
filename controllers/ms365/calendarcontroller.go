package ms365

import (
	"fire_heart/models/msgraph"
	"fire_heart/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type CalendarController struct {}

func (calendarController *CalendarController)Index(c *gin.Context) {
	utils.GetToken()
	if utils.Ms365Token != "" {
		endpoint := "https://graph.microsoft.com/v1.0/me/calendars"
		res := utils.AuthenticatedHttpGet(endpoint)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		c.HTML(200, "calendars", gin.H{
			"response": string(body),
		})
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func (calendarController *CalendarController)Show(c *gin.Context) {
	calendarId := c.Param("calendarId")

	utils.GetToken()
	if utils.Ms365Token != "" {
		endpoint := "https://graph.microsoft.com/v1.0/me/calendars/" + calendarId + "/events"
		res := utils.AuthenticatedHttpGet(endpoint)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		endpoint2 := "https://graph.microsoft.com/v1.0/me/calendars/" + calendarId
		res2 := utils.AuthenticatedHttpGet(endpoint2)
		defer res2.Body.Close()
		body2, _ := ioutil.ReadAll(res2.Body)
		c.HTML(200, "calendar", gin.H{
			"events": string(body),
			"calendar": string(body2),
		})
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func (calendarController *CalendarController)Store(c *gin.Context) {
	type CalendarCreateInput struct {
		Name string `form:"name"`
	}

	var input CalendarCreateInput
	if err := c.Bind(&input); err != nil {
		c.Redirect(http.StatusFound, "/calendars")
	}

	if input.Name != "" {
		data := map[string]string{
			"name": input.Name,
		}

		endpoint := "https://graph.microsoft.com/v1.0/me/calendars"
		res := utils.AuthenticatedHttpPost(endpoint, data)

		defer res.Body.Close()
	}

	c.Redirect(http.StatusFound, "/calendars")
}

func (calendarController *CalendarController)CreateEvent(c *gin.Context) {
	type CalendarEventCreateInput struct {
		Subject string `form:"subject"`
		Content string `form:"content"`
		Location string `form:"location"`
		AttendeeName string `form:"attendee_name"`
		AttendeeEmail string `form:"attendee_email"`
		Start string `form:"start"`
		End string `form:"end"`
	}

	calendarId := c.Param("calendarId")

	var input CalendarEventCreateInput

	if err := c.Bind(&input); err != nil {
		c.Redirect(http.StatusFound, "/calendars/" + calendarId)
	}

	eventBody := msgraph.Body{}
	eventBody.ContentType = "HTML"
	eventBody.Content = input.Content

	eventStart := msgraph.Start{}
	eventStart.TimeZone = "Pacific Standard Time"
	eventStart.DateTime = input.Start

	eventEnd := msgraph.End{}
	eventEnd.TimeZone = "Pacific Standard Time"
	eventEnd.DateTime = input.End

	eventLocation := msgraph.Location{}
	eventLocation.DisplayName = input.Location

	eventAttendee := msgraph.Attendee{}
	emailAddress := msgraph.EmailAddress{}
	emailAddress.Name = input.AttendeeName
	emailAddress.Address = input.AttendeeEmail
	eventAttendee.EmailAddress = emailAddress
	eventAttendee.Type = "required"

	event := msgraph.CalendarEvent{}
	event.Subject = input.Subject
	event.Body = eventBody
	event.Start = eventStart
	event.End = eventEnd
	event.Location = eventLocation
	//event.Attendee = []msgraph.Attendee{eventAttendee}
	event.IsOnlineMeeting = true
	event.OnlineMeetingProvider = "teamsForBusiness"

	endpoint := "https://graph.microsoft.com/v1.0/me/calendars/" + calendarId + "/events"
	res := utils.AuthenticatedHttpPost1(endpoint, event)

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	defer res.Body.Close()

	c.Redirect(http.StatusFound, "/calendars/" + calendarId)
}