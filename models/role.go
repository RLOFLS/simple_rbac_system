package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id int
	Name string
	MenuPermission string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

//RoleMenuPermission 角色菜单
type RoleMenuPermission map[string][]string

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Role))
}

//GetMenuPermission 获取角色权限菜单
func (role *Role) GetMenuPermission() (m RoleMenuPermission, err error) {
	m = make(RoleMenuPermission, 0)
	o := orm.NewOrm()
	if err = o.Read(role); err != nil {
		return
	} 
	err = json.Unmarshal([]byte(role.MenuPermission), &m)
	return
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (role *Role) GetList(page int, pageSize int) (count int, list []*Role) {
	table := orm.NewOrm()
	qs := table.QueryTable(role)
	offset := (page - 1) * pageSize
	list = make([]*Role, 0)
	num, err := qs.Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return 
}

//Insert 插入
func (role *Role) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(role)
	return int(id), err
}

//Find 根据id查找
func (role *Role) Find() error {
	o := orm.NewOrm()
	err := o.Read(role)
	return err
}


//Update 更新
func (role *Role) Update() (int, error) {
	o := orm.NewOrm()
	mn := &Role{Id: role.Id}
	if err := o.Read(mn); err != nil {
		return 0, errors.New("角色不存在")
	}
	mn.Name = role.Name
	mn.MenuPermission = role.MenuPermission
	row , err2 := o.Update(mn)
	return int(row), err2
}

//Del 删除
func (role *Role) Del() (int, error) {
	o := orm.NewOrm()
	num, err := o.Delete(role)
	return int(num), err
}

//All 全部
func (role *Role) All() (list []Role) {
	o := orm.NewOrm()
	_, _ = o.QueryTable(role).All(&list, "Id", "Name")
	return
}

//MapAllByID  mapid
func (role *Role) MapAllByID() (l map[int]map[string]interface{}) {
	l = make(map[int]map[string]interface{}, 0)
	all := role.All()
	for _, role := range all {
		l[role.Id] = role.ToMap()
	}
	return
}

//ToMap 转成map
func (role *Role) ToMap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["id"] = role.Id
	m["name"] = role.Name
	m["menuPermission"] = role.MenuPermission
	return 
}

//IsExsit 根据id是否存在
func (role *Role) IsExsit() bool {
	o := orm.NewOrm()
	if err := o.Read(role); err != nil {
		return false
	}
	return true
}