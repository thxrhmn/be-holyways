package authdto

type AuthRequest struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"form"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"pasword" form:"pasword"`
}
