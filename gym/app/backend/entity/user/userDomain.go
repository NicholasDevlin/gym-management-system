package user

import "gorm.io/gorm"

func ConvertReqToDto(input UserReq) *UserDto {
	return &UserDto{
		Id:             input.Id,
		UUID:           input.UUID,
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
		UUID:           input.UUID,
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
		UUID:           input.UUID,
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
		UUID:           input.UUID,
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
