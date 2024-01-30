package services

import (
	"gym/app/backend/feature/repositories"
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IMembershipPlanService interface {
	CreateMembershipPlan(input membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error)
	GetAllMembershipPlan(filter membershipplan.MembershipPlanReq) ([]membershipplan.MembershipPlanRes, error)
	GetMembershipPlan(filter membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error)
	UpdateMembershipPlan(input membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error)
	DeleteMembershipPlan(id uuid.UUID) (membershipplan.MembershipPlanRes, error)
}

type membershipPlanService struct {
	membershipPlanRepository repositories.IMembershipPlanRepository
}

func NewMembershipPlanService(repo repositories.IMembershipPlanRepository) *membershipPlanService {
	return &membershipPlanService{membershipPlanRepository: repo}
}

func (mp *membershipPlanService) CreateMembershipPlan(input membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error) {
	if input.Name == "" {
		return membershipplan.MembershipPlanRes{}, errors.ERR_MEMBERSHIP_NAME_EMPTY
	}
	if input.Duration == 0 {
		return membershipplan.MembershipPlanRes{}, errors.ERR_MEMBERSHIP_DURATION_EMPTY
	}
	if input.Price == 0 {
		return membershipplan.MembershipPlanRes{}, errors.ERR_MEMBERSHIP_PRICE_EMPTY
	}

	res, err := mp.membershipPlanRepository.CreateMembershipPlan(*membershipplan.ConvertReqToDto(input))
	if err != nil {
		return membershipplan.MembershipPlanRes{}, errors.ERR_BCRYPT_PASSWORD
	}
	return *membershipplan.ConvertDtoToRes(res), nil
}

func (mp *membershipPlanService) GetAllMembershipPlan(filter membershipplan.MembershipPlanReq) ([]membershipplan.MembershipPlanRes, error) {
	res, err := mp.membershipPlanRepository.GetAllMembershipPlan(*membershipplan.ConvertReqToDto(filter))
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	var resMembershipPlan []membershipplan.MembershipPlanRes
	for i := 0; i < len(res); i++ {
		roleVm := membershipplan.ConvertDtoToRes(res[i])
		resMembershipPlan = append(resMembershipPlan, *roleVm)
	}
	return resMembershipPlan, nil
}

func (mp *membershipPlanService) GetMembershipPlan(filter membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error) {
	res, err := mp.membershipPlanRepository.GetMembershipPlan(*membershipplan.ConvertReqToDto(filter))

	if err != nil || (filter.Id == 0 && filter.UUID == uuid.Nil) {
		return membershipplan.MembershipPlanRes{}, errors.ERR_NOT_FOUND
	}
	return *membershipplan.ConvertDtoToRes(res), nil
}

func (mp *membershipPlanService) UpdateMembershipPlan(input membershipplan.MembershipPlanReq) (membershipplan.MembershipPlanRes, error) {
	res, err := mp.membershipPlanRepository.GetMembershipPlan(*membershipplan.ConvertReqToDto(input))
	if err != nil {
		return membershipplan.MembershipPlanRes{}, errors.ERR_NOT_FOUND
	}
	res, err = mp.membershipPlanRepository.UpdateMembershipPlan(res, *membershipplan.ConvertReqToDto(input))
	if err != nil {
		return membershipplan.MembershipPlanRes{}, errors.ERR_UPDATE_ROLE
	}
	return *membershipplan.ConvertDtoToRes(res), nil
}

func (mp *membershipPlanService) DeleteMembershipPlan(id uuid.UUID) (membershipplan.MembershipPlanRes, error) {
	res, err := mp.membershipPlanRepository.GetMembershipPlan(membershipplan.MembershipPlanDto{UUID: id})
	if err != nil {
		return membershipplan.MembershipPlanRes{}, errors.ERR_NOT_FOUND
	}

	res, err = mp.membershipPlanRepository.DeleteMembershipPlan(id.String())
	if err != nil {
		return membershipplan.MembershipPlanRes{}, errors.ERR_DELETE_USER
	}
	return *membershipplan.ConvertDtoToRes(res), nil
}
