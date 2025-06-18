package mail

type MailService interface {
	SendEmail(email string, code string) error
}
