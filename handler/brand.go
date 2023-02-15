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

	var input brand.CreateBrandInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	newBrand, err := h.Service.RegisterBrand(input)
	if err != nil {
		response := helper.APIResponse("Failed to create brand", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to create brand", http.StatusOK, "success", newBrand)
	c.JSON(http.StatusOK, response)
}

func (h *brandHandler) GetBrandsHandler(c *gin.Context) {

	result, err := h.Service.GetAllBrands()
	if err != nil {
		response := helper.APIResponse("data brand failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("data brand succesfuly", http.StatusOK, "success", result)

	c.JSON(http.StatusOK, response)
}

func (h *brandHandler) DeleteBrandHandler(c *gin.Context) {
	var inputID brand.GetBrandDetailInput

	err := h.Service.Delete(inputID)
	if err != nil {
		response := helper.APIResponse("delete brand failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("data brand succesfuly deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *brandHandler) UpdateBrandHandler(c *gin.Context) {
	var inputID brand.GetBrandDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData brand.CreateBrandInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update brand", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedBrand, err := h.Service.UpdateBrand(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update brand", http.StatusOK, "success", updatedBrand)
	c.JSON(http.StatusOK, response)
}

func (h *brandHandler) GetBrandByIDHandler(c *gin.Context) {
	var inputID brand.GetBrandDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Wrong Parameter", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := h.Service.GetBrandByID(inputID)
	if err != nil {
		response := helper.APIResponse("data brand failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("data brand succesfuly", http.StatusOK, "success", result)

	c.JSON(http.StatusOK, response)
}
