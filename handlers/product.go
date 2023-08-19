package handlers

import (
	"fmt"
	"net/http"
	productdto "nutech/dto/product"
	dto "nutech/dto/result"
	"nutech/models"
	"nutech/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	selling, _ := strconv.Atoi(c.FormValue("selling_price"))
	purchase, _ := strconv.Atoi(c.FormValue("purchase_price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	request := productdto.CreateProductRequest{
		Name:          c.FormValue("name"),
		SellingPrice:  selling,
		PurchasePrice: purchase,
		Stock:         stock,
		Image:         dataFile,
	}

	validion := validator.New()
	err := validion.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	product := models.Product{
		Name:          request.Name,
		SellingPrice:  request.SellingPrice,
		PurchasePrice: request.PurchasePrice,
		Stock:         request.Stock,
		Image:         request.Image,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerProduct) FindProduct(c echo.Context) error {
	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("This is data file", dataFile)

	id, _ := strconv.Atoi(c.Param("id"))
	product, _ := h.ProductRepository.GetProduct(id)

	selling, _ := strconv.Atoi(c.FormValue("selling_price"))
	purchase, _ := strconv.Atoi(c.FormValue("purchase_price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	request := productdto.UpdateProductRequest{
		Name:          c.FormValue("name"),
		SellingPrice:  selling,
		PurchasePrice: purchase,
		Stock:         stock,
		Image:         dataFile,
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.SellingPrice != 0 {
		product.SellingPrice = request.SellingPrice
	}
	if request.PurchasePrice != 0 {
		product.PurchasePrice = request.PurchasePrice
	}
	if request.Stock != 0 {
		product.Stock = request.Stock
	}
	if request.Image != "" {
		product.Image = request.Image
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(product, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
