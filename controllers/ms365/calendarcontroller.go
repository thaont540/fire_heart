package ms365

import (
	"fire_heart/utils"
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
		c.HTML(200, "calendars", gin.H{
			"response": "",
		})
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
		c.HTML(200, "calendar", gin.H{
			"response": "",
		})
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
