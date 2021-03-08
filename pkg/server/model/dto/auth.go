package dto

//Request struct

//アカウント作成リクエスト
type SignUpRequest struct {
	Id string `json:"userId"`
	Pass string `json:"password"`
}
