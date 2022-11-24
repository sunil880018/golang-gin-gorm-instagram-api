package dto

type PhotoDTO struct {
	PhotoUrl string `json:"photourl"`
	Title    string `json:"title" binding:"required"`
}
