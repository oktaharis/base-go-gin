package service

import (
	"base-gin/domain/dto"
	"base-gin/repository"
)

type PublisherService struct {
	repo *repository.PublisherRepository
}

func NewPublisherService(publisherRepo *repository.PublisherRepository) *PublisherService {
	return &PublisherService{repo: publisherRepo}
}

func (s *PublisherService) Create(params *dto.PublisherCreateReq) error {
	newItem := params.ToEntity()
	return s.repo.Create(&newItem)
}
