package user

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Controller interface{
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service Service
}

type ResponseMetadata struct {
	StatusCode int `json:"statusCode"`
	Title string `json:"title"`
}

type ResponseErrorMessage struct {
	Metadata ResponseMetadata
	Message string `json:"message"`
}

func NewController(service Service) Controller {
	return &controller{service}
}

func (controller *controller) Register(w http.ResponseWriter, r *http.Request) {
		var userValue User
		var err error
		
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userValue)

	if err != nil {
		http.Error(w, "Error decoding data", http.StatusBadRequest)
	}

	userValue = User{
		Id: uuid.NewString(),
		Username: userValue.Username,
		Email: userValue.Email,
		Password: userValue.Password,
	}

	userValue, err = controller.service.Register(r.Context(), userValue)

	if err != nil {
		json.NewEncoder(w).Encode(ResponseErrorMessage{
			Metadata: ResponseMetadata{
				StatusCode: 404,
				Title: "Faield",
			},
			Message: err.Error(),
		})
		return 
	}

	json.NewEncoder(w).Encode(userValue)
}

func (controller *controller) Login(w http.ResponseWriter, r *http.Request) {}

func (controller *controller) Logout(w http.ResponseWriter, r *http.Request) {}
