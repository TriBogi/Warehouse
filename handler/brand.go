package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/brand"
	"warehouse/helper"
)

type brandHandler struct {
	brand.Service
}

func NewBrandHandler(brandService brand.Service) *brandHandler {
	return &brandHandler{brandService}

}

func (h *brandHandler) RegisterBrandHandler(c *gin.Context) {

	var input brand.Brand

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, input)
}

func (h *brandHandler) GetBrandHandler(c *gin.Context) {

	result, err := h.Service.GetBrand()
	if err != nil {
		response := helper.APIResponse("data brand failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("data brand succesfuly", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *brandHandler) DeleteBrandHandler(c *gin.Context) {

	delBrand, err := h.Service.DeleteBrand()
	if err != nil {
		response := helper.APIResponse("delete brand failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("data brand succesfuly", http.StatusOK, "success", delBrand)
	c.JSON(http.StatusOK, response)
}