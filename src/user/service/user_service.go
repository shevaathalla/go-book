package user

import (
	userDto "gobook/src/user/dto"
	userEntity "gobook/src/user/entity"
	userRepository "gobook/src/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll() []userEntity.User
	FindById(id int) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Create(userDto userDto.CreateUserRequestBody) (*userEntity.User, error)
	Update(id int, userDto userDto.UpdateUserRequestBody) (*userEntity.User, error)
	Delete(id int) error
}

type UserServiceImpl struct {
	userRepository userRepository.UserRepository
}

func NewUserService(userRepository userRepository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (userService *UserServiceImpl) FindAll() []userEntity.User {
	return userService.userRepository.FindAll()
}

func (userService *UserServiceImpl) FindById(id int) (*userEntity.User, error) {
	return userService.userRepository.FindById(id)
}

func (userService *UserServiceImpl) FindByEmail(email string) (*userEntity.User, error) {
	return userService.userRepository.FindByEmail(email)
}

func (userService *UserServiceImpl) Create(userDto userDto.CreateUserRequestBody) (*userEntity.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dataUser := userEntity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: string(hashedPassword),
		Address:  userDto.Address,
		Instance: userDto.Instance,
		Phone:    userDto.Phone,
	}

	createdUser, err := userService.userRepository.Create(dataUser)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (userService *UserServiceImpl) Update(id int, userDto userDto.UpdateUserRequestBody) (*userEntity.User, error) {

	user, err := userService.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	user.Name = userDto.Name
	user.Address = userDto.Address
	user.Phone = userDto.Phone
	user.Instance = userDto.Instance

	updatedUser, err := userService.userRepository.Update(*user)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (userService *UserServiceImpl) Delete(id int) error {
	return userService.userRepository.Delete(id)
}
