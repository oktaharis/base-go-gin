package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForm_Create_Success(t *testing.T) {
	params := dto.FormCreateReq{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}

	w := doTest(
		"POST",
		server.RootFormData,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestForm_Update_Success(t *testing.T) {
	o := dao.FormData{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}
	_ = formDataRepo.Create(&o)

	params := dto.FormDataUpdateReq{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    false,
		InputText:   "updated text",
		SelectBox:   "updated select",
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootFormData, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := formDataRepo.GetByID(o.ID)
	assert.Equal(t, params.RadioButton, item.RadioButton)
	assert.Equal(t, params.CheckBox, item.CheckBox)
	assert.Equal(t, params.InputText, item.InputText)
	assert.Equal(t, params.SelectBox, item.SelectBox)
	assert.Equal(t, false, item.DeletedAt.Valid)
}
func TestForm_Delete_Success(t *testing.T) {
	o := dao.FormData{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}
	_ = formDataRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootFormData, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := formDataRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestForm_GetList_Success(t *testing.T) {
	o1 := dao.FormData{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}
	_ = formDataRepo.Create(&o1) 
	

	o2 := dao.FormData{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}
	_ = formDataRepo.Create(&o2)

	w := doTest(
		"GET",
		server.RootFormData,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o1.RadioButton)
	assert.Contains(t, body, o2.RadioButton)

	w = doTest(
		"GET",
		server.RootFormData+"?q="+o1.RadioButton,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body = w.Body.String()
	assert.Contains(t, body, o1.RadioButton)
	assert.NotContains(t, body, o2.RadioButton)
}

func TestForm_GetDetail_Success(t *testing.T) {
	o := dao.FormData{
		RadioButton: util.RandomStringAlpha(6),
		CheckBox:    true,
		InputText:   "test",
		SelectBox:   "test",
	}
	_ = formDataRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootFormData, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o.RadioButton)
}
