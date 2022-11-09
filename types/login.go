package types

type UserLoginRequest struct {
	Email    string `msg:"email"`
	Password string `msg:"password"`
}
type UserLoginResponse struct {
	Token string `msg:"token"`
}
