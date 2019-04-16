package main

import (
	"crypto/tls"
	"gosmtp/pkg/config"
	"gosmtp/pkg/smpt"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	config := config.New()

	smtpServer := &smpt.SmtpServer{
		Host: config.Host,
		Port: config.Port,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         config.Host,
		},
		Auth: smtp.PlainAuth("", config.Sender, config.Password, config.Host),
	}

	mail := &smpt.Mail{
		Sender:  config.Sender,
		To:      strings.Split(config.To, ","),
		Subject: config.Subject,
		Body:    config.Body,
	}

	serivce := smpt.New(smtpServer, mail)
	err := serivce.Send()
	if err != nil {
		log.Printf("[Error] Unable to send emails: %v", err)
	}
}
