package main

import (
	"net/http"

	"github.com/NoSkillGirl/auth-service/routers"
)

func main() {
	// AuthRoutes Initilization
	routers.AuthRoutes()
	http.ListenAndServe(":8083", nil)
}
