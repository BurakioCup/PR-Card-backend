package dto

//Request struct

//アカウント作成リクエスト
type SignUpRequest struct {
	ID string `json:"userId"`
	Pass string `json:"password"`
}
