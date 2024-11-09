package service

import (
    "base-gin/domain/dto"
    "base-gin/exception"
    "base-gin/repository"
)

type FormDataService struct {
    repo *repository.FormDataRepository
}

func NewFormDataService(formDataRepo *repository.FormDataRepository) *FormDataService {
    return &FormDataService{repo: formDataRepo}
}

func (s *FormDataService) Create(params *dto.FormCreateReq) error {
    newFormData := params.ToEntity()
    return s.repo.Create(&newFormData)
}

func (s *FormDataService) GetByID(id uint) (dto.FormDataResp, error) {
    var resp dto.FormDataResp

    form, err := s.repo.GetByID(id)
    if err != nil {
        return resp, err
    }
    if form == nil {
        return resp, exception.ErrDataNotFound
    }

    resp.FromEntity(form)

    return resp, nil
}

func (s *FormDataService) GetList(params *dto.Filter) ([]dto.FormDataResp, error) {
    var resp []dto.FormDataResp

    formDatas, err := s.repo.GetList(params)
    if err != nil {
        return nil, err
    }
    if len(formDatas) < 1 {
        return nil, exception.ErrDataNotFound
    }

    for _, formData := range formDatas {
        var a dto.FormDataResp
        a.FromEntity(&formData)
        resp = append(resp, a)
    }

    return resp, nil
}
func (s *FormDataService) Update(params *dto.FormDataUpdateReq) error {
    if params.ID <= 0 {
        return exception.ErrDataNotFound
    }

    return s.repo.Update(params)
}

func (s *FormDataService) Delete(id uint) error {
    if id <= 0 {
        return exception.ErrDataNotFound
    }

    return s.repo.Delete(id)
}