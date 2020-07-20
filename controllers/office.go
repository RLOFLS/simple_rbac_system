package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//OfficeController 控制器
type OfficeController struct {
	beego.Controller
}

//GetList 获取office 列表
func (c *OfficeController) GetList() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 10)

	office := new(models.Office)
	count, list := office.GetList(page, pageSize)

	c.Data["json"] = utils.Paginate(list, count, page, pageSize, "")
	logs.Emergency(c.Data["json"])
	c.ServeJSON()
}