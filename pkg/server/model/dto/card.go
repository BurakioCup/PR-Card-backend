package dto

type MyCard struct{
	UserName string
	FaceImage string
	NickName string
	StatusImage string
	Words Words
	FreeText string
}

type Words struct{
	Word [4]string
}