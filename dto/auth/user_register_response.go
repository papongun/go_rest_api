package auth

type UserRegisterResponse struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}
