package dto

//アカウント認証リクエスト
type SignInRequest struct {
	LoginID string `json:"loginID"`
}