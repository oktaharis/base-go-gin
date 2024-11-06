package rest

import (
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublisherHandler struct {
	hr      *server.Handler
	service *service.PublisherService
}

func NewPublisherHandler(handler *server.Handler, publisherService *service.PublisherService) *PublisherHandler {
	return &PublisherHandler{hr: handler, service: publisherService}
}

func (h *PublisherHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootPublisher)
	grp.POST("", h.hr.AuthAccess(), h.create)
}

func (h *PublisherHandler) create(c *gin.Context) {
	var req dto.PublisherCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil disimpan",
	})
}
