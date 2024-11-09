package rest

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/server"
	"base-gin/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	hr      *server.Handler
	service *service.BookService
}

func NewBookHandler(handler *server.Handler, BookService *service.BookService) *BookHandler {
	return &BookHandler{hr: handler, service: BookService}
}

func (h *BookHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootBook)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
//	@Summary Create a publisher
//	@Description Create a publisher.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param detail body dto.BookUpdateReq true "Publisher's detail"
//	@Success 201 {object} dto.SuccessResponse[any]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [post]
func (h *BookHandler) create(c *gin.Context) {
	var req dto.BookCreateReq
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

// getList godoc
//
//	@Summary Get a list of publishers
//	@Description Get a list of publishers.
//	@Produce json
//	@Param q query string false "Publisher's name"
//	@Param s query int false "Data offset"
//	@Param l query int false "Data limit"
//	@Success 200 {object} dto.SuccessResponse[[]dto.BookResp]
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [get]
func (h *BookHandler) getList(c *gin.Context) {
	var req dto.Filter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.GetList(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.BookResp]{
		Success: true,
		Message: "Daftar penerbit",
		Data:    data,
	})
}

// getByID godoc
//
//	@Summary Get a publisher's detail
//	@Description Get a publisher's detail.
//	@Produce json
//	@Param id path int true "Publisher's ID"
//	@Success 200 {object} dto.SuccessResponse[dto.BookResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [get]
func (h *BookHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.BookResp]{
		Success: true,
		Message: "Detail penerbit",
		Data:    data,
	})
}

// update godoc
//
//	@Summary Update a publisher's detail
//	@Description Update a publisher's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Publisher's ID"
//	@Param detail body dto.BookUpdateReq true "Publisher's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [put]
func (h *BookHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.BookUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil disimpan",
	})
}

// delete godoc
//
//	@Summary Delete a publisher
//	@Description Delete a publisher.
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Publisher's ID"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [delete]
func (h *BookHandler) delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil dihapus",
	})
}
