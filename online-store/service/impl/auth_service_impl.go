package service_impl

import (
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/logger"
	"online-store/model"
	"online-store/repository"
	"online-store/service"
	"online-store/validation"
	"time"
)

func NewAuthService(userRepository *repository.UserRepository, configuration config.Config) service.AuthService {
	return &authServiceImpl{
		UserRepository: *userRepository,
		Configuration: configuration,
	}
}

type authServiceImpl struct {
	UserRepository repository.UserRepository
	Configuration config.Config
}

func (service *authServiceImpl) Logout(token string) {
	service.UserRepository.DeleteToken(token)
}

func (service *authServiceImpl) Login(request model.LoginRequest) (response model.GetUserResponse) {
	validation.ValidateLogin(request)
	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil{
		panic(exception.ValidationError{
			Message: "invalid email or password",
			Status:  http.StatusUnauthorized,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil{
		panic(exception.ValidationError{
			Message: "invalid email or password",
			Status:  http.StatusUnauthorized,
		})
	}

	token := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	service.UserRepository.UpdateToken(user.Id, token)

	response = model.GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Balance:   user.Balance,
		Token:     token,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	request.Password = "******"
	logger.Info().Interface("Login: ", request).Msg("Login success")
	return response
}

func (service *authServiceImpl) Register(request model.RegisterRequest) (response model.GetUserResponse) {
	validation.ValidateRegister(request)

	password,_ := bcrypt.GenerateFromPassword([]byte(request.Password),5)

	user := entity.User{
		Id:        uuid.New().String(),
		Name:      request.Name,
		Email:     request.Email,
		Password:  string(password),
		Balance:   0,
		Token:     base64.StdEncoding.EncodeToString([]byte(uuid.New().String())),
		Address:   request.Address,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	service.UserRepository.Insert(user)

	response = model.GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Balance:   user.Balance,
		Token:     user.Token,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	request.Password = "******"
	logger.Info().Interface("Register: ", request).Msg("Register success")
	return response
}

func (service *authServiceImpl) GetProfile(token string) (response model.GetUserResponse) {
	user,err := service.UserRepository.FindByToken(token)
	if err!=nil{
		panic(exception.ValidationError{
			Message: "invalid user id",
			Status:  http.StatusUnauthorized,
		})
	}
	response = model.GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Balance:   user.Balance,
		Token:     user.Token,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	logger.Info().Interface("Profile: ", nil).Msg("Get profile success")
	return response
}

