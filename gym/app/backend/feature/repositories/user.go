package repositories

import (
	"gym/app/backend/models/role"
	"gym/app/backend/models/user"
	"gym/app/backend/utils/bcrypt"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data user.UserDto) (user.UserDto, error)
	LoginUser(data user.UserDto) (user.UserDto, error)
	GetAllUser(filter user.UserFilter) ([]user.UserDto, error)
	GetUser(filter user.UserDto) (user.UserDto, error)
	UpdateUser(data, input user.UserDto) (user.UserDto, error)
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

func (u *userRepository) GetAllUser(filter user.UserFilter) ([]user.UserDto, error) {
	var allUser []user.User
	var resAllUser []user.UserDto

	query := u.db.Joins("Role")
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
	if filter.Active != nil {
		if *filter.Active {
			today := time.Now()
			date := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
			query = query.Where("subscription_expiration_date >= ?", date)
		}
		if !*filter.Active {
			today := time.Now()
			date := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
			query = query.Where("subscription_expiration_date < ?", date)
		}
	}
	if filter.Role != "" {
		query = query.Joins("Role", query.Where("role = ?", filter.Role))
	}
	if filter.LastDayActive {
		today := time.Now()
		date := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
		query = query.Where("subscription_expiration_date >= ? AND subscription_expiration_date < ?", date, date.AddDate(0, 0, 1))
	}
	query = query.Order("display_name")

	err := query.Find(&allUser).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allUser); i++ {
		user := user.ConvertModelToDto(allUser[i])
		resAllUser = append(resAllUser, *user)
	}

	return resAllUser, nil
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
		if input.Gender != consts.MALE && input.Gender != consts.FEMALE {
			userData.Gender = consts.SECRET
		}
	}
	if input.SubscriptionExpirationDate != nil {
		userData.SubscriptionExpirationDate = input.SubscriptionExpirationDate
	}
	if input.RoleId != 0 {
		userData.RoleId = input.RoleId
		userData.Role = role.Role{
			Role: input.Role.Role,
			Model: gorm.Model{
				ID: input.RoleId,
			},
		}
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
