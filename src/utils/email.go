package utils
import (
    "gopkg.in/gomail.v2"
	"bd2-backend/src/config"

	
)

type EmailCredentials struct {
	smtpUsername string
	smtpPassword string
}

var credentials EmailCredentials

func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		ErrorLogger.Fatal("cannot load config:", err)
	}
	credentials = EmailCredentials{
		smtpUsername: cfg.SMTPUsername,
		smtpPassword: cfg.SMTPPassword,
	}	
}

func SendEmail(to string, subject string, body string) error {

	m := gomail.NewMessage()
    m.SetHeader("From", credentials.smtpUsername)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, credentials.smtpUsername, credentials.smtpPassword)

	err := d.DialAndSend(m)
	if err != nil {
		ErrorLogger.Fatalf("cannot sendEmail to %s, error %v", to, err)
		return err
		
	}
	return nil
}

