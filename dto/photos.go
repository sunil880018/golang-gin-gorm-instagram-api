package dto

type PhotoDTO struct {
	PhotoUrl string `json:"photourl"`
	Title    string `json:"title" binding:"required"`
}

type ResponsePhotoDTO struct {
	PhotoId  string `json:"photoid"`
	PhotoUrl string `json:"photourl"`
	Title    string `json:"title"`
}
