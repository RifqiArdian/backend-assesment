package service

import "online-store/model"

type AuthService interface {
	Login(request model.LoginRequest)(response model.GetUserResponse)
	Register(request model.RegisterRequest)(response model.GetUserResponse)
	GetProfile(token string)(response model.GetUserResponse)
	Logout(token string)()
}
