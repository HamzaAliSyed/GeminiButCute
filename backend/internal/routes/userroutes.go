package routes

import (
	"backend/internal/controller"
	"net/http"
)

func RegisterUserRoutes() {
	http.HandleFunc("/v1/user/register", controller.RegisterUserHandle)
}
