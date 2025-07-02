package request

type CreateUser_Role struct {
	User_ID string `json:"user_id"`
	Role_ID string `json:"role_id"`
}

type UpdateUser_Role struct {
	CreateRole
}

type ListUser_Role struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
}

type GetByIDUser_Role struct {
	ID string `uri:"id" binding:"required"`
}
