package service

import (
    "base-gin/domain/dto"
    "base-gin/exception"
    "base-gin/repository"
)

type AuthorService struct {
    repo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
    return &AuthorService{repo: authorRepo}
}

func (s *AuthorService) Create(params *dto.AuthorCreateReq) error {
    newAuthor := params.ToEntity()
    return s.repo.Create(&newAuthor)
}

func (s *AuthorService) GetByID(id uint) (dto.AuthorResp, error) {
    var resp dto.AuthorResp

    author, err := s.repo.GetByID(id)
    if err != nil {
        return resp, err
    }
    if author == nil {
        return resp, exception.ErrDataNotFound
    }

    resp.FromEntity(author)

    return resp, nil
}

func (s *AuthorService) GetList(params *dto.Filter) ([]dto.AuthorResp, error) {
    var resp []dto.AuthorResp

    authors, err := s.repo.GetList(params)
    if err != nil {
        return nil, err
    }
    if len(authors) < 1 {
        return nil, exception.ErrDataNotFound
    }

    for _, author := range authors {
        var a dto.AuthorResp
        a.FromEntity(&author)
        resp = append(resp, a)
    }

    return resp, nil
}

func (s *AuthorService) Update(params *dto.AuthorUpdateReq) error {
    if params.ID <= 0 {
        return exception.ErrDataNotFound
    }

    return s.repo.Update(params)
}

func (s *AuthorService) Delete(id uint) error {
    if id <= 0 {
        return exception.ErrDataNotFound
    }

    return s.repo.Delete(id)
}