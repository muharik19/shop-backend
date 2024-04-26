package services

import (
	"net/http"

	"github.com/muharik19/shop-backend/constant"
	"github.com/muharik19/shop-backend/pkg/util"

	"github.com/google/uuid"
	"github.com/muharik19/shop-backend/models"
	"github.com/muharik19/shop-backend/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) CreateUser(user *models.User) (res *models.Response, err error) {
	_, err = service.UserRepository.GetUserByEmail(user.Email)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Email already exist",
		}, nil
	}

	_, err = service.UserRepository.GetUserByPhoneNumber(user.PhoneNumber)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Phone Number already exist",
		}, nil
	}

	password, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New().String()
	user.Password = password
	err = service.UserRepository.AddUser(user)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "User created successfully",
	}, nil
}
