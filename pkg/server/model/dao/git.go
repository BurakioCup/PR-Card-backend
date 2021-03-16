package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"encoding/base64"
	"os"

	"github.com/canhlinh/svg2png"
	"github.com/sirupsen/logrus"
)

type gitStatus struct{

}

func MakeGitStatusClient()gitStatus{
	return gitStatus{}
}

func (info  *gitStatus)Request(gitName string)(dto.GitStatus,error){
	var gitStatus dto.GitStatus
	var err error
	url :="https://github-readme-stats.vercel.app/api/top-langs/?username="+gitName
	chrome := svg2png.NewChrome().SetHeight(250).SetWith(280)
	filepath := "git-status.png"
	if err := chrome.Screenshoot(url, filepath); err != nil {
		logrus.Println(err)
	}

	file, _ := os.Open("git-status.png")
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	gitStatus.GitImage = base64.StdEncoding.EncodeToString(data)
	//exec.Command("open", filepath).Run()
	return gitStatus,err
}
