package controllers

import (
	"DataCertProject/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get() {
	//0.获取前端页面get请求时携带的cert_id数据
	certId := c.GetString("cert_id")
	fmt.Println("要查询的认证ID:", certId)
	//1.准备数据:根据cert_id到区块链上查询具体信息，获得到区块信息
	block, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		c.TplName = "error.html"
		return
	}
	//查询未遇到错误,有两种情况:查到了和未查到
	if block == nil {
		c.Ctx.WriteString("抱歉,未查询到链上数据，请重试")
	}
	//certId = hex.EncodeToString(block.Data)
	c.Data["CertId"] = strings.ToUpper(string(block.Data))

	//2.跳转页面
	c.TplName = "cert_detail.html"
}
