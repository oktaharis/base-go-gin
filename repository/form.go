package repository

import (
    "base-gin/domain/dao"
    "base-gin/domain/dto"
    "base-gin/exception"
    "base-gin/storage"
    "errors"
    "fmt"
    "gorm.io/gorm"
)

type FormDataRepository struct {
    db *gorm.DB
}

func NewFormDataRepository(db *gorm.DB) *FormDataRepository {
    return &FormDataRepository{db: db}
}

func (r *FormDataRepository) Create(newFormData *dao.FormData) error {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    tx := r.db.WithContext(ctx).Create(&newFormData)
    if tx.Error != nil {
        return tx.Error
    }

    return nil
}

func (r *FormDataRepository) GetByID(id uint) (*dao.FormData, error) {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    var formData dao.FormData
    tx := r.db.WithContext(ctx).First(&formData, id)
    if tx.Error != nil {
        if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
            return nil, exception.ErrUserNotFound
        }
        return nil, tx.Error
    }

    return &formData, nil
}

func (r *FormDataRepository) GetList(params *dto.Filter) ([]dao.FormData, error) {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    var formDatas []dao.FormData
    tx := r.db.WithContext(ctx)

    if params.Keyword != "" {
        q := fmt.Sprintf("%%%s%%", params.Keyword)
        tx = tx.Where("radio_button LIKE ?", q)
    }
    if params.Start >= 0 {
        tx = tx.Offset(params.Start)
    }
    if params.Limit > 0 {
        tx = tx.Limit(params.Limit)
    }

    tx = tx.Order("radio_button ASC").Find(&formDatas)
    if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return nil, tx.Error
    }

    return formDatas, nil
}

func (r *FormDataRepository) Update(params *dto.FormDataUpdateReq) error {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    tx := r.db.WithContext(ctx).Model(&dao.FormData{}).
        Where("id = ?", params.ID).
        Updates(map[string]interface{}{
            "radio_button":  params.RadioButton,
            "check_box":    params.CheckBox,
            "input_text": params.InputText,
			"select_box": params.SelectBox,
        })

    return tx.Error
}

func (r *FormDataRepository) Delete(id uint) error {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    tx := r.db.WithContext(ctx).Delete(&dao.FormData{}, id)

    return tx.Error
}
