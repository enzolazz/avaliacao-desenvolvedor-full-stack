package dtos

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserInfoResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
}

type RefreshRequest struct {
	Token string `json:"token"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}
