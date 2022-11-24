package dto

type PhotoDTO struct {
	PhotoId  string `json:"photoid" binding:"required"`
	PhotoUrl string `json:"photourl" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

type ResponsePhotoDTO struct {
	PhotoId  string `json:"photoid"`
	PhotoUrl string `json:"photourl"`
	Title    string `json:"title"`
}
