package user

import (
	userEntity "gobook/src/user/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []userEntity.User
	FindById(id int) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Create(userEntity userEntity.User) (*userEntity.User, error)
	Update(userEntity userEntity.User) (*userEntity.User, error)
	Delete(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// FindAll implements UserRepository.
func (userRepository *UserRepositoryImpl) FindAll() []userEntity.User {

	var user []userEntity.User

	userRepository.db.Find(&user)

	return user
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (userRepository *UserRepositoryImpl) FindById(id int) (*userEntity.User, error) {
	// declare user variable
	var user userEntity.User

	// find user by id using gorm
	dataUser := userRepository.db.First(&user, id)

	// return error if any
	if dataUser.Error != nil {
		return nil, dataUser.Error
	}

	// return user
	return &user, nil
}

func (userRepository *UserRepositoryImpl) FindByEmail(email string) (*userEntity.User, error) {
	// declare user variable
	var user userEntity.User

	// find user by email using gorm
	dataUser := userRepository.db.Where("email = ?", email).First(&user)

	// return error if any
	if dataUser.Error != nil {
		return nil, dataUser.Error
	}

	// return user
	return &user, nil
}

func (userRepository *UserRepositoryImpl) Create(userEntity userEntity.User) (*userEntity.User, error) {
	newUser := userRepository.db.Create(&userEntity)

	if newUser.Error != nil {
		return nil, newUser.Error
	}

	return &userEntity, nil
}

func (userRepository *UserRepositoryImpl) Update(userEntity userEntity.User) (*userEntity.User, error) {
	updateUser := userRepository.db.Save(&userEntity)

	if updateUser.Error != nil {
		return nil, updateUser.Error
	}

	return &userEntity, nil
}

func (userRepository *UserRepositoryImpl) Delete(id int) error {
	deleteUser := userRepository.db.Delete(&userEntity.User{}, id)

	if deleteUser.Error != nil {
		return deleteUser.Error
	}

	return nil
}
