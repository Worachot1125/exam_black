package response

type EmergencyReportResponse struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	EmergencyTypeID string `json:"emergency_type_id"`
	Description      string `json:"description"`
	ImageURL         string `json:"image_url"`
	Status           string `json:"status"` // Should be one of the enum values
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
