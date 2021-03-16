package view

import "PR-Card_backend/pkg/server/model/dto"

type GitStatusResponse struct{
	GitImage string `json:"gitImage"`
}

func ReturnGitStatusResponse(res dto.GitStatus)GitStatusResponse{
	return GitStatusResponse{
		GitImage: res.GitImage,
	}
}
