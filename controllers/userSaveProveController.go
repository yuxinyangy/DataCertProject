package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type SaveProveController struct {
	beego.Controller
}
func (s *SaveProveController) Post(){
	f,h,_ :=s.GetFile("uploadOne")
	fileName := h.Filename
	//arr :=strings.Split(fileName,":")
	//if len(arr)>1{
	//	index := len(arr)-1
	//	fileName = arr[index]
	//}
	//fmt.Println("文件名称：",fileName)
	//f.Close()
	//s.SaveToFile("uploadOne",path.Join("static/upload",fileName))
	//s.TplName="home.html"
	uploadDir :="./static/upload/"+fileName
	saveFile,err :=os.OpenFile(uploadDir,os.O_RDWR|os.O_CREATE,777)
	writer := bufio.NewWriter(saveFile)
	file_size,err :=io.Copy(writer,f)
	if err != nil {
		s.TplName="error.html"
		return
	}
	fmt.Println(file_size)

}
