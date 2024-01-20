package user

import (
	"encoding/json"
	"net/http"
)

type Controller interface{
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{service}
}

func (controller *controller) Register(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(User{
		Id: "u000",
		Username: "doe",
		Email: "Doe@gmail.com",
		Password: "eyxxx",
	})
}

func (controller *controller) Login(w http.ResponseWriter, r *http.Request) {}

func (controller *controller) Logout(w http.ResponseWriter, r *http.Request) {}
