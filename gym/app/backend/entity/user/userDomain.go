package user

import "gorm.io/gorm"

func ConvertReqToDto(input UserReq) *UserDto {
	return &UserDto{
		Id:             input.Id,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
	}
}

func ConvertDtoToModel(input UserDto) *User {
	return &User{
		Model: gorm.Model{
			ID: input.Id,
		},
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
	}
}

func ConvertModelToDto(input User) *UserDto {
	return &UserDto{
		Id:             input.ID,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
	}
}

func ConvertDtoToRes(input UserDto) *UserRes {
	return &UserRes{
		Id:             input.Id,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
	}
}
