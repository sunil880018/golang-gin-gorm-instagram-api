package dto

type UserDTO struct {
	UserId string `json:"userid" binding:"required"`
	Name   string `json:"username" binding:"required"`
	Email  string `json:"email" binding:"required"`
	DOB    string `json:"dob" binding:"required"`
}

type ResponseUserDTO struct {
	UserId string `json:"userid"`
	Name   string `json:"username"`
	Email  string `json:"email"`
	DOB    string `json:"dob"`
}

// {
// "userid":"123",
// "username":"sunil",
// "email":"aba@gmail.com",
// "dob":"01-01-2001"
// }
