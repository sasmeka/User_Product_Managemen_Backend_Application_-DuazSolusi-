package handlers

import (
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"
	"github.com/sasmeka/user_product_management_duazsolusi/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Users struct {
	repositories.Repo_Users_IF
}

func New_Users(r repositories.Repo_Users_IF) *Handler_Users {
	return &Handler_Users{r}
}

func (h *Handler_Users) Get_Data_Users(ctx *gin.Context) {
	var user models.Users
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Users(&user, page, limit)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Users) Get_Data_Users_byId(ctx *gin.Context) {
	var user models.Users
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Users_byId(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Users) Post_Data_User(ctx *gin.Context) {
	user := models.Users{}
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	count_by_email := h.Get_Count_by_Email(user.Email)
	if count_by_email > 0 {
		pkg.Responses(400, &config.Result{Message: "e-mail already registered."}).Send(ctx)

		return
	}

	hash_pass, err_has := pkg.HashPassword(user.Pass)
	if err_has != nil {
		pkg.Responses(400, &config.Result{Message: err_has.Error()}).Send(ctx)
		return
	}
	user.Pass = hash_pass

	response, err := h.Insert_User(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Users) Put_Data_User(ctx *gin.Context) {
	var user models.Users
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	count_by_id := h.Get_Count_by_Id(user.Id_user)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)

		return
	}

	count_by_email := h.Get_Count_by_IdEmail(user.Email, user.Id_user)
	if count_by_email > 0 {
		pkg.Responses(400, &config.Result{Message: "email has been used by another user."}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	hash_pass, err_has := pkg.HashPassword(user.Pass)
	if err_has != nil {
		pkg.Responses(400, &config.Result{Message: err_has.Error()}).Send(ctx)
		return
	}
	user.Pass = hash_pass

	response, err := h.Update_User(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Users) Delete_Data_User(ctx *gin.Context) {
	var user models.Users
	user.Id_user = ctx.Param("id")

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	count_by_id := h.Get_Count_by_Id(user.Id_user)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	response1, err1 := h.Delete_User(&user)
	if err1 != nil {
		pkg.Responses(400, &config.Result{Message: err1.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, &config.Result{Message: response1}).Send(ctx)
}
