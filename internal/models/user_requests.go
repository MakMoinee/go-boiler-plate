package models

type UserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
}
