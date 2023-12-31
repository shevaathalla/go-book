package user

import (
	"database/sql"
	userDto "gobook/src/user/dto"
	userEntity "gobook/src/user/entity"
	userRepository "gobook/src/user/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll(offset int, limit int) []userEntity.User
	FindById(id int) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Verify(id int) (*userEntity.User, error)
	Create(userDto userDto.CreateUserRequest) (*userEntity.User, error)
}

type UserServiceImpl struct {
	userRepository userRepository.UserRepository
}

// FindByEmail implements UserService.
func (userService *UserServiceImpl) FindByEmail(email string) (*userEntity.User, error) {
	return userService.userRepository.FindByEmail(email)
}

// Verify implements UserService.
func (userService *UserServiceImpl) Verify(id int) (*userEntity.User, error) {

	user, err := userService.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	user.EmailVerifiedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	// update user to database
	updatedUser, err := userService.userRepository.Update(*user)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Create implements UserService.
func (userService *UserServiceImpl) Create(userDto userDto.CreateUserRequest) (*userEntity.User, error) {
	var user userEntity.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Name = userDto.Name
	user.Email = userDto.Email
	user.Password = string(hashedPassword)
	user.Address = userDto.Address
	user.Phone = userDto.Phone
	user.Instance = userDto.Instance

	// create new user to database
	userService.userRepository.Create(user)

	return &user, nil
}

// FindAll implements UserService.
func (userService *UserServiceImpl) FindAll(offset int, limit int) []userEntity.User {
	return userService.userRepository.FindAll(offset, limit)
}

// FindById implements UserService.
func (userService *UserServiceImpl) FindById(id int) (*userEntity.User, error) {
	return userService.userRepository.FindById(id)
}

func NewUserService(userRepository userRepository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}
