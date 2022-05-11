package mail

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

// MailClient object interacrs with the smtp server
type MailClient struct {
	l        *log.Logger
	username string
	password string
}

// NewMailClient returns a new instance of MailClient with a logger
func NewMailClient(username, password string) *MailClient {
	l := log.New(os.Stdout, "{MAIL-CLIENT}", log.LstdFlags)
	return &MailClient{l, username, password}
}

// SendConfirmEmail will send a confirmation email to the given email along with their contestant id
func (mc MailClient) SendConfirmEmail(email, userID, name string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "anotherstory105@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Audition Confirmation")
	body := fmt.Sprintf("Hi %s this is to confirm your registration to audition for another story. You are contestant number %s", name, userID)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, mc.username, mc.password)
	err := d.DialAndSend(m)
	if err != nil {
		mc.l.Printf("unable to send confirmation email %s", err.Error())
		return errors.New("unable to send confirmation email pls try again ")
	}
	return nil
}
