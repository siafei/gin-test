package controller

import (
	"github.com/gin-gonic/gin"
	"github/siafei/gin-test/pkg/response"
)

type UserController struct {
}

func (u UserController) GetUsers(c *gin.Context) {

	//param := struct {
	//	Name  string `form:"name" binding:"max=100"`
	//	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	//}{}
	res := response.NewResponse(c)
	res.ToResponse("xixi")
	//valid, errs := validError.BindAndValid(c, &param)
	//if !valid {
	//	global.Logger.ErrorOf("app.BindAndValid errs: %v", errs)
	//	res.ToErrorResponse(1,errs.Error())
	//	return
	//}
	//var users []model.UserModel
	//global.Logger.Error("test-log")
	//if err := global.DB.Find(&users).Error; err != nil {
	//	res.ToErrorResponse(1,errs.Error())
	//	return
	//}
	return
}
