package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"github.com/astaxie/beego/logs"
)

//OfficeSettingController 设置
type OfficeSettingController struct {
	BaseController
}

//Showlist 列表
func (c *OfficeSettingController) Showlist() {
	c.Data["title"] = "办公室设置"
	c.LayoutSections["SelfScript"] = "tpl/office/showOfficeList.html"
	c.TplName = "office/showOfficeList.html"
}

//Getshowlist  获取列表
func (c *OfficeSettingController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Office)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Add 添加
func (c *OfficeSettingController) Add() {
	name := strings.TrimSpace(c.GetString("name")) 
	serialNumber := strings.TrimSpace(c.GetString("serial_number")) 
	if name == "" || serialNumber == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Office{Name: name, SerialNumber: serialNumber}
	_, err := b.Insert()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Modify 修改
func (c *OfficeSettingController) Modify() {
	id, err := c.GetInt("id")
	if err != nil {
		c.outputJSON(utils.Error("参数错误"))
	} 
	name := strings.TrimSpace(c.GetString("name")) 
	serialNumber := strings.TrimSpace(c.GetString("serial_number")) 
	if name == "" || serialNumber == "" || id <= 0{
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Office{Id: id}
	_, err = b.Update(name, serialNumber)
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("修改成功"))
}

//Delete 删除
func (c *OfficeSettingController) Delete() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Office{Id: id}
	_, err := b.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	c.outputJSON(utils.Success("删除失败"))
}