package main

import (
	"fire_heart/models/db"
	"fire_heart/routers"
)

func main() {
	db.Connection()
	router := routers.InitRouter()
	_ = router.Run()
}
