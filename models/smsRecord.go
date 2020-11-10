package models

import (
	"DataCertProject/db_mysql"
)

type SmsRecord struct {
	BizId     string `form:"biz_id"`  //业务号
	Phone     string `form:"phone"`   //手机号
	Code      string `form:"code"`    //验证码
	Status    string `form:"status"`  //阿里云状态码
	Message   string `form:"message"` //短信调用sdk信息
	TimeStamp int64  `form:"timestamp"`
}

/**
 *该方法根据BizId，phone以及code条件查询出符合条件的验证码记录
 */

func (s SmsRecord) QuerySmsByBizId() (*SmsRecord,error) {
	var sms SmsRecord
	row := db_mysql.Db.QueryRow("select biz_id, phone, code, status, message, timestamp from sms_record where biz_id = ? and phone = ? and code = ?",
		s.BizId, s.Phone, s.Code)
	err := row.Scan(&sms.BizId,&sms.Phone,&sms.Code,&sms.Status,&sms.Message,&sms.TimeStamp)
	if err != nil {
		return nil,err
	}
	return &sms,nil
}

/*
 *该方法用于将smsrecord结构体实例保存到数据库中
 */
func (s SmsRecord) SaveSms() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp)"+"values(?,?,?,?,?,?)",
		s.BizId, s.Phone, s.Code, s.Status, s.Message, s.TimeStamp)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
