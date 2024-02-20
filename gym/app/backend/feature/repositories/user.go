package repositories

import (
	"gym/app/backend/models/user"
	"gym/app/backend/utils/bcrypt"
	"gym/app/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data user.UserDto) (user.UserDto, error)
	LoginUser(data user.UserDto) (user.UserDto, error)
	GetAllUser(filter user.UserDto, page, pageSize int) ([]user.UserDto, int, error)
	GetUser(filter user.UserDto) (user.UserDto, error)
	UpdateUser(data ,input user.UserDto) (user.UserDto, error)
	DeleteUser(id string) (user.UserDto, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data user.UserDto) (user.UserDto, error) {
	dataUser := user.ConvertDtoToModel(data)
	dataUser.UUID = uuid.NewV4()
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return user.UserDto{}, err
	}
	err = u.db.Preload("Role").First(&dataUser, "id = ?", dataUser.ID).Error
	if err != nil {
		return user.UserDto{}, errors.ERR_LOGIN
	}
	return *user.ConvertModelToDto(*dataUser), nil
}

func (u *userRepository) LoginUser(data user.UserDto) (user.UserDto, error) {
	dataUser := user.ConvertDtoToModel(data)
	err := u.db.Preload("Role").First(&dataUser, "email = ?", dataUser.Email).Error
	if err != nil {
		return user.UserDto{}, errors.ERR_EMAIL_NOT_FOUND
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return user.UserDto{}, errors.ERR_WRONG_PASSWORD
	}

	return *user.ConvertModelToDto(*dataUser), nil
}

func (u *userRepository) GetAllUser(filter user.UserDto, page, pageSize int) ([]user.UserDto, int, error) {
	var allUser []user.User
	var resAllUser []user.UserDto

	query := u.db.Preload("Role")
	if filter.DisplayName != "" {
		query = query.Where("display_name LIKE ? ", "%"+filter.DisplayName+"%")
	}
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}

	offset := (page - 1) * pageSize
	query = query.Limit(pageSize).Offset(offset)

	err := query.Find(&allUser).Error
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(allUser); i++ {
		user := user.ConvertModelToDto(allUser[i])
		resAllUser = append(resAllUser, *user)
	}

	var allItems int64
	query.Count(&allItems)

	return resAllUser, int(allItems), nil
}

func (u *userRepository) GetUser(filter user.UserDto) (user.UserDto, error) {
	var userData user.User
	query := u.db.Preload("Role")
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	err := query.First(&userData).Error
	if err != nil {
		return user.UserDto{}, err
	}
	return *user.ConvertModelToDto(userData), nil
}

func (u *userRepository) UpdateUser(data, input user.UserDto) (user.UserDto, error) {
	userData := *user.ConvertDtoToModel(data)

	if input.DisplayName != "" {
		userData.DisplayName = input.DisplayName
	}
	if input.Email != "" {
		userData.Email = input.Email
	}
	if input.Password != "" {
		userData.Password = input.Password
	}
	if input.PhoneNumber != "" {
		userData.PhoneNumber = input.PhoneNumber
	}
	if input.BirthDate != nil {
		userData.BirthDate = input.BirthDate
	}
	if input.Gender != "" {
		userData.Gender = input.Gender
	}
	if input.RoleId != 0 {
		userData.RoleId = input.RoleId
	}

	if err := u.db.Save(&userData).Error; err != nil {
		return user.UserDto{}, err
	}
	return *user.ConvertModelToDto(userData), nil
}

func (u *userRepository) DeleteUser(id string) (user.UserDto, error) {
	userData := user.User{}

	err := u.db.Delete(&userData, "uuid = ?", id).Error
	if err != nil {
		return user.UserDto{}, err
	}

	return *user.ConvertModelToDto(userData), nil
}

// func (u *userRepository) FindByEmail(email string) (*models.Users, error) {
// 	user := models.Users{}
// 	res := u.db.Where("email = ?", email).First(&user).Error
// 	if res != nil {
// 		return nil, res
// 	}
// 	return &user, nil
// }

// func (u *userRepository) CreateUser(user *models.Users) (*models.Users, error) {
// 	result := u.db.Create(&user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return user, nil
// }
