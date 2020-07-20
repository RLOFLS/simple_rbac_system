package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

//LoginController 用户信息控制器
type LoginController struct {
	beego.Controller
}

type singInValidate struct {
	StaffNumber string `valid:"Required"`
	Password string `valid:"Required"`
}

//Index 首页
func (c *LoginController) Index() {
	c.TplName = "staff/login.html"
}

//Signin 登录
func (c *LoginController) Signin() {
	id := c.GetString("staffNumber", "")
	password := c.GetString("password", "")

	//验证
	validate := &singInValidate{id, password}
	valid := validation.Validation{}
	b, err := valid.Valid(validate)
	if err != nil {
		logs.Error(err.Error())
		c.Data["json"] = utils.Error(utils.MsgBuzy)
		c.outputJSON()
	}
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = utils.Error(err.Message)
			break
		}
		c.outputJSON()
	}
	//查找
	staff := &models.Staff{Id: id}
	if isExist := staff.IsExistByID(); !isExist {
		c.Data["json"] = utils.Error("账户或者密码错误")
		c.outputJSON()
	}
	//密码验证
	encryptP := utils.GeneratePassword(password, staff.Salt)
	if encryptP != staff.Password {
		c.Data["json"] = utils.Error("账户或者密码错误")
		c.outputJSON()
	}
	var row int64
	row, err = staff.SaveLoginStatus(c.Ctx.Input.IP())
	if row <= 0 || err != nil {
		logs.Error(err)
		c.Data["json"] = utils.Error(utils.MsgBuzy)
		c.outputJSON()
	}
	c.setSession(*staff)
	c.Data["json"] = utils.Success("登录成功")
	c.outputJSON()
}

//Logout 退出
func (c *LoginController) Logout() {
	c.DelSession("user")
	c.Data["json"] = utils.Success("退出成功")
	c.outputJSON()
}

func (c *LoginController) setSession(staff models.Staff) {
	c.SetSession("user", staff)
}

func (c *LoginController) outputJSON() {
	c.ServeJSON()
	c.StopRun()
}
