package integration_test

import (
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublisher_Create_Success(t *testing.T) {
	params := dto.PublisherCreateReq{
		Name: util.RandomStringAlpha(6),
		City: util.RandomStringAlpha(8),
	}

	w := doTest("POST", server.RootPublisher, params, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 201, w.Code)
}
