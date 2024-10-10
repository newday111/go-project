package utils

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(email string, subject string, body string) error {
	/*
		email:	接收人的邮箱
		subject:	邮件主题
		body:	邮件内容
	*/

	//	读取关于邮件的配置信息
	smtpHost := viper.GetString("email_config.smtp_host")
	smtpPort := viper.GetInt("email_config.smtp_port")
	smtpUser := viper.GetString("email_config.smtp_user")
	smtpPass := viper.GetString("email_config.smtp_pass")

	m := gomail.NewMessage()
	m.SetHeader("From", "jackgeng1314@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	err := d.DialAndSend(m)
	return err
}
