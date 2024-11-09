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

type FormDataHandler struct {
	hr      *server.Handler
	service *service.FormDataService
}

func NewFormDataHandler(handler *server.Handler, FormDataService *service.FormDataService) *FormDataHandler {
	return &FormDataHandler{hr: handler, service: FormDataService}
}

func (h *FormDataHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootFormData)
	grp.POST("", h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.update)
	grp.DELETE("/:id", h.delete)
}

// create godoc
//
//	@Summary Create an author
//	@Description Create an author.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param detail body dto.FormCreateReq true "Author's detail"
//	@Success 201 {object} dto.SuccessResponse[any]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /forms [post]
func (h *FormDataHandler) create(c *gin.Context) {
	var req dto.FormCreateReq
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
		Message: "Form berhasil disimpan",
	})
}


// getList godoc
//
//	@Summary Get a list of forms
//	@Description Get a list of forms.
//	@Produce json
//	@Param q query string false "Author's name"
//	@Param s query int false "Data offset"
//	@Param l query int false "Data limit"
//	@Success 200 {object} dto.SuccessResponse[[]dto.AuthorResp]
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /forms [get]
func (h *FormDataHandler) getList(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.FormDataResp]{
		Success: true,
		Message: "Daftar penulis",
		Data:    data,
	})
}

// getByID godoc
//
//	@Summary Get an author's detail
//	@Description Get an author's detail.
//	@Produce json
//	@Param id path int true "Author's ID"
//	@Success 200 {object} dto.SuccessResponse[dto.AuthorResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /forms/{id} [get]
func (h *FormDataHandler) getByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.FormDataResp]{
		Success: true,
		Message: "Detail penulis",
		Data:    data,
	})
}

// update godoc
//
//	@Summary Update an author's detail
//	@Description Update an author's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Author's ID"
//	@Param detail body dto.FormDataUpdateReq true "Author's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /forms/{id} [put]
func (h *FormDataHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.FormDataUpdateReq
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
		Message: "Data author berhasil diperbarui",
	})
}

// delete godoc
//
//	@Summary Delete an author
//	@Description Delete an author.
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Author's ID"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /forms/{id} [delete]
func (h *FormDataHandler) delete(c *gin.Context) {
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
		Message: "Author berhasil dihapus",
	})
}
