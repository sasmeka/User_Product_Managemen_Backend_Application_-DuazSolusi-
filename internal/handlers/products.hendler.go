package handlers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"
	"github.com/sasmeka/user_product_management_duazsolusi/pkg"
)

type Handler_Products struct {
	repositories.Repo_Products_IF
}

func New_Products(r repositories.Repo_Products_IF) *Handler_Products {
	return &Handler_Products{r}
}

func (h *Handler_Products) Get_Data_Products(ctx *gin.Context) {
	var product models.Products

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	orderby := ctx.Query("order_by")

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&product, page, limit, search, orderby)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Products) Post_Data_Product(ctx *gin.Context) {
	var productset models.Products

	id := ctx.GetString("userId")
	productset.Id_user = &id

	if err := ctx.Bind(&productset); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&productset)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&productset)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Products) Put_Data_Product(ctx *gin.Context) {
	var product models.Products

	product.Id_product = ctx.Param("id")
	id := ctx.GetString("userId")
	product.Id_user = &id

	count_by_id := h.Get_Count_by_Id(product.Id_product)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&product)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Update_Data(&product)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Products) Delete_Data_Product(ctx *gin.Context) {
	var product models.Products
	product.Id_product = ctx.Param("id")
	id := ctx.GetString("userId")
	product.Id_user = &id

	count_by_id := h.Get_Count_by_Id(product.Id_product)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&product)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
