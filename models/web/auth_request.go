package web

type AuthRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,lte=6" json:"password"`
}
