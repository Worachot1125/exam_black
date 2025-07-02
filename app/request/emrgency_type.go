package request

type CreateEmergencyType struct {
	Name string `json:"name"`
}

type UpdateEmergencyType struct {
	CreateEmergencyType
}

type ListEmergencyType struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
}

type GetByIDEmergencyType struct {
	ID string `uri:"id" binding:"required"`
}
