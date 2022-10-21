package models

type UserToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTokenResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func CreateUserTokenResponse(hasError bool, message string, token string) *UserTokenResponse {

	return &UserTokenResponse{
		Error:   hasError,
		Message: message,
		Token:   token,
	}

}
