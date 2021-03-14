package dto

//Request struct

//アカウント認証リクエスト
type SignInRequest struct {
	LoginID string `json:"loginID"`
}

//アカウント作成リクエスト
type SignUpRequest struct {
	ID   string `json:"userId"`
	Pass string `json:"password"`
}
