package main

import (
	"encoding/json"
	"fire_heart/models/db"
	"fire_heart/models/msgraph"
	"fire_heart/utils"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	db.Connection()
	//router := routers.InitRouter()
	cache1 := cache.New(50*time.Minute, 100*time.Minute)
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		authenticatedCache, found := cache1.Get("authenticated")
		if found {
			authenticated := &msgraph.Token{}
			err1 := json.Unmarshal([]byte(authenticatedCache.(string)), authenticated)
			fmt.Println(err1)
			fmt.Println(authenticated.AccessToken)
		}
		c.HTML(200, "index", gin.H{
			"client_id": utils.Env("MS365_CLIENT_ID"),
			"call_back": utils.Env("MS365_CALL_BACK"),
		})
	})

	router.GET("/calendars", func(c *gin.Context) {
		authenticatedCache, found := cache1.Get("authenticated")
		if found {
			authenticated := &msgraph.Token{}
			_ = json.Unmarshal([]byte(authenticatedCache.(string)), authenticated)
			fmt.Println(authenticated.AccessToken)
			endpoint := "https://graph.microsoft.com/v1.0/me/calendars"

			client := &http.Client{}
			r, err := http.NewRequest("GET", endpoint, nil) // URL-encoded payload
			if err != nil {
				log.Fatal(err)
			}
			r.Header.Add("Content-Type", "application/json")
			r.Header.Add("Authorization", "Bearer " + authenticated.AccessToken)

			res, err := client.Do(r)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)

			fmt.Println(string(body))

			c.HTML(200, "calendars", gin.H{
				"response": string(body),
			})
		}

		c.HTML(200, "calendars", gin.H{
			"response": "",
		})
	})
	router.GET("/callback", func(c *gin.Context) {
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
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
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

		cache1.Set("authenticated", string(body), cache.DefaultExpiration)


		c.Redirect(http.StatusFound, "/calendars")
	})

	_ = router.Run()
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/index.html")
	r.AddFromFiles("callback", "templates/callback.html")
	r.AddFromFiles("calendars", "templates/calendars.html")
	return r
}
