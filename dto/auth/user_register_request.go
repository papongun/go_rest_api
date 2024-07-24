package auth

type UserRegisterRequest struct {
	Username    string `json:"username" validate:"required,max=32"`
	DisplayName string `json:"displayName" validate:"required,max=32"`
	Password    string `json:"password" validate:"required,min=8"`
}
