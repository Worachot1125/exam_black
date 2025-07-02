package response

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserNumber string `json:"user_number"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	RoleID    string `json:"role_id"` // Assuming RoleID is a string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
