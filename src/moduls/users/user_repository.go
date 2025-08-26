package users

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() []User
	FindById(id int) User
	Create(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur *UserRepositoryImpl) FindAll() []User {
	var users []User

	_ = ur.db.Find(&users)

	return users
}

func (ur *UserRepositoryImpl) FindById(id int) User {
	var user User
	_ = ur.db.First(&user, id)

	return user
}

func (ur *UserRepositoryImpl) Create() []User {

}

func (ur *UserRepositoryImpl) Update() []User {

}

func (ur *UserRepositoryImpl) Delete() []User {

}
