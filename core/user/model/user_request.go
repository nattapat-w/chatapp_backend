package model

type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"displayname"`
}

type Response struct {
	Data         interface{} `json:"data"`
	Status       bool        `json:"status"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
}

type RegisterDataResponse struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayName"`
}
