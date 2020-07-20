package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego/logs"
)

//OfficeApproveController 办公室
type OfficeApproveController struct {
	BaseController
}

//Showlist 列表
func (c *OfficeApproveController) Showlist() {
	c.Data["title"] = "办公室审批"
	c.LayoutSections["SelfScript"] = "tpl/office/showApproveList.html"
	c.TplName = "office/showApproveList.html"
}

//Getshowlist  获取列表
func (c *OfficeApproveController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.OfficeApply)
	count, list := m.GetList(page, pageSize, "")
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Close 关闭
func (c *OfficeApproveController) Close() {
	id, err := c.GetInt("id", 0)
	approveReason := c.GetString("approve_reason")
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.OfficeApply{Id: id, ApproveStaffId: c.loginUser.Id, ApproveReason: approveReason, Status: models.OAClose}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("关闭成功"))
}
