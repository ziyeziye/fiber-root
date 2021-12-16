package login

import (
	"fiber-root/app/api"
	"fiber-root/app/model"
	"fiber-root/app/service"
	"fiber-root/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

// Login 登录
// @Summary 登录
// @Description 登录
// @Tags login
// @Accept json
// @Produce json
// @Param user body model.User true "login user"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Failure 400 {object} api.ResponseHTTP{}
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return api.Response(c, err, nil)
	}
	err := service.UserService.Login(user)
	if err != nil {
		return api.Response(c, err, nil)
	}

	t, err := jwt.GenerateToken(user.ID, user.Name, user.Password)

	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, nil, t)
}
