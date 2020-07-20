package controllers

import (
	"hello2/utils"
)

//JumpController 
type JumpController struct {
	BaseController
}

//Welcome 首页
func (c *JumpController) Welcome() {
	c.TplName = "welcome.html"
}

//Showprofile 个人信息
func (c *JumpController) Showprofile() {
	c.Data["title"] = "个人信息"
	c.Data["info"] = *c.loginUser
	c.TplName = "profile.html"
}

//Modifypassword 修改密码
func (c *JumpController) Modifypassword() {
	newp := c.GetString("newPassword")
	if newp == "" {
		c.outputJSON(utils.Error("密码不能为空"))
	}
	enp := utils.GeneratePassword(newp, c.loginUser.Salt)
	if c.loginUser.Password == enp {
		c.outputJSON(utils.Error("新密码与旧密码一致，请重新修改"))
	}
	row, err := c.loginUser.UpdatePassword(enp)
	if err != nil || row == 0 {
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("修改成功"))
}
