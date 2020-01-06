package routers

import (
	"net/http"

	"github.com/NoSkillGirl/auth-service/controllers"
)

// AuthRoutes - All User related Routes.
func AuthRoutes() {
	http.HandleFunc("/users/login", controllers.UserLogin)
	http.HandleFunc("/users/register", controllers.RegisterUser)
}
