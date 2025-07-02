package request

type LoginRequest struct {
	User_Number string `json:"user_number"`
	Password string `json:"password"`
}