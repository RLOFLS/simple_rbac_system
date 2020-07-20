package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type LeaveType struct {
	Id int
	Name string
	Rule string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(LeaveType))
}


//GetList  获取列表
//page 页码 pageSzie 页码大小
func (b *LeaveType) GetList(page int, pageSize int) (count int, list []*LeaveType) {
	table := orm.NewOrm()
	qs := table.QueryTable(b)
	offset := (page - 1) * pageSize
	list = make([]*LeaveType, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return 
}

//Insert 插入
func (b *LeaveType) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(b)
	return int(id), err
}


//Update 更新
func (b *LeaveType) Update(name string, rule string) (int, error) {
	o := orm.NewOrm()
	if err := o.Read(b); err != nil {
		return 0, errors.New("假期不存在")
	}
	b.Name = name
	b.Rule = rule
	row , err2 := o.Update(b)
	return int(row), err2
}

//Del 删除
func (b *LeaveType) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(b)
	return int(num), err
}

//All 全部
func (b *LeaveType) All() (list []LeaveType) {
	o := orm.NewOrm()
	_, _ = o.QueryTable(b).All(&list, "Id", "Name", "Rule")
	return
}

//IsExsit 根据id是否存在
func (b *LeaveType) IsExsit() bool {
	o := orm.NewOrm()
	if err := o.Read(b); err != nil {
		return false
	}
	return true
} 