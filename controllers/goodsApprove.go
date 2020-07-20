package controllers

import (
	"hello2/models"
	"hello2/utils"

	"github.com/astaxie/beego/logs"
)

//GoodsApproveController 物品
type GoodsApproveController struct {
	BaseController
}

//Showlist 列表
func (c *GoodsApproveController) Showlist() {
	c.Data["title"] = "假期审批"
	c.LayoutSections["SelfScript"] = "tpl/goods/showApproveList.html"
	c.TplName = "goods/showApproveList.html"
}

//Getshowlist  获取列表
func (c *GoodsApproveController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.GoodsApply)
	count, list := m.GetList(page, pageSize, "")
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Agree 同意
func (c *GoodsApproveController) Agree() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.GoodsApply{Id: id, ApproveStaffId: c.loginUser.Id, ApproveReason: "/", Status: models.GAAgree}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("同意成功"))
}

//Reject 拒绝
func (c *GoodsApproveController) Reject() {
	id, err := c.GetInt("id", 0)
	approveReason := c.GetString("approve_reason")
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.GoodsApply{Id: id, ApproveStaffId: c.loginUser.Id, ApproveReason: approveReason, Status: models.GAReject}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("拒绝成功"))
}

//Confirm 确认领取
func (c *GoodsApproveController) Confirm() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.GoodsApply{Id: id}
	_, err = m.Confirm()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success(utils.MsgSuccess))
}