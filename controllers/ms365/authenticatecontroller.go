package ms365

import (
	"fire_heart/utils"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AuthenticateController struct {
}

func (authenticateController *AuthenticateController)Index(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"client_id": utils.Env("MS365_CLIENT_ID"),
		"call_back": utils.Env("MS365_CALL_BACK"),
	})
}

func (authenticateController *AuthenticateController)CallBack(c *gin.Context) {
	query := c.Request.URL.Query()
	endpoint := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	data := url.Values{}
	data.Set("client_id", utils.Env("MS365_CLIENT_ID"))
	data.Set("scope", "https://graph.microsoft.com/Calendars.Read")
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", utils.Env("MS365_CALL_BACK"))
	data.Set("client_secret", utils.Env("MS365_CLIENT_SECRET"))
	data.Set("code", strings.TrimSpace(query.Get("code")))

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	utils.GlobalCache.Set("authenticated", string(body), cache.DefaultExpiration)

	c.Redirect(http.StatusFound, "/calendars")
}