package controllers

import (
	//"github.com/astaxie/beego"
)

//UserInfoController 用户信息控制器
type UserInfoController struct {
	BaseController
}

//Get get请求
func (c *UserInfoController) Get() {
	c.Ctx.WriteString("userInfo controller get !!!")
	user := c.GetSession("name")
	if user == nil {
		c.Ctx.WriteString("userInfo controller get empty !!!")
	}
	c.Ctx.WriteString(user.(string))
}

func (c *UserInfoController) Set() {
	c.Ctx.WriteString("userInfo controller set !!!")
	c.SetSession("name", "tym")
}

//GetName 获取名字
func (c *UserInfoController) GetName() {
	type S struct {
		Name string
	}
	data := &S{"tttt"}
	c.Data["json"] = data
	c.ServeJSON()
}
