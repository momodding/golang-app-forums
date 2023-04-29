package controller

import (
	"errors"
	"forum-app/entity"
	"forum-app/helper"
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/bytebufferpool"
	"strconv"
	"time"
)

type AuthController interface {
	LoginView(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	oauthService service.OauthService
}

func NewAuthController(oauthService service.OauthService) *AuthControllerImpl {
	return &AuthControllerImpl{oauthService: oauthService}
}

func (ctrl *AuthControllerImpl) LoginView(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"Title": "Login",
	}, "layouts/main")
}

func (ctrl *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	clientVal := ctx.Locals("client")
	if clientVal == nil {
		return errors.New("invalid ClientID")
	}
	client, ok := clientVal.(*entity.OauthClient)
	if !ok {
		return nil
	}

	user, err := ctrl.oauthService.AuthUser(ctx.FormValue("username"), ctx.FormValue("password"))
	helper.PanicIfError(err)

	scope, err := ctrl.oauthService.GetScope(ctx.FormValue("scope"))
	helper.PanicIfError(err)

	accessToken, refreshToken, err := ctrl.oauthService.GetToken(client, user, scope)

	cookies := map[string]string{
		"ClientID":     strconv.FormatUint(client.ID, 10),
		"Username":     user.Username,
		"AccessToken":  accessToken.Token,
		"RefreshToken": refreshToken.Token,
	}

	for key, element := range cookies {
		cookie := new(fiber.Cookie)
		cookie.Name = key
		cookie.Value = element
		cookie.Expires = time.Now().Add(24 * time.Hour)
		ctx.Cookie(cookie)
	}

	redirectUri := ctx.Query("redirectUrl")
	if redirectUri == "" {
		redirectUri = "/web/admin"
	}

	bb := bytebufferpool.Get()
	defer bytebufferpool.Put(bb)
	_, err = bb.WriteString(redirectUri + "?")
	helper.PanicIfError(err)
	_, err = bb.Write(ctx.Context().Request.URI().QueryString())
	helper.PanicIfError(err)
	err = ctx.Redirect(bb.String())
	helper.PanicIfError(err)

	return nil
}
