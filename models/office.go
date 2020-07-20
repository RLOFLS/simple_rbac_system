package models

import (
	"time"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Office struct {
	Id int
	SerialNumber string
	Name string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Office))
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (b *Office) GetList(page int, pageSize int) (count int, list []*Office) {
	table := orm.NewOrm()
	qs := table.QueryTable(b)
	offset := (page - 1) * pageSize
	list = make([]*Office, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return 
}

//Insert 插入
func (b *Office) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(b)
	return int(id), err
}


//Update 更新
func (b *Office) Update(name string, serialNumber string) (int, error) {
	o := orm.NewOrm()
	if err := o.Read(b); err != nil {
		return 0, errors.New("办公室不存在")
	}
	b.Name = name
	b.SerialNumber = serialNumber
	row , err2 := o.Update(b)
	return int(row), err2
}

//Del 删除
func (b *Office) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(b)
	return int(num), err
}

//All 全部
func (b *Office) All() (list []Office) {
	o := orm.NewOrm()
	_, _ = o.QueryTable(b).All(&list, "Id", "Name", "SerialNumber")
	return
}

//IsExsit 根据id是否存在
func (b *Office) IsExsit() bool {
	o := orm.NewOrm()
	if err := o.Read(b); err != nil {
		return false
	}
	return true
}