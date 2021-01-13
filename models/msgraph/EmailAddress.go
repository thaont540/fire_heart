package msgraph

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func (e EmailAddress) String() string {
	return e.Name + "<" + e.Address + ">"
}
