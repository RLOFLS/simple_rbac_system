package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego/logs"
)

//HomeController 首页
type HomeController struct {
	BaseController
}

//Index 首页
func (c *HomeController) Index() {
	c.LayoutSections["SelfScript"] = "tpl/index.html"
	c.TplName = "index.html"
}

//Getbulletinlist 获取公告
func (c *HomeController) Getbulletinlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Bulletin)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Getbulletindetail 获取详情
func (c *HomeController) Getbulletindetail () {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Bulletin{Id: id}
	err := b.Detail()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Item(b))
}