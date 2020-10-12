package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type SaveProveController struct {
	beego.Controller
}

func (s *SaveProveController) Get()  {
	s.TplName="save_prove.html"
}

func (s *SaveProveController) Post(){
	f,h,_ :=s.GetFile("uploadOne")
	fileName := h.Filename
	arr :=strings.Split(fileName,":")
	if len(arr)>1{
		index := len(arr)-1
		fileName = arr[index]
	}
	fmt.Println("文件名称：",fileName)
	f.Close()
	s.SaveToFile("uploadOne",path.Join("static/upload",fileName))
	s.TplName="save_prove.html"
}
