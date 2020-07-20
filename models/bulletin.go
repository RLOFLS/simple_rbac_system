package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Bulletin struct {
	Id int
	Title string
	Content string
	PostStaffId string
	Author string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Bulletin))
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (b *Bulletin) GetList(page int, pageSize int) (count int, list []*Bulletin) {
	table := orm.NewOrm()
	qs := table.QueryTable(b)
	offset := (page - 1) * pageSize
	list = make([]*Bulletin, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return 
}

//Detail 根据id获取详情
func (b *Bulletin) Detail() (err error) {
	o := orm.NewOrm()
	err = o.Read(b)
	return
}

//Insert 插入
func (b *Bulletin) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(b)
	return int(id), err
}

//Del 删除
func (b *Bulletin) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(b)
	return int(num), err
}
