package models

import (
	"encoding/gob"
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Staff struct {
	Id string `orm:"pk"`
	Name string
	Gender int
	Phone string
	Address string
	Password string
	Salt string
	Email string
	ProfilePictrue string
	Title string
	IsValid int
	IsFreeze int
	IsApprover int
	LastLoginTime time.Time `orm:"type(datetime)"`
	LastLoginIp string
	RoleId int
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	gob.RegisterName("hello2/models.Staff", *new(Staff))
    // 需要在init中注册定义的model
	orm.RegisterModel(new(Staff)) 
}

//IsExistByID 根据id查找
func (staff *Staff) IsExistByID() bool {
	o := orm.NewOrm()
	if err := o.Read(staff); err != nil {
		return false
	}
	return true
}

//SaveLoginStatus 保存登录状态
func (staff *Staff) SaveLoginStatus(ip string) (row int64, err error) {
	o := orm.NewOrm()
	if err = o.Read(staff); err != nil {
		return
	}
	staff.LastLoginTime = time.Now()
	staff.LastLoginIp = ip
	row, err = o.Update(staff)
	return
}

//UpdatePassword 修改密码
func (staff *Staff) UpdatePassword(np string) (row int64, err error) {
	o := orm.NewOrm()
	if err = o.Read(staff); err != nil {
		return
	}
	staff.Password = np
	row, err = o.Update(staff)
	return
}

//Approves 审核人
func (staff *Staff) Approves() (list []Staff) {
	o := orm.NewOrm()
	_, _ = o.QueryTable(staff).Filter("is_approver", 1).All(&list, "Id", "Name")
	return
}

//HasApproves 根据id查看是否有审核人存在
func (staff *Staff) HasApproves() bool {
	o := orm.NewOrm()
	return o.QueryTable(staff).Filter("is_approver", 1).Exist();	
}

//GetList  获取列表
//page 页码 pageSzie 页码大小
func (staff *Staff) GetList(page int, pageSize int) (count int, list []*Staff) {
	table := orm.NewOrm()
	qs := table.QueryTable(staff)
	offset := (page - 1) * pageSize
	list = make([]*Staff, 0)
	num, err := qs.Exclude("id__exact", "000000").Filter("is_valid", 1).Limit(pageSize).Offset(offset).OrderBy("-update_time").All(&list)
	if (err != nil) {
		logs.Error(err.Error())
		return
	}
	count = int(num)
	return 
}

//Insert 插入
func (staff *Staff) Insert() (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(staff)
	return int(id), err
}

//Find 根据id查找
func (staff *Staff) Find() error {
	o := orm.NewOrm()
	err := o.Read(staff)
	return err
}

//Update 更新
func (staff *Staff) Update() (int, error) {
	o := orm.NewOrm()
	mn := &Staff{Id: staff.Id}
	if err := o.Read(mn); err != nil {
		return 0, errors.New("角色不存在")
	}
	mn.RoleId = staff.RoleId
	row , err2 := o.Update(mn)
	return int(row), err2
}

//Opfreeze 冻结
func (staff *Staff) Opfreeze() (int, error) {
	o := orm.NewOrm()
	mn := &Staff{Id: staff.Id}
	if err := o.Read(mn); err != nil {
		return 0, errors.New("角色不存在")
	}
	mn.IsFreeze = staff.IsFreeze
	row , err2 := o.Update(mn)
	return int(row), err2
}

//Opapprover 审核源
func (staff *Staff) Opapprover() (int, error) {
	o := orm.NewOrm()
	mn := &Staff{Id: staff.Id}
	if err := o.Read(mn); err != nil {
		return 0, errors.New("角色不存在")
	}
	mn.IsApprover = staff.IsApprover
	row , err2 := o.Update(mn)
	return int(row), err2
}

//Del 删除
func (staff *Staff) Del() (int, error) {
	o := orm.NewOrm()
	if err := o.Read(staff); err != nil {
		return 0, errors.New("角色不存在")
	}
	staff.IsValid = 0
	num, err := o.Update(staff)
	return int(num), err
}

//All 全部
func (staff *Staff) All() (list []Staff) {
	o := orm.NewOrm()
	_, _ = o.QueryTable(staff).All(&list, "Id", "Name", "SerialNumber")
	return
}