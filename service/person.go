package service

import (
	"base-gin/domain/dto"
	"base-gin/repository"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(personRepo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: personRepo}
}

func (s *PersonService) GetAccountProfile(accountID uint) (dto.AccountProfileResp, error) {
	var resp dto.AccountProfileResp

	item, err := s.repo.GetByAccountID(accountID)
	if err != nil {
		return resp, err
	}

	resp.FromPerson(&item)

	return resp, nil
}
