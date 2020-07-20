package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type OfficeApply struct {
	Id int
	OfficeId int
	OfficeName string
	StartTime time.Time `orm:"type(datetime)"`
	EndTime time.Time `orm:"type(datetime)"`
	StaffId string
	StaffName string
	ApplyReason string
	ApproveStaffId string
	ApproveStaffName string
	ApproveTime time.Time `orm:"type(datetime)"`
	ApproveReason string
	Status string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

const (
	//OAPending //以预约
	OAPending = "pending"
	//OAClose 关闭请求
	OAClose = "close"
	//OACancel 取消
	OACancel = "cancel"
)

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(OfficeApply))
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (m *OfficeApply) GetList(page int, pageSize int, id string) (count int, list []*OfficeApply) {
	table := orm.NewOrm()
	qs := table.QueryTable(m)
	if id != "" {
		qs = qs.Filter("staff_id", id)
	}
	offset := (page - 1) * pageSize
	list = make([]*OfficeApply, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return
}

//ChangeStatus 更改状态
func (m *OfficeApply) ChangeStatus() (int, error) {
	o := orm.NewOrm()
	mcp := &OfficeApply{Id: m.Id}
	if err := o.Read(mcp); err != nil {
		return 0, errors.New("申请不存在")
	}
	if mcp.Status != OAPending {
		return 0, errors.New("办公室状态已发生改变，请刷新重新操作")
	}
	switch m.Status {
	case OAClose:
		if mcp.ApproveStaffId != m.ApproveStaffId {
			return 0, errors.New("审核人不是你，你无权操作")
		}
		if m.ApproveReason == "" {
			return 0, errors.New("关闭理由不能为空")
		}
		mcp.ApproveReason = m.ApproveReason
		mcp.ApproveTime = time.Now()
	case OACancel:
	default:
	}
	mcp.Status = m.Status
	row, err := o.Update(mcp)
	return int(row), err
}

//Insert 插入
func (m *OfficeApply) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return int(id), err
}

//Del 删除
func (m *OfficeApply) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	return int(num), err
}

//ApplyTimeCheck 申请时间检测
func (m *OfficeApply) ApplyTimeCheck(officeID int, st string, et string) error {
	o := orm.NewOrm()
	q := o.QueryTable(m)
	if q.Filter("office_id", officeID).Filter("start_time__gt", st).Filter("start_time__lt", et).Exist() {
		return fmt.Errorf("时间冲突，时间段：%s~%s已被预约", st, et)
	}
	if q.Filter("office_id", officeID).Filter("end_time__gt", st).Filter("end_time__lt", et).Exist() {
		return fmt.Errorf("时间冲突，时间段：%s~%s已被预约", st, et)
	}
	if q.Filter("office_id", officeID).Filter("start_time__lt", st).Filter("end_time__gt", st).Exist() {
		return fmt.Errorf("时间冲突，时间段：%s~%s已被预约", st, et)
	}
	if q.Filter("office_id", officeID).Filter("start_time__lt", et).Filter("end_time__gt", et).Exist() {
		return fmt.Errorf("时间冲突，时间段：%s~%s已被预约", st, et)
	}
	return nil
}
