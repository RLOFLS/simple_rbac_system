package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type LeaveApply struct {
	Id int
	LeaveTypeId int
	LeaveTypeName string
	StartTime time.Time `orm:"type(datetime)"`
	EndTime time.Time `orm:"type(datetime)"`
	StaffId string
	StaffName string
	ApplyReason string
	ApproveStaffId string
	ApproveStaffName string
	ApproveTime time.Time `orm:"auto_now_add;type(datetime)"`
	ApproveReason string
	Status string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

const (
	LAPending = "pending"
	LAReject = "reject"
	LAAgree = "agree"
	LACancel = "cancel"
)

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(LeaveApply))
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (m *LeaveApply) GetList(page int, pageSize int, id string) (count int, list []*LeaveApply) {
	table := orm.NewOrm()
	qs := table.QueryTable(m)
	if id != "" {
		qs = qs.Filter("staff_id", id)
	}
	offset := (page - 1) * pageSize
	list = make([]*LeaveApply, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return
}

//ChangeStatus 更改状态
func (m *LeaveApply) ChangeStatus() (int, error) {
	o := orm.NewOrm()
	mcp := &LeaveApply{Id: m.Id}
	if err := o.Read(mcp); err != nil {
		return 0, errors.New("申请不存在")
	}
	if mcp.Status != LAPending {
		return 0, errors.New("假期状态已发生改变，请刷新重新操作")
	}
	switch m.Status {
	case LAAgree :
		fallthrough
	case LAReject:
		if mcp.ApproveStaffId != m.ApproveStaffId {
			return 0, errors.New("审核人不是你，你无权操作")
		}
		if m.ApproveReason == "" {
			return 0, errors.New("审批理由不能为空")
		}
		mcp.ApproveReason = m.ApproveReason
		mcp.ApproveTime = time.Now()
	case LACancel:
	default:
	}
	mcp.Status = m.Status
	row, err := o.Update(mcp)
	return int(row), err
}

//Insert 插入
func (m *LeaveApply) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return int(id), err
}

//Del 删除
func (m *LeaveApply) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	return int(num), err
}
