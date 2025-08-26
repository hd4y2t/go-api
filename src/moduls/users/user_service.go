package users

import (
	"github.com/gin-gonic/gin"
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
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	return us.userRepository.Create(user)
}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*User, error) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
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
