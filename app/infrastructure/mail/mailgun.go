package mail

import (
    "context"
    "fmt"
    "time"
    "github.com/mailgun/mailgun-go/v4"
    mailDomain "github.com/onion0904/CarShareSystem/app/domain/mail"
    "github.com/onion0904/CarShareSystem/app/config"
)

type mailService struct {}

func NewMailRepository() mailDomain.MailService {
	return &mailService{}
}

func(mr *mailService) SendEmail(email string, code string) error {
    mailConfig := config.GetConfig()
    mailgunDomain := mailConfig.Mailgun.Domain
    mailgunPrivateAPIKey := mailConfig.Mailgun.Private_Key
    senderEmail := mailConfig.Mailgun.Sender_email
    recipientEmail := email
    
    // Mailgunクライアントの作成
    mg := mailgun.NewMailgun(mailgunDomain, mailgunPrivateAPIKey)

    // メッセージの作成
    subject := "認証コード"
    body := "認証コード: "+code
    message := mg.NewMessage(
        senderEmail,
        subject,
        body,
        recipientEmail,
    )

    
    // コンテキストの作成（タイムアウト設定）
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    // メールの送信
    _, _, err := mg.Send(ctx, message)
    if err != nil {
        fmt.Println("メールの送信に失敗しました:", err)
        return err
    }

    return nil
}
