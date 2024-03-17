package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Response struct {
	Data         interface{} `json:"data"`
	Status       bool        `json:"status"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
}
type UserLoginResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
