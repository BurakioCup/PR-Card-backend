package dto

type MyCard struct{
	UserName string
	FaceImage string
	NickName string
	StatusImage string
	words []Word
	FreeText string
}

type Word struct{
	word string
}
