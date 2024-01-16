package repositories

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/utils/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data user.UserDto) (user.UserDto, error)
	// LoginUser(data user.User) (user.UserDto, error)
	GetAllUser(filter user.UserDto, page, pageSize int) ([]user.UserDto, int, error)
	// CountUsersByRole(roleId uint) (int, error)
	// GetUser(id string) (user.UserDto, error)
	// //getRoleName(roleID uint) string
	// UpdateUser(id string, input user.UserDto) (user.UserDto, error)
	// DeleteUser(id string) (user.UserDto, error)
	// FindByEmail(email string) (user.UserDto, error)
	// CreateUser(user user.UserDto) (user.UserDto, error)
	//SaveOTP(otp *models.OTP) (*models.OTP, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data user.UserDto) (user.UserDto, error) {
	dataUser := user.ConvertDtoToModel(data)
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

// func (u *userRepository) LoginUser(data *request.User) (user.UserDto, error) {
// 	dataUser := domain.UserDtoToModel(*data)
// 	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
// 	if err != nil {
// 		return user.UserDto{}, errors.ERR_EMAIL_NOT_FOUND
// 	}

// 	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
// 	if err != nil {
// 		return user.UserDto{}, errors.ERR_WRONG_PASSWORD
// 	}
// 	var creator models.Creator
// 	err = u.db.Preload("Role").Preload("Users").First(&creator, "user_id = ?", dataUser.ID).Error
// 	if err != nil {
// 		creator.Users = *dataUser
// 		return *domain.ConvertFromModelToCreatorsRes(creator), nil
// 	}
// 	return *domain.ConvertFromModelToCreatorsRes(creator), nil
// }

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

	offset := (page - 1) * pageSize
	query = query.Limit(pageSize).Offset(offset)

	err := query.Find(&allUser).Error
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(allUser); i++ {
		userVm := user.ConvertModelToDto(allUser[i])
		resAllUser = append(resAllUser, *userVm)
	}

	var allItems int64
	query.Count(&allItems)

	return resAllUser, int(allItems), nil
}

// func (u *userRepository) CountUsersByRole(roleId uint) (int, error) {
// 	var count int64
// 	var query = u.db.Model(&models.Users{})

// 	query.Where("role_id = ?", roleId)

// 	err := query.Count(&count).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return int(count), nil
// }

// func (u *userRepository) GetUser(id string) (user.UserDto, error) {
// 	var userData models.Users
// 	err := u.db.Preload("Role").First(&userData, "id = ?", id).Error
// 	if err != nil {
// 		return user.UserDto{}, err
// 	}
// 	return *user.ConvertModelToDto(userData), nil
// }

// func (u *userRepository) UpdateUser(id string, input request.User) (user.UserDto, error) {
// 	userData := models.Users{}
// 	err := u.db.First(&userData, "id = ?", id).Error

// 	if err != nil {
// 		return user.UserDto{}, err
// 	}

// 	if input.FirstName != "" {
// 		userData.FirstName = input.FirstName
// 	}
// 	if input.LastName != "" {
// 		userData.LastName = input.LastName
// 	}
// 	if input.Username != "" {
// 		userData.Username = input.Username
// 	}
// 	if input.Email != "" {
// 		userData.Email = input.Email
// 	}
// 	if input.Password != "" {
// 		userData.Password = input.Password
// 	}
// 	if input.PhoneNumber != "" {
// 		userData.PhoneNumber = input.PhoneNumber
// 	}
// 	if input.NIK != "" {
// 		userData.NIK = input.NIK
// 	}
// 	if input.Gender != "" {
// 		userData.Gender = input.Gender
// 	}
// 	// if input.RoleId != 0 {
// 	// 	userData.RoleId = input.RoleId
// 	// }

// 	if err = u.db.Save(&userData).Error; err != nil {
// 		return user.UserDto{}, err
// 	}
// 	return *user.ConvertModelToDto(userData), nil
// }

// func (u *userRepository) DeleteUser(id string) (user.UserDto, error) {
// 	userData := models.Users{}
// 	res := user.UserDto{}

// 	find := u.db.Preload("Role").First(&userData, "id = ?", id).Error
// 	if find == nil {
// 		res = *user.ConvertModelToDto(userData)
// 	}

// 	err := u.db.Delete(&models.Creator{}, "user_id = ?", id).Error
// 	if err != nil {
// 		return user.UserDto{}, err
// 	}

// 	err = u.db.Delete(&userData, "id = ?", id).Error
// 	if err != nil {
// 		return user.UserDto{}, err
// 	}

// 	return res, nil
// }

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
