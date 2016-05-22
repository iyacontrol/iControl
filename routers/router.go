package routers

import (
	"iControl/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")

	beego.Router("/realtime", &controllers.RealTimeController{}, "*:List")
	beego.Router("/realtime/getdata", &controllers.RealTimeController{}, "get:GetData")

	beego.Router("/ipset", &controllers.IpSetController{}, "*:List")
	beego.Router("/ipset/getdata", &controllers.IpSetController{}, "get:GetData")
	beego.Router("/ipset/adddata", &controllers.IpSetController{}, "post:AddData")
	beego.Router("/ipset/updatedata", &controllers.IpSetController{}, "post:UpdateData")
	beego.Router("/ipset/deldata", &controllers.IpSetController{}, "post:DelData")

	beego.Router("/protocolset", &controllers.ProtocolSetController{}, "*:List")
	beego.Router("/protocolset/getdata", &controllers.ProtocolSetController{}, "get:GetData")
	beego.Router("/protocolset/adddata", &controllers.ProtocolSetController{}, "post:AddData")
	beego.Router("/protocolset/updatedata", &controllers.ProtocolSetController{}, "post:UpdateData")
	beego.Router("/protocolset/deldata", &controllers.ProtocolSetController{}, "post:DelData")
}
