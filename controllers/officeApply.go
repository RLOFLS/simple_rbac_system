package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

//OfficeApplyController 办公室
type OfficeApplyController struct {
	BaseController
}

type officeApplyValidate struct {
	OfficeId int `valid:"Required;Min(1)"`
	OfficeName string `valid:"Required"`
	StartTime string `valid:"Required"`
	EndTime string `valid:"Required"`
	ApproveStaffId string `valid:"Required"`
	ApproveStaffName string `valid:"Required"`
	ApplyReason string `valid:"Required"`
}

//Showlist 列表
func (c *OfficeApplyController) Showlist() {
	c.Data["title"] = "办公室申请"
	c.LayoutSections["SelfScript"] = "tpl/office/showList.html"
	c.TplName = "office/showList.html"
}

//Getshowlist  获取列表
func (c *OfficeApplyController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.OfficeApply)
	count, list := m.GetList(page, pageSize, c.loginUser.Id)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Cancel 取消
func (c *OfficeApplyController) Cancel() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.OfficeApply{Id: id, Status: models.OACancel}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("取消成功"))
}

//Showadd 添加
func (c *OfficeApplyController) Showadd() {
	c.Data["title"] = "假期申请"
	c.Data["selector"] = new(models.Office).All()
	c.Data["approves"] = new(models.Staff).Approves()
	c.LayoutSections["SelfScript"] = "tpl/office/applyOffice.html"
	c.LayoutSections["SelfCss"] = "tpl/leave/applyLeaveCss.html"
	c.TplName = "office/applyOffice.html"
}

//Apply 添加
func (c *OfficeApplyController) Apply() {
	officeId, _ := c.GetInt("office_id", 0)
	officeName := strings.TrimSpace(c.GetString("office_name")) 
	startTime := strings.TrimSpace(c.GetString("start_time")) 
	endTime := strings.TrimSpace(c.GetString("end_time")) 
	approveStaffId := strings.TrimSpace(c.GetString("approve_staff_id")) 
	approveStaffName := strings.TrimSpace(c.GetString("approve_staff_name")) 
	applyReason := strings.TrimSpace(c.GetString("apply_reason")) 
	
	//验证
	validate := &leaveApplyValidate{officeId, officeName, startTime, endTime, approveStaffId, approveStaffName, applyReason}
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

	office := &models.Office{Id: officeId}
	if ! office.IsExsit() {
		c.outputJSON(utils.Error("办公室不存在"))
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

	officeApply := &models.OfficeApply{OfficeId: officeId}
	if err4 := officeApply.ApplyTimeCheck(officeId, startTime, endTime); err4 != nil {
		c.outputJSON(utils.Error(err4.Error()))
	}
	//插入
	officeApply = &models.OfficeApply{
		OfficeId: officeId,
		OfficeName: officeName,
		StartTime: st,
		EndTime: et,
		StaffId: c.loginUser.Id,
		StaffName: c.loginUser.Name,
		ApplyReason: applyReason,
		ApproveStaffId: approveStaffId,
		ApproveStaffName: approveStaffName,
		Status: models.OAPending,
	}
	if _, err := officeApply.Insert(); err != nil {
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("申请成功"))

}