package smpt

import (
	"log"
	"net/smtp"
)

type SmtpService struct {
	server SmtpServerI
	mail   *Mail
	Auth   smtp.Auth
}

func New(s *SmtpServer, m *Mail) *SmtpService {
	return &SmtpService{
		server: s,
		mail:   m,
		Auth:   smtp.PlainAuth("", s.Email, s.Password, s.Host),
	}
}

func (s *SmtpService) Send() error {
	client, err := s.server.new()
	if err != nil {
		return err
	}

	if err = client.Auth(s.Auth); err != nil {
		return err
	}

	if err = client.Mail(s.mail.Sender); err != nil {
		return err
	}

	receivers := append(s.mail.To, s.mail.Cc...)
	receivers = append(receivers, s.mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(s.mail.BuildMessage()))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	return nil
}
