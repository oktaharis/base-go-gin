package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// parseDate function to parse date strings into time.Time in UTC timezone
func parseDate(dateStr string) time.Time {
	parsedDate, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.UTC)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}
	return parsedDate
}

func TestAuthor_Create_Success(t *testing.T) {
	params := dto.AuthorCreateReq{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "m",
		BirthDate: parseDate("2015-09-02 08:04:00"),
	}

	w := doTest(
		"POST",
		server.RootAuthor,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestAuthor_Update_Success(t *testing.T) {
	o := dao.Author{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "f",
		BirthDate: parseDate("1995-05-15 00:00:00"),
	}
	_ = authorRepo.Create(&o)

	params := dto.AuthorUpdateReq{
		Fullname:  util.RandomStringAlpha(12),
		Gender:    "m",
		BirthDate: parseDate("1992-03-20 00:00:00"),
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)

	// Convert item.BirthDate to UTC for comparison
	assert.Equal(t, params.Fullname, item.Fullname)
	assert.Equal(t, params.Gender, item.Gender)
	assert.Equal(t, params.BirthDate.UTC(), item.BirthDate.UTC()) // Normalize both to UTC
}

func TestAuthor_Delete_Success(t *testing.T) {
	o := dao.Author{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "f",
		BirthDate: parseDate("1995-05-15 00:00:00"),
	}
	_ = authorRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestAuthor_GetList_Success(t *testing.T) {
	o1 := dao.Author{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "m",
		BirthDate: parseDate("1990-01-01 00:00:00"),
	}
	_ = authorRepo.Create(&o1)

	o2 := dao.Author{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "f",
		BirthDate: parseDate("1992-03-20 00:00:00"),
	}
	_ = authorRepo.Create(&o2)

	w := doTest(
		"GET",
		server.RootAuthor,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.Contains(t, body, o2.Fullname)

	w = doTest(
		"GET",
		server.RootAuthor+"?q="+o1.Fullname,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body = w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.NotContains(t, body, o2.Fullname)
}

func TestAuthor_GetDetail_Success(t *testing.T) {
	o := dao.Author{
		Fullname:  util.RandomStringAlpha(10),
		Gender:    "m",
		BirthDate: parseDate("1990-01-01 00:00:00"),
	}
	_ = authorRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o.Fullname)
}
