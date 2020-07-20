package controllers

import (
	"encoding/json"
	"hello2/models"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
	loginUser *models.Staff
}

const (
	//SUPERADAMINID 超级管理员
	SUPERADAMINID = "000000"
)

func (c *BaseController) Prepare() {
	user := c.GetSession("user")
	if (user == nil) {
		c.Redirect("/", http.StatusFound)
		return
	}
	user2 := user.(models.Staff)
	c.loginUser = &user2
	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Nav"] = "inc/nav.html"
    c.LayoutSections["Sidebar"] = "inc/sidebar.html"
	c.LayoutSections["Footer"] = "inc/footer.html"
	c.Data["title"] = "GO OA SYSTEM"
	c.Data["menu"] = c.getMenuTree()
	c.Data["staffId"] = c.loginUser.Id
	c.Data["staffName"] = c.loginUser.Name
	c.Data["profilePicture"] = c.loginUser.ProfilePictrue
	if c.loginUser.ProfilePictrue == "" {
		c.Data["profilePicture"] = "/static/admin/images/avatars/profile-pic.jpg"
	}
	
	c.Data["cc"], _ = c.GetControllerAndAction()
	c.Data["cc"] = strings.Replace(strings.ToLower(c.Data["cc"].(string)), "controller", "", 1) 
	logs.Emergency(c.Data["cc"])
	c.Data["noOperations"] = c.getNoOperations()
}	

func (c *BaseController) getMenuTree() []interface{} {
	menu := new(models.Menu)
	allMenus, _, _ := menu.List()
	
	if c.loginUser.Id != SUPERADAMINID {
		role := &models.Role{Id: c.loginUser.RoleId}
		roleMenu, _ := role.GetMenuPermission()
		tempMenu := make([]models.Menu, 0)
		for _, item := range allMenus {
			if _, ok := roleMenu[item.MenuCode]; ok {
				tempMenu = append(tempMenu, item)
			}
		}
		return c.getTree(tempMenu, 0)
	}
	return c.getTree(allMenus, 0)
}

//todo 权限按钮筛选
func (c *BaseController) getNoOperations() string {
	str := "{}"
	if c.loginUser.Id == SUPERADAMINID {
		return str
	}
	menuCode := c.Data["cc"]
	role := &models.Role{Id: c.loginUser.RoleId}
	roleMenu, _ := role.GetMenuPermission()
	
	roleOp, _ := roleMenu[menuCode.(string)]
	op := &models.OperateMenu{MenuCode: menuCode.(string)}
	opList, _, _ := op.GetOpByMenuCode()
	if len(opList) == 0 {
		return str
	}

	noOp := make([]string, 0)
	roleOpStr := strings.Join(roleOp, " ")
	for _, item := range opList {
		if ! strings.Contains(roleOpStr, item.(string)) {
			noOp = append(noOp, item.(string))
		}
	}

	if len(noOp) == 0 {
		return str
	}

	js, err := json.Marshal(noOp)
	if err != nil {
		logs.Error(err.Error())
		return str
	}

	return string(js)
}

func (c *BaseController) getTree(menus []models.Menu, parentID int) (m []interface{}) {
	m = make([]interface{}, 0)
	for _, menu := range menus {
		if menu.ParentId == parentID {
			mmap := menu.ToMap()
			mmap["children"] = c.getTree(menus, menu.Id)
			m = append(m, mmap)
		}
	}
	return 
}

func (c *BaseController) outputJSON(rs interface{}) {
	c.Data["json"] = rs
	c.ServeJSON()
	c.StopRun()
}
