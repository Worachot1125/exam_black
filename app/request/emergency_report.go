package request

type CreateEmergencyReport struct {
	User_ID           string `json:"user_id"`
	Emergency_Type_ID string `json:"emergency_type_id"`
	Description       string `json:"description"`
	Image_URL         string `json:"image_url"`
	Status            string `json:"status"` // Should be one of the enum values
}

type UpdateEmergencyReport struct {
	CreateEmergencyReport
}

type ListEmergencyReport struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Search   string `form:"search"`
	SearchBy string `form:"search_by"`
	SortBy   string `form:"sort_by"`
	OrderBy  string `form:"order_by"`
}

type GetByIDEmergencyReport struct {
	ID string `uri:"id" binding:"required"`
}

type GetByUserIDEmergency struct {
    UserID string `uri:"id" binding:"required"`
}