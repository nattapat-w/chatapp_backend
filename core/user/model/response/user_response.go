package response

type Response struct {
	Data         interface{} `json:"data"`
	Status       bool        `json:"status"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
}

type RegisterDataResponse struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayName"`
}
