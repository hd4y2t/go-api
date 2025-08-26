package users

import (
	dto "go-api/src/moduls/users/dto"

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
	var input dto.UpdateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	return us.userRepository.Update(user)
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*User, error) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return us.userRepository.Delete(user)
}
