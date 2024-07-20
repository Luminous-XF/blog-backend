package email

import (
	"blog-backend/global"
	"blog-backend/pkg/logger"
	"bytes"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
)

func SendEmail(toEmail string, emailParam interface{}) {
	// 生成邮件
	t, err := template.ParseFiles("storage/templates/verify-code.html")
	if err != nil {
		logger.Error("Create email template failed! err:", zap.Error(err))
	}

	var emailBody bytes.Buffer
	err = t.Execute(&emailBody, emailParam)
	if err != nil {
		logger.Error("Create email template failed! err:", zap.Error(err))
	}

	e := &email.Email{
		To:      []string{toEmail},
		From:    global.CONFIG.EmailConfig.Addr,
		Subject: "XOJ VerifyCode",
		HTML:    []byte(emailBody.String()),
	}
	emailAuth := smtp.PlainAuth(
		"",
		global.CONFIG.EmailConfig.Addr,
		global.CONFIG.EmailConfig.LicenseKey,
		"smtp.qq.com",
	)
	// 发送邮件
	err = e.Send("smtp.qq.com:587", emailAuth)
	if err != nil {
		logger.Error("Send email failed! err:", zap.Error(err))
	}
}
