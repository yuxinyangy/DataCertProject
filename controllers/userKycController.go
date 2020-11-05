package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}
func (u *UserKycController) Get(){
	u.TplName = "user_kyc.html"
}

func (u *UserKycController) Post()  {
	var user models.User
	err := u.ParseForm(&user)
	fmt.Println("user:",user)
	if err != nil {
		u.Ctx.WriteString("用户认证数据解析失败，请重试")
		return
	}
	_,err =user.Update()
	if err != nil {
		u.Ctx.WriteString("用户实名认证失败，请重试")
		return
	}
	//准备数据跳转
	records,err :=models.QueryRecordByPhone(user.Phone)
	if err != nil {
		u.Ctx.WriteString("抱歉，获取认证数据失败, 请重试!")
		return
	}
	u.Data["Records"]=records
	u.Data["Phone"]=user.Phone

	u.TplName = "list_record.html"
}
