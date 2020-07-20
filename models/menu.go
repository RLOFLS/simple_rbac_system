package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id int
	Name string
	ParentId int
	MenuCode string
	Url string
	Icon string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Menu))
}

//ToMap 转成map
func (menu *Menu) ToMap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["id"] = menu.Id
	m["name"] = menu.Name
	m["parentId"] = menu.ParentId
	m["menuCode"] = menu.MenuCode
	m["url"] = menu.Url
	m["icon"] = menu.Icon
	return 
}

//List 获取全部菜单
func (menu *Menu) List() (menus []Menu, count int64, err error) {
	o := orm.NewOrm()
	q := o.QueryTable(menu)
	count, err = q.All(&menus)
	return
}