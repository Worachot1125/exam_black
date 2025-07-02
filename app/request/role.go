package request

type CreateRole struct {
	Name string `json:"name"`
}

type UpdateRole struct {
	CreateRole
}

type ListRole struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
}

type GetByIDRole struct {
	ID string `uri:"id" binding:"required"`
}
