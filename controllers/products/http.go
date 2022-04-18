package products

import (
	"net/http"
	"rest-go-stock/businesses/products"
	"rest-go-stock/controllers"
	"rest-go-stock/controllers/products/request"
	"rest-go-stock/controllers/products/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase products.ProductUsecaseInterface
}

func NewProductController(ProductUseCase products.ProductUsecaseInterface) *ProductController {
	return &ProductController{productUseCase: ProductUseCase}
}

func (ctrl *ProductController) Add(c echo.Context) error {
	req := request.ProductRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	var data products.Domain
	data, err = ctrl.productUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainProduct(data))
}

func (ctrl *ProductController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	Product, err := ctrl.productUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainProductArray(Product))
}

func (ctrl *ProductController) GetByID(c echo.Context) error{
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := ctrl.productUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainProduct(data))
}