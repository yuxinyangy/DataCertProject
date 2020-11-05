package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type SaveProveController struct {
	beego.Controller
}

func (s *SaveProveController) Get() {
	phone := s.GetString("phone")
	s.Data["Phone"] = phone
	s.TplName = "home.html"
}

func (s *SaveProveController) Post() {
	//标题
	fileTitle := s.Ctx.Request.PostFormValue("upload_title")
	phone := s.Ctx.Request.PostFormValue("phone")
	//文件
	file, header, err := s.GetFile("upload_file")
	if err != nil {
		s.TplName = "error.html"
		return
	}
	//关闭文件
	defer file.Close()

	fmt.Println("自定义的文件标题:", fileTitle)
	fmt.Println("文件名称:", header.Filename)
	fmt.Println("文件的大小:", header.Size) //字节大小
	//保存文件到本地目录
	uploadDir := "./static/upload/" + header.Filename
	//arr :=strings.Split(fileName,":")
	//if len(arr)>1{
	//	index := len(arr)-1
	//	fileName = arr[index]
	//}
	//fmt.Println("文件名称：",fileName)
	//f.Close()
	//s.SaveToFile("uploadOne",path.Join("static/upload",fileName))
	//s.TplName="home.html"
	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)
	writer := bufio.NewWriter(saveFile)
	_, err = io.Copy(writer, file)
	if err != nil {
		fmt.Println(err.Error())
		s.TplName = "error.html"
		return
	}
	defer saveFile.Close()
	//计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := util.MD5HashReader(hashFile)


	t :=time.Now().Unix()
	//保存到数据库中
	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix()
	record.FileCert = hash
	record.Phone = phone
	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println(err.Error())
		s.TplName = "error.html"
		return
	}

	//新增逻辑：将要认证的文件hash值及个人实名信息，保存到区块链上，即上链
	//①准备认证数据的用户相关的数据
	us,err := models.QueryUserByPhone(phone)
	fmt.Println("用户的信息：", us.Name, us.Card)
	if err != nil {
		fmt.Println(err.Error())
		s.Ctx.WriteString("抱歉，数据认证失败，请重试")
		return
	}
	certhash,_ :=util.SHA256HashReader(hashFile)
	certRecord := models.CertRecord{
		CertId:     []byte(hash),
		CertHash:   []byte(certhash),
		CertAuthor: us.Name,
		AuthorCard: us.Card,
		Phone:      us.Phone,
		FileName:   header.Filename,
		FileSize:   header.Size,
		CertTime:   t,
	}
	//序列化
	certBytes,err := certRecord.SerializeRecord()
	_, err = blockchain.CHAIN.SaveData(certBytes)
	if err != nil {
		s.TplName = "error.html"
		return
	}
	//从数据库中读取phone用户对应的所有认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		s.TplName = "error.html"
		return
	}
	s.Data["Records"] = records
	s.Data["Phone"] = phone
	s.TplName = "list_record.html"
}
