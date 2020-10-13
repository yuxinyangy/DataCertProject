package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)
/*
*router.go的作用：路由功能用于接收并发接收到的浏览器的请求，用于匹配请求
 */
func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册的接口请求
    beego.Router("/user_register",&controllers.RegisterController{})
    //直接登录的页面请求接口
    beego.Router("/login.html",&controllers.LoginController{})
    //用户登录请求接口
    beego.Router("/user_login",&controllers.LoginController{})
    //用户忘记密码接口
    beego.Router("/forget_login", &controllers.ForgetController{})
    //用户登入转注册
    beego.Router("/register_register",&controllers.RegisterController{})
    //重置密码界面转登入
    beego.Router("/forget_login",&controllers.LoginController{})
    //存证
    beego.Router("/upload",&controllers.SaveProveController{})
}
