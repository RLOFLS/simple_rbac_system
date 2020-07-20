package controllers

import (
	"encoding/json"
	"hello2/models"
	"hello2/utils"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

//GoodsApplyController 物品
type GoodsApplyController struct {
	BaseController
}

type goodsApplyValidate struct {
	GoodsId int `valid:"Required;Min(1)"`
	GoodsName string `valid:"Required"`
	ApproveStaffId string `valid:"Required"`
	ApproveStaffName string `valid:"Required"`
	ApplyReason string `valid:"Required"`
}

//Showlist 列表
func (c *GoodsApplyController) Showlist() {
	c.Data["title"] = "办公室申请"
	c.LayoutSections["SelfScript"] = "tpl/goods/showList.html"
	c.TplName = "goods/showList.html"
}

//Getshowlist  获取列表
func (c *GoodsApplyController) Getshowlist() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 10
	m := new(models.GoodsApply)
	count, list := m.GetList(page, pageSize, c.loginUser.Id)
	c.outputJSON(utils.Paginate(list, count, page, pageSize, ""))
}

//Cancel 取消
func (c *GoodsApplyController) Cancel() {
	id, err := c.GetInt("id", 0)
	if err != nil || id <= 0 {
		c.outputJSON(utils.Error("参数错误"))
	} 
	m := &models.GoodsApply{Id: id, Status: models.GACancel}
	_, err = m.ChangeStatus()
	if err != nil {
		logs.Error(err.Error())
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("取消成功"))
}

//Showadd 添加
func (c *GoodsApplyController) Showadd() {
	c.Data["title"] = "假期申请"
	c.Data["selector"] = new(models.Goods).All()
	c.Data["approves"] = new(models.Staff).Approves()
	j, _ := json.Marshal(c.Data["selector"])
	c.Data["selectorJson"] = string(j)
	c.LayoutSections["SelfScript"] = "tpl/goods/applyGoods.html"
	c.LayoutSections["SelfCss"] = "tpl/leave/applyLeaveCss.html"
	c.TplName = "goods/applyGoods.html"
}

//Apply 添加
func (c *GoodsApplyController) Apply() {
	goodsId, _ := c.GetInt("goods_id", 0)
	goodsName := strings.TrimSpace(c.GetString("goods_name")) 
	approveStaffId := strings.TrimSpace(c.GetString("approve_staff_id")) 
	approveStaffName := strings.TrimSpace(c.GetString("approve_staff_name")) 
	applyReason := strings.TrimSpace(c.GetString("apply_reason")) 
	
	//验证
	validate := &goodsApplyValidate{goodsId, goodsName, approveStaffId, approveStaffName, applyReason}
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

	goods := &models.Goods{Id: goodsId}
	if ! goods.IsExsit() {
		c.outputJSON(utils.Error("物品不存在"))
	}
	approver := &models.Staff{Id: approveStaffId}
	if ! approver.HasApproves() {
		c.outputJSON(utils.Error("审核人不对"))
	}
	//插入
	goodsApply := &models.GoodsApply{
		GoodsId: goodsId,
		GoodsName: goodsName,
		StaffId: c.loginUser.Id,
		StaffName: c.loginUser.Name,
		ApplyReason: applyReason,
		ApproveStaffId: approveStaffId,
		ApproveStaffName: approveStaffName,
		Status: models.GAPending,
	}
	if _, err := goodsApply.Insert(); err != nil {
		c.outputJSON(utils.Error(err.Error()))
	}
	c.outputJSON(utils.Success("申请成功"))

}