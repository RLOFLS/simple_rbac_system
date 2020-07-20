package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"github.com/astaxie/beego/logs"
)

//GoodsSettingController 物品设置
type GoodsSettingController struct {
	BaseController
}

//Showlist 列表
func (c *GoodsSettingController) Showlist() {
	c.Data["title"] = "物品设置"
	c.LayoutSections["SelfScript"] = "tpl/goods/showGoodsList.html"
	c.TplName = "goods/showGoodsList.html"
}

//Getshowlist  获取列表
func (c *GoodsSettingController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Goods)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Add 添加
func (c *GoodsSettingController) Add() {
	name := strings.TrimSpace(c.GetString("name")) 
	description := strings.TrimSpace(c.GetString("description")) 
	if name == "" || description == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Goods{Name: name, Description: description}
	_, err := b.Insert()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Modify 修改
func (c *GoodsSettingController) Modify() {
	id, err := c.GetInt("id")
	if err != nil {
		c.outputJSON(utils.Error("参数错误"))
	} 
	name := strings.TrimSpace(c.GetString("name")) 
	description := strings.TrimSpace(c.GetString("description")) 
	if name == "" || description == "" || id <= 0{
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Goods{Id: id}
	_, err = b.Update(name, description)
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("修改成功"))
}

//Delete 删除
func (c *GoodsSettingController) Delete() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Goods{Id: id}
	_, err := b.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("删除失败"))
}