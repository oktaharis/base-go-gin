package dto

import (
    "base-gin/domain/dao"
    "time"
)

// Request untuk membuat data baru
type FormCreateReq struct {
    RadioButton string `json:"radio_button" binding:"required,min=2,max=56"`
    CheckBox    bool   `json:"check_box"`
    InputText   string `json:"input_text" binding:"required"`
    SelectBox   string `json:"select_box" binding:"required"`
}

func (a *FormCreateReq) ToEntity() dao.FormData {
    return dao.FormData{
        RadioButton: a.RadioButton,
        CheckBox:    a.CheckBox,
        InputText:   a.InputText,
        SelectBox:   a.SelectBox,
        UpdatedAt:   time.Now(),
    }
}

// Response data
type FormDataResp struct {
    ID          uint      `json:"id"`
    RadioButton string    `json:"radio_button"`
    CheckBox    bool      `json:"check_box"`
    InputText   string    `json:"input_text"`
    SelectBox   string    `json:"select_box"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func (a *FormDataResp) FromEntity(formData *dao.FormData) {
    a.ID = formData.ID
    a.RadioButton = formData.RadioButton
    a.CheckBox = formData.CheckBox
    a.InputText = formData.InputText
    a.SelectBox = formData.SelectBox
    a.UpdatedAt = formData.UpdatedAt
}

// Request untuk update data
type FormDataUpdateReq struct {
	ID          uint   `json:"id"`
    RadioButton string `json:"radio_button" binding:"required,min=2,max=56"`
    CheckBox    bool   `json:"check_box"`
    InputText   string `json:"input_text" binding:"required"`
    SelectBox   string `json:"select_box" binding:"required"`
}
