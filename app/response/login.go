package response

type FindUser_Number struct {
	User_Number string `bun:"user_number" json:"user_number"`
}

type LoginResponse struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	UserNumber string `json:"user_number"`
	Role_ID    string `json:"role_id"`
}
