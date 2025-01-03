package dao

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
)

// GetDialer 获取邮箱连接
func (d *Dao) GetDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(d.c.Email.EmailHost, d.c.Email.EmailPort, d.c.Email.EmailUser, d.c.Email.EmailPwd)
	return dialer
}

// GenEmailHeaderMessage 设置邮件头部信息 邮件一件一件发，防止用户隐私暴露
func (d *Dao) GenEmailHeaderMessage(addr string, htmlText string, title string) (m *gomail.Message) {
	m = gomail.NewMessage()
	m.SetHeader("From", d.c.Email.EmailUser)
	m.SetHeader("To", addr)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", htmlText)
	return
}

func (d *Dao) SendEmail(ctx context.Context, email, content, title string) (err error) {
	message := d.GenEmailHeaderMessage(email, content, title)
	sendErr := d.GetDialer().DialAndSend(message)
	if sendErr != nil {
		fmt.Printf("[dao | SendEmail] send developer email failed err:%+v", sendErr)
	}
	fmt.Printf("[dao | SendEmail] send email success email:%+v\n", email)
	return
}
