package controllers

import (
	"hello2/models"
	"hello2/utils"
	"strings"
	"github.com/astaxie/beego/logs"
)

//PermissionManageController 权限
type PermissionManageController struct {
	BaseController
}

//Showrolelist 列表
func (c *PermissionManageController) Showrolelist() {
	c.Data["title"] = "权限管理"
	c.LayoutSections["SelfScript"] = "tpl/permission/showRoleList.html"
	c.TplName = "permission/showRoleList.html"
}

//Getshowlist  获取列表
func (c *PermissionManageController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	b := new(models.Role)
	count, list := b.GetList(page, pageSize)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Showaddrole 添加角色
func (c *PermissionManageController) Showaddrole() {
	c.Data["title"] = "添加角色"
	c.Data["htmlMenu"] = c.getPessionTree()
	//logs.Emergency(c.Data["hemlMenu"])
	c.LayoutSections["SelfScript"] = "tpl/permission/addRole.html"
	c.LayoutSections["SelfCss"] = "tpl/permission/addRoleCss.html"
	c.TplName = "permission/addRole.html"
}

//Addrole 添加
func (c *PermissionManageController) Addrole() {
	name := strings.TrimSpace(c.GetString("name")) 
	menuPermission := strings.TrimSpace(c.GetString("menu_permission")) 
	if name == "" || menuPermission == "" {
		c.outputJSON(utils.Error("参数错误"))
	}
	b := &models.Role{Name: name, MenuPermission: menuPermission}
	_, err := b.Insert()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("添加成功"))
}

//Showmodifyrole 角色修改
func (c *PermissionManageController) Showmodifyrole() {
	id, _ := c.GetInt("roleId")
	role := &models.Role{Id: id}
	err := role.Find()
	c.Data["title"] = "添加角色"
	if err != nil {
		c.Redirect("/notFound",404)
		c.StopRun()
	}
	c.Data["roleName"] = role.Name
	c.Data["roleId"] = role.Id
	c.Data["roleMenu"] = role.MenuPermission
	c.Data["htmlMenu"] = c.getPessionTree()
	c.LayoutSections["SelfScript"] = "tpl/permission/modifyRole.html"
	c.LayoutSections["SelfCss"] = "tpl/permission/addRoleCss.html"
	c.TplName = "permission/modifyRole.html"
}

//Modifyrole 修改
func (c *PermissionManageController) Modifyrole() {
	id, err := c.GetInt("roleId")
	if err != nil {
		c.outputJSON(utils.Error("参数错误"))
	} 
	name := strings.TrimSpace(c.GetString("name")) 
	menuPermission := strings.TrimSpace(c.GetString("menu_permission")) 
	if name == "" || menuPermission == "" || id <= 0{
		c.outputJSON(utils.Error("参数错误"))
	}
	m := &models.Role{Id: id, Name: name, MenuPermission: menuPermission}
	_, err = m.Update()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("修改成功"))
}

//Delrole 删除
func (c *PermissionManageController) Delrole() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	}
	m := &models.Role{Id: id}
	_, err := m.Del()
	if err != nil {
		logs.Error(err)
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("删除失败"))
}


func (c *PermissionManageController) getPessionTree() string {
	menu := new(models.Menu)
	allMenu, _, _ := menu.List()
	operate := new(models.OperateMenu)
	operateMenu := operate.GetOperateMenu()
	mapMenu := make([]map[string]interface{}, 0)
	for _, item := range allMenu {
		temp := item.ToMap()
		if ops, ok := operateMenu[item.MenuCode]; ok {
			temp["operateChildren"] = ops
		}
		mapMenu = append(mapMenu, temp)
	}
	tree := c.getMenuTree(mapMenu, 0)
	return c.getTreeHTML(tree)
}

func (c *PermissionManageController) getMenuTree(menus []map[string]interface{}, parentID int) (m []map[string]interface{}) {
	m = make([]map[string]interface{}, 0)
	for _, menu := range menus {
		if menu["parentId"] == parentID {
			menu["children"] = c.getMenuTree(menus, menu["id"].(int))
			m = append(m, menu)
		}
	}
	return 
}

func(c *PermissionManageController) getTreeHTML(menus []map[string]interface{}) string {
	html := ""
	for _, item := range menus {
		ch , ok1 := item["children"]
		ops , ok2 := item["operateChildren"]
		if (ok1 && len(ch.([]map[string]interface{})) > 0) || ok2 {
			html += `<li class="tree-branch" role="treeitem" aria-expanded="false" data-code="` + item["menuCode"].(string) + `">`;
			html += `<i class="icon-caret fa fa-plus-minus-o"></i>&nbsp;`
			html += `<div class="tree-branch-header">`
			html += `<span class="tree-branch-name">`
			html += `<i class="icon-folder ace-icon tree-plus"></i>`
			html += `<span class="tree-label">` + item["name"].(string) + `</span></span>`
			html += `</div>`
			html += `<ul class="tree-branch-children hidden" role="group">`
			var children interface{}
			if len(ch.([]map[string]interface{})) > 0 {
				children = ch
			} else {
				children = ops
			}
			html = html + c.getTreeHTML(children.([]map[string]interface{}));
			html += `</ul>`
			html += `</li>`
		} else {
			html += `<li class="tree-item" role="treeitem" data-code="`
			if _ , ok := item["operateCode"]; ok {
				html += item["menuCode"].(string) + "-" + item["operateCode"].(string)
			} else {
				html += item["menuCode"].(string)
			}
			html += `">`;  
			html += `<span class="tree-item-name">`
			html += `<i class="icon-item ace-icon fa fa-check"></i>`
			html += `<span class="tree-label">`
			if _, ok := item["operateName"]; ok {
				html += item["operateName"].(string)
			} else {
				html += item["name"].(string)
			}
			html += `</span></span>`
			html += `</li>`
		}
	}
	return html;
}