package user

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
