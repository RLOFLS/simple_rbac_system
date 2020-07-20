package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego/logs"
)

//LeaveApproveController 假期
type LeaveApproveController struct {
	BaseController
}

//Showlist 列表
func (c *LeaveApproveController) Showlist() {
	c.Data["title"] = "假期审批"
	c.LayoutSections["SelfScript"] = "tpl/leave/showApproveList.html"
	c.TplName = "leave/showApproveList.html"
}

//Getshowlist  获取列表
func (c *LeaveApproveController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.LeaveApply)
	count, list := m.GetList(page, pageSize, "")
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Agree 同意
func (c *LeaveApproveController) Agree() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.LeaveApply{Id: id, ApproveStaffId: c.loginUser.Id, ApproveReason: "/", Status: models.LAAgree}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("同意成功"))
}

//Reject 拒绝
func (c *LeaveApproveController) Reject() {
	id, err := c.GetInt("id", 0)
	approveReason := c.GetString("approve_reason")
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.LeaveApply{Id: id, ApproveStaffId: c.loginUser.Id, ApproveReason: approveReason, Status: models.LAReject}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("拒绝成功"))
}