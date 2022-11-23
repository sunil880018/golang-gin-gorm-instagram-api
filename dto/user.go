package dto

type UserDTO struct {
	UserId string `json:"userid" binding:"required"`
	Name   string `json:"username" binding:"required"`
	Email  string `json:"email" binding:"required"`
	DOB    string `json:"dob" binding:"required"`
}

// {
// "userid":"123",
// "username":"sunil",
// "email":"aba@gmail.com",
// "dob":"01-01-2001"
// }
