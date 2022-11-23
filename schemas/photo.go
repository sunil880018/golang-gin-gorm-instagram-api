package schemas

type Photo struct {
	UserId         string `json:"userid" binding:"required"`
	PhotoPath      string `json:"photopath" binding:"required"`
	PhotoLatitude  int    `json:"photolatitude" binding:"required"`
	PhotoLongitude int    `json:"photolongitude" binding:"required"`
	UserLatitude   int    `json:"userlatitude" binding:"required"`
	UserLongitude  int    `json:"userlongitude" binding:"required"`
}
