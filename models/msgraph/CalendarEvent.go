package msgraph

type Body struct {
	ContentType string `json:"contentType"`
	Content string `json:"content"`
}

type Start struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type Location struct {
	DisplayName string `json:"displayName"`
}

type End struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Attendee struct {
	EmailAddress EmailAddress
	Type string `json:"type"`
}

type CalendarEvent struct {
	Subject string `json:"subject"`
	Body Body
	Start Start
	End End
	Location Location
	//Attendee []Attendee
	IsOnlineMeeting bool `json:"isOnlineMeeting"`
	OnlineMeetingProvider string `json:"onlineMeetingProvider"`
}
