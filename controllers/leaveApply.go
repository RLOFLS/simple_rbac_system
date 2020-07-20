package controllers

import (
	"encoding/json"
	"hello2/models"
	"hello2/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

//LeaveApplyController 假期
type LeaveApplyController struct {
	BaseController
}

type leaveApplyValidate struct {
	LeaveTypeId int `valid:"Required;Min(1)"`
	LeaveTypeName string `valid:"Required"`
	StartTime string `valid:"Required"`
	EndTime string `valid:"Required"`
	ApproveStaffId string `valid:"Required"`
	ApproveStaffName string `valid:"Required"`
	ApplyReason string `valid:"Required"`
}

//Showlist 列表
func (c *LeaveApplyController) Showlist() {
	c.Data["title"] = "假期申请"
	c.LayoutSections["SelfScript"] = "tpl/leave/showList.html"
	c.TplName = "leave/showList.html"
}

//Getshowlist  获取列表
func (c *LeaveApplyController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.LeaveApply)
	count, list := m.GetList(page, pageSize, c.loginUser.Id)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Cancel 取消
func (c *LeaveApplyController) Cancel() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.LeaveApply{Id: id, Status: models.LACancel}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("取消成功"))
}

//Showadd 添加
func (c *LeaveApplyController) Showadd() {
	c.Data["title"] = "假期申请"
	c.Data["leaves"] = new(models.LeaveType).All()
	c.Data["approves"] = new(models.Staff).Approves()
	j, _ := json.Marshal(c.Data["leaves"])
	c.Data["leavesJson"] = string(j)
	c.LayoutSections["SelfScript"] = "tpl/leave/applyLeave.html"
	c.LayoutSections["SelfCss"] = "tpl/leave/applyLeaveCss.html"
	c.TplName = "leave/applyLeave.html"
}

//Add 添加
func (c *LeaveApplyController) Apply() {
	leaveTypeId, _ := c.GetInt("leave_type_id", 0)
	leaveTypeName := strings.TrimSpace(c.GetString("leave_type_name")) 
	startTime := strings.TrimSpace(c.GetString("start_time")) 
	endTime := strings.TrimSpace(c.GetString("end_time")) 
	approveStaffId := strings.TrimSpace(c.GetString("approve_staff_id")) 
	approveStaffName := strings.TrimSpace(c.GetString("approve_staff_name")) 
	applyReason := strings.TrimSpace(c.GetString("apply_reason")) 
	
	//验证
	validate := &leaveApplyValidate{leaveTypeId, leaveTypeName, startTime, endTime, approveStaffId, approveStaffName, applyReason}
	valid := validation.Validation{}
	b, err := valid.Valid(validate)
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	if !b {
		for _, err := range valid.Errors {
			c.outputJSON(utils.Error(err.Message))
			break
		}
		
	}

	leavetype := &models.LeaveType{Id: leaveTypeId}
	if ! leavetype.IsExsit() {
		c.outputJSON(utils.Error("假期不存在"))
	}

	approver := &models.Staff{Id: approveStaffId}
	if ! approver.HasApproves() {
		c.outputJSON(utils.Error("审核人不对"))
	}
	
	loc, _ := time.LoadLocation("Asia/Shanghai")
	st, err2 :=time.ParseInLocation("2006-01-02 15:04:05", startTime, loc)
	if err2 != nil {
		c.outputJSON(utils.Error("开始时间格式错误"))
	}
	et, err3 :=time.ParseInLocation("2006-01-02 15:04:05", endTime, loc)
	if err3 != nil {
		c.outputJSON(utils.Error("开始时间格式错误"))
	}

	if st.Before(time.Now()) {
		c.outputJSON(utils.Error("开始时间必须大于当前时间"))
	}

	if st.After(et) {
		c.outputJSON(utils.Error("结束时间必须大于开始时间"))
	}
	//插入
	leaveApply := &models.LeaveApply{
		LeaveTypeId: leaveTypeId,
		LeaveTypeName: leaveTypeName,
		StartTime: st,
		EndTime: et,
		StaffId: c.loginUser.Id,
		StaffName: c.loginUser.Name,
		ApplyReason: applyReason,
		ApproveStaffId: approveStaffId,
		ApproveStaffName: approveStaffName,
		Status: models.LAPending,
	}
	if _, err := leaveApply.Insert(); err != nil {
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("申请成功"))

}