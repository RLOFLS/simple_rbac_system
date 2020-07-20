package controllers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"hello2/models"
	"hello2/utils"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

//StaffManageController 员工管理
type StaffManageController struct {
	BaseController
}

//StaffValidate 员工信息验证
type StaffValidate struct {
	Id string `valid:"Required"`
	Name string `valid:"Required"`
	Gender int `valid:"Range(0, 1)"`
	Phone string `valid:"Required;Phone"`
	Address string `valid:"Required"`
	Password string `valid:"Required"`
	Email string `valid:"Required;Email"`
	RoleId int `valid:"Required"`
	Title string `valid:"Required"`
	IsApprover int `valid:"Range(0, 1)"`
}

//Showlist 列表
func (c *StaffManageController) Showlist() {
	c.Data["title"] = "员工管理"
	mapRole := new(models.Role)
	js, _ := json.Marshal(mapRole.MapAllByID())
	c.Data["roleInfo"] = string(js)
	c.LayoutSections["SelfScript"] = "tpl/staff/showList.html"
	c.TplName = "staff/showList.html"
}

//Getstafflist  获取列表
func (c *StaffManageController) Getstafflist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Staff)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Showaddstaff 添加
func (c *StaffManageController) Showaddstaff() {
	c.Data["title"] = "添加员工"
	role := new(models.Role)
	c.Data["roleInfo"] = role.All()
	c.LayoutSections["SelfScript"] = "tpl/staff/addStaff.html"
	c.LayoutSections["SelfCss"] = "tpl/permission/addRoleCss.html"
	c.TplName = "staff/addStaff.html"
}

//Addstaff 添加
func (c *StaffManageController) Addstaff() {
	id := strings.TrimSpace(c.GetString("id")) 
	name := strings.TrimSpace(c.GetString("name")) 
	password := strings.TrimSpace(c.GetString("password")) 
	gender,_ := c.GetInt("gender")
	phone := strings.TrimSpace(c.GetString("phone")) 
	address := strings.TrimSpace(c.GetString("address")) 
	email := strings.TrimSpace(c.GetString("email")) 
	title := strings.TrimSpace(c.GetString("title")) 
	roleId, _ := c.GetInt("role_id") 
	isApprover,_ := c.GetInt("is_approver", 0) 

	//验证
	validate := &StaffValidate{id, name, gender, phone, address, password, email, roleId, title, isApprover}
	valid := validation.Validation{}
	b, err := valid.Valid(validate)
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	if !b {
		for _, err := range valid.Errors {
			c.outputJSON(utils.Error(err.Message))
			break
		}
	}
	
	staff := &models.Staff{Id: id}
	if staff.IsExistByID() {
		c.outputJSON(utils.Error("员工编号已存在"))
	}
	role := &models.Role{Id: roleId}
	logs.Emergency(roleId)
	if ! role.IsExsit() {
		c.outputJSON(utils.Error("角色不存在"))
	}

	sl := make([]byte, 3)
	if _, err2 := rand.Read(sl); err2 != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(utils.MsgBuzy))
	}
	logs.Emergency(sl)
	salt := fmt.Sprintf("%x", sl)
	logs.Emergency(salt)
	enp := utils.GeneratePassword(password, salt)

	//插入
	m := &models.Staff{
		Id: id,
		Name: name,
		Gender: gender,
		Phone: phone,
		Address: address,
		Password: enp,
		Salt: salt,
		Email: email,
		IsApprover: isApprover,
		RoleId: roleId,
		Title: title,
		IsValid: 1,
	}
	_, err3 := m.Insert()
	if err3 != nil {
		logs.Error(err3.Error())
		c.outputJSON(utils.Error(err3.Error()))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Modifystaffrole 修改
func (c *StaffManageController) Modifystaffrole() {
	id := strings.TrimSpace(c.GetString("id"))
	roleId, err2 := c.GetInt("role_id")
	if err2 != nil {
		c.outputJSON(utils.Error("参数错误"))
	} 
	if roleId <= 0 || id == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	m := &models.Staff{Id: id, RoleId: roleId}
	_, err := m.Update()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("修改成功"))
}

//Delstaff 删除
func (c *StaffManageController) Delstaff() {
	id := strings.TrimSpace(c.GetString("id"))
	if id == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	m := &models.Staff{Id: id}
	_, err := m.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("删除失败"))
}

//Isfreeze 冻结
func (c *StaffManageController) Isfreeze() {
	mapfreeze := map[string]int {
		"freeze": 1,
		"unfreeze": 0,
	} 
	id := strings.TrimSpace(c.GetString("id"))
	if id == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	flag := c.GetString("type")
	v, ok := mapfreeze[flag]
	if !ok {
		c.outputJSON(utils.Error("参数错误"))
	}
	staff := &models.Staff{Id: id, IsFreeze: v}
	_, err := staff.Opfreeze()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("操作成功"))
}

//Isapprover 审核
func (c *StaffManageController) Isapprover() {
	mapfreeze := map[string]int {
		"confirm": 1,
		"cancel": 0,
	} 
	id := strings.TrimSpace(c.GetString("id"))
	if id == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	flag := c.GetString("type")
	v, ok := mapfreeze[flag]
	if !ok {
		c.outputJSON(utils.Error("参数错误"))
	}
	staff := &models.Staff{Id: id, IsApprover: v}
	_, err := staff.Opapprover()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("操作成功"))
}