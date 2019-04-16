package main

import (
	"crypto/tls"
	"gosmtp"
	"gosmtp/config"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	config := config.New()

	mail := &gosmtp.Mail{}
	mail.Sender = config.Sender
	mail.To = strings.Split(config.To, ",")
	mail.Subject = config.Subject
	mail.Body = config.Body

	s := &gosmtp.SmtpServer{
		Host: config.Host,
		Port: config.Port,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         config.Host,
		},
		Auth: smtp.PlainAuth("", config.Sender, config.Password, config.Host),
	}

	serivce := gosmtp.New(s, mail)
	err := serivce.Send()
	if err != nil {
		log.Printf("[Error] Unable to send emails: %v", err)
	}
}
