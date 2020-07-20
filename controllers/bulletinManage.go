package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"github.com/astaxie/beego/logs"
)

// "hello2/models"
// "hello2/utils"
// "github.com/astaxie/beego/logs"

//BulletinManageController 公告管理
type BulletinManageController struct {
	BaseController
}

//Showlist 列表
func (c *BulletinManageController) Showlist() {
	c.Data["title"] = "公告管理"
	c.LayoutSections["SelfScript"] = "tpl/bulletin/showList.html"
	c.TplName = "bulletin/showList.html"
}

//Getbulletinlist 获取公共列表
func (c *BulletinManageController) Getbulletinlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Bulletin)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Showaddbulletion 添加公告
func (c *BulletinManageController) Showaddbulletion() {
	c.Data["title"] = "添加公告"
	c.LayoutSections["SelfScript"] = "tpl/bulletin/showAddBulletion.html"
	c.LayoutSections["SelfCss"] = "tpl/bulletin/showAddBulletionCss.html"
	c.TplName = "bulletin/showAddBulletion.html"
}

//Addbulletion 添加公告
func (c *BulletinManageController) Addbulletion() {
	title := strings.TrimSpace(c.GetString("title")) 
	content := strings.TrimSpace(c.GetString("content")) 
	if title == "" || content == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Bulletin{Title: title, Content: content, PostStaffId: c.loginUser.Id, Author: c.loginUser.Name}
	_, err := b.Insert()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Getbulletindetail 获取详情
func (c *BulletinManageController) Getbulletindetail() {
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

//Deletebulletion 删除公告
func (c *BulletinManageController) Deletebulletion() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Bulletin{Id: id}
	_, err := b.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("删除失败"))
}