package routers

import (
	"hello2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{}, "get:Index")
    beego.Router("/userInfo", &controllers.UserInfoController{})
	beego.Router("/userInfo/name", &controllers.UserInfoController{}, "*:GetName")
    beego.Router("/office/list", &controllers.OfficeController{}, "Get:GetList")
	ns := beego.NewNamespace("/admin",
		beego.NSAutoRouter(&controllers.LoginController{}),
		beego.NSAutoRouter(&controllers.JumpController{}),
		beego.NSAutoRouter(&controllers.HomeController{}),
		beego.NSAutoRouter(&controllers.BulletinManageController{}),
		beego.NSAutoRouter(&controllers.LeaveSettingController{}),
		beego.NSAutoRouter(&controllers.OfficeSettingController{}),
		beego.NSAutoRouter(&controllers.GoodsSettingController{}),
		beego.NSAutoRouter(&controllers.LeaveApplyController{}),
		beego.NSAutoRouter(&controllers.OfficeApplyController{}),
		beego.NSAutoRouter(&controllers.GoodsApplyController{}),
		beego.NSAutoRouter(&controllers.LeaveApproveController{}),
		beego.NSAutoRouter(&controllers.OfficeApproveController{}),
		beego.NSAutoRouter(&controllers.GoodsApproveController{}),
		beego.NSAutoRouter(&controllers.PermissionManageController{}),
		beego.NSAutoRouter(&controllers.StaffManageController{}),
	)
	beego.AddNamespace(ns)
}
