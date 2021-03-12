package view

type SignUpResponse struct{
	Token string `json:"token"`
	LoginId string `json:"loginID"`
}

func ReturnSignUpResponse(token string, uuid string) SignUpResponse {
	body := SignUpResponse{
		Token: token,
		LoginId: uuid,
	}
	return body
}
