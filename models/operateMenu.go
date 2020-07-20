package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type OperateMenu struct {
	Id int
	OperateCode string
	OperateName string
	MenuCode string
	Methods string
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(OperateMenu))
}

func (m *OperateMenu) toMap() map[string]interface{} {
	return map[string]interface{}{
		"id" : m.Id,
		"operateCode" : m.OperateCode,
		"operateName" : m.OperateName,
		"menuCode" : m.MenuCode,
		"methods" : m.Methods,
	}
}

//GetOperateMenu 获取权限操作
func (m *OperateMenu) GetOperateMenu() (omenu map[string][]map[string]interface{}) {
	o := orm.NewOrm()
	list := make([]OperateMenu, 0)
	_, _ = o.QueryTable(m).All(&list)
	omenu = make(map[string][]map[string]interface{}, 0)
	for _, item := range list { 
		if _, ok := omenu[item.MenuCode]; ok {
			omenu[item.MenuCode] = append(omenu[item.MenuCode], item.toMap())
		} else {
			omenu[item.MenuCode] = []map[string]interface{}{item.toMap()}
		}
	}
	return
}

//GetOpByMenuCode 获取操作
func (m *OperateMenu) GetOpByMenuCode() (l orm.ParamsList, row int64, err error){
	l = make(orm.ParamsList, 0)
	o := orm.NewOrm()
	row, err = o.QueryTable(m).Filter("menu_code", m.MenuCode).ValuesFlat(&l, "operate_code")
	return 
}