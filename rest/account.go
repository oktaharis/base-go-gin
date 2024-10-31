package rest

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/server"
	"base-gin/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	hr      *server.Handler
	service *service.AccountService
}

func NewAccountHandler(
	hr *server.Handler,
	accountService *service.AccountService,
) *AccountHandler {
	return &AccountHandler{hr: hr, service: accountService}
}

func (h *AccountHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootAccount)
	grp.POST(server.PathLogin, h.login)
}

// login godoc
//
//	@Summary Account login
//	@Description Account login using username & password combination.
//	@Accept json
//	@Produce json
//	@Param cred body dto.AccountLoginReq true "Credential"
//	@Success 200 {object} dto.SuccessResponse[dto.AccountLoginResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /account/login [post]
func (h *AccountHandler) login(c *gin.Context) {
	var req dto.AccountLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound),
			errors.Is(err, exception.ErrUserLoginFailed):
			c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(exception.ErrUserLoginFailed.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AccountLoginResp]{
		Success: true,
		Message: "Login berhasil",
		Data:    data,
	})
}
