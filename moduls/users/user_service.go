package users

import (
	dto "go-api/moduls/users/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetAll() []User
	GetById(id int) User
	Create(ctx *gin.Context) (*User, error)
	Update(ctx *gin.Context) (*User, error)
	Delete(ctx *gin.Context) (*User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (us *UserServiceImpl) GetAll() []User {
	return us.userRepository.FindAll()
}

func (us *UserServiceImpl) GetById(id int) User {
	return us.userRepository.FindById(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context) (*User, error) {
	var input dto.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateUserInput

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	user := User{
		ID:    uint(id),
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := User{
		ID: uint(id),
	}

	result, err := us.userRepository.Delete(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
