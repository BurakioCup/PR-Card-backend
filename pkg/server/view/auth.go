package view

type SignInResponse struct {
	LoginId string `json:"loginID"`
}

func ReturnSignInResponse(loginID string) SignInResponse {
	body := SignInResponse{
		LoginId: loginID,
	}
	return body
}

type SignUpResponse struct {
	Token   string `json:"token"`
	LoginId string `json:"loginID"`
}

func ReturnSignUpResponse(token string, uuid string) SignUpResponse {
	body := SignUpResponse{
		Token:   token,
		LoginId: uuid,
	}
	return body
}
