package web

type AuthRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,gte=6" json:"password"`
}
