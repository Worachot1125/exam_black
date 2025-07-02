package request

type CreateUser struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	UserNumber   string `json:"user_number"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	RoleID        string `json:"role_id"` // Assuming RoleID is a string,

}

type UpdateUser struct {
	CreateUser
}

type ListUser struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
}

type GetByIDUser struct {
	ID string `uri:"id" binding:"required"`
}
