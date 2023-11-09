package handlers

import (
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"
	"github.com/sasmeka/user_product_management_duazsolusi/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Auth struct {
	repositories.Repo_AuthIF
}

func New_Auth(r repositories.Repo_AuthIF) *Handler_Auth {
	return &Handler_Auth{r}
}

func (h *Handler_Auth) Login(ctx *gin.Context) {
	var user models.Auth
	var err_val error
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_User(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(response.Pass, user.Pass); err != nil {
		pkg.Responses(400, &config.Result{Message: "wrong password"}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(response.Id_user, response.Role, response.Email)
	tokens, err := jwtt.Generate()
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, &config.Result{Token: tokens}).Send(ctx)
}

func (h *Handler_Auth) Register(ctx *gin.Context) {
	var user models.Users
	var err_val error
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

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

	response, err := h.Register_rep(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
