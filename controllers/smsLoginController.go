package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type SmsLoginController struct {
	beego.Controller
}

/**
 *该方法用于处理浏览器中的跳转手机验证码登入页面
 */
func (s *SmsLoginController) Get() {
	s.TplName = "login_sms.html"
}

/**
 *该方法用于处理post请求：使用手机号和验证码进行登入
 */
func (s *SmsLoginController) Post() {
	//1.解析数据
	var sms models.SmsRecord
	if err := s.ParseForm(&sms); err != nil {
		s.Ctx.WriteString("解析数据失败，请重试!")
		return
	}
	//1.1使用用户提交的手机号进行用户表查询,判断用户是否已注册
	user,err :=models.QueryUserByPhone(sms.Phone)
	if err != nil {//遇到error,说明没查到数据,说明该用户还未注册
		s.Ctx.WriteString("登入失败,请重试")
		return
	}
	//2.将解析到的数据作为条件进行数据库查询
	codeRecord,err :=sms.QuerySmsByBizId()
	//3.判断数据库查询结果
	//a.查询错误或未查到数据：返回提示信息
	if err != nil {
		s.Ctx.WriteString("手机号或验证码错误,请重试!")
		return
	}
	//b.查到了结果且在有效期内：登入成功跳转页面
	now := time.Now().Unix()
	if now -codeRecord.TimeStamp > 1000*60*3{
		//超时了
		s.Ctx.WriteString("验证码已失效，请重新获取验证码!")
		//s.TplName="login_sms.html"
		return
	}
	//判断用户是否已经实名认证，如果未实名，先跳转实名认证
	name := strings.TrimSpace(user.Name)
	card := strings.TrimSpace(user.Card)
	if name == "" || card == ""{
		s.Data["Phone"] = sms.Phone
		s.TplName="user_kyc.html"
	}

	//主页面 列表页面
	//查询列表数据
	records, err := models.QueryRecordByPhone(sms.Phone)
	if err != nil {
		s.Ctx.WriteString("获取认证数据列表失败,请重试！")
		return
	}
	s.Data["Records"]=records
	s.Data["Phone"]=sms.Phone
	s.TplName="list_record.html"
}
