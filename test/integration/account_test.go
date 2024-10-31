package integration_test

import (
	"base-gin/domain/dto"
	"base-gin/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount_Login_Success(t *testing.T) {
	req := dto.AccountLoginReq{
		Username: "admin",
		Password: password,
	}

	w := doTest("POST", server.RootAccount+server.PathLogin, req, "")
	assert.Equal(t, 200, w.Code)
}
