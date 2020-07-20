package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"github.com/astaxie/beego/logs"
)

//LeaveSettingController 假期设置
type LeaveSettingController struct {
	BaseController
}

//Showlist 列表
func (c *LeaveSettingController) Showlist() {
	c.Data["title"] = "假期设置"
	c.LayoutSections["SelfScript"] = "tpl/leave/showLeaveList.html"
	c.TplName = "leave/showLeaveList.html"
}

//Getshowlist  获取列表
func (c *LeaveSettingController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.LeaveType)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Add 添加
func (c *LeaveSettingController) Add() {
	name := strings.TrimSpace(c.GetString("name")) 
	rule := strings.TrimSpace(c.GetString("rule")) 
	if name == "" || rule == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.LeaveType{Name: name, Rule: rule}
	_, err := b.Insert()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Modify 修改
func (c *LeaveSettingController) Modify() {
	id, err := c.GetInt("id")
	if err != nil {
		c.outputJSON(utils.Error("参数错误"))
	} 
	name := strings.TrimSpace(c.GetString("name")) 
	rule := strings.TrimSpace(c.GetString("rule")) 
	if name == "" || rule == "" || id <= 0{
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.LeaveType{Id: id}
	_, err = b.Update(name, rule)
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("修改成功"))
}

//Delete 删除
func (c *LeaveSettingController) Delete() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.LeaveType{Id: id}
	_, err := b.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("删除失败"))
}