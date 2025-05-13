package mail

import (
	"log"
	"net/smtp"
	"github.com/jordan-wright/email"
	
	mailDomain "github.com/onion0904/CarShareSystem/app/domain/mail"
	"github.com/onion0904/CarShareSystem/app/config"
)

type mailService struct {}

func NewMailRepository() mailDomain.MailService {
	return &mailService{}
}

func (mr *mailService)SendEmail(toEmail string, code string)error{
	mailConfig := config.GetConfig()
    gmailpass := mailConfig.Mail.GmailPass

	e := email.NewEmail()
	e.From = "takerudaze0904@gmail.com"
	e.To = []string{toEmail}
	e.Subject = "CarShareSystemの確認コード"
	e.Text = []byte("見覚えのない連絡でしたら無視してください\n確認コード:"+code)
	// GmailのSMTPサーバ情報
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	smtpAddr := smtpServer + ":" + smtpPort

	err := e.Send(smtpAddr, smtp.PlainAuth(
		"",               // identity（通常は空文字）
		"takerudaze0904@gmail.com",   // Gmailアドレス
		gmailpass,      // Appパスワード
		smtpServer,       // ホスト名（smtp.gmail.com）
	))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully")
	return nil
}