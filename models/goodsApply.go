package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsApply struct {
	Id int
	GoodsId int
	GoodsName string
	StaffId string
	StaffName string
	ApplyReason string
	ApproveStaffId string
	ApproveStaffName string
	ApproveTime time.Time `orm:"auto_now_add;type(datetime)"`
	ApproveReason string
	Status string
	IsGet int
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

const (
	GAPending = "pending"
	GAReject = "reject"
	GAAgree = "agree"
	GACancel = "cancel"
)

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(GoodsApply))
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (m *GoodsApply) GetList(page int, pageSize int, id string) (count int, list []*GoodsApply) {
	table := orm.NewOrm()
	qs := table.QueryTable(m)
	if id != "" {
		qs = qs.Filter("staff_id", id)
	}
	offset := (page - 1) * pageSize
	list = make([]*GoodsApply, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return
}

//ChangeStatus 更改状态
func (m *GoodsApply) ChangeStatus() (int, error) {
	o := orm.NewOrm()
	mcp := &GoodsApply{Id: m.Id}
	if err := o.Read(mcp); err != nil {
		return 0, errors.New("申请不存在")
	}
	if mcp.Status != GAPending {
		return 0, errors.New("物品申请状态已发生改变，请刷新重新操作")
	}
	switch m.Status {
	case GAAgree:
		fallthrough
	case GAReject:
		if mcp.ApproveStaffId != m.ApproveStaffId {
			return 0, errors.New("审核人不是你，你无权操作")
		}
		if m.ApproveReason == "" {
			return 0, errors.New("审核理由不能为空")
		}
		mcp.ApproveReason = m.ApproveReason
		mcp.ApproveTime = time.Now()
	default:
	}
	mcp.Status = m.Status
	row, err := o.Update(mcp)
	return int(row), err
}

//Confirm 确认领取
func (m *GoodsApply) Confirm() (int, error) {
	o := orm.NewOrm()
	if err := o.Read(m); err != nil {
		return 0, errors.New("申请不存在")
	}
	if m.Status != GAAgree {
		return 0, errors.New("物品申请状态已发生改变，请刷新重新操作")
	}
	m.IsGet = 1
	row, err := o.Update(m)
	return int(row), err
}

//Insert 插入
func (m *GoodsApply) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return int(id), err
}

//Del 删除
func (m *GoodsApply) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(m)
	return int(num), err
}
