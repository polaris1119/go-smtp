package smpt

import (
	"log"
	"net/smtp"
)

type SmtpService struct {
	s *SmtpServer
	m *Mail
}

func New(smtp *SmtpServer, mail *Mail) *SmtpService {
	return &SmtpService{
		s: smtp,
		m: mail,
	}
}

func (s *SmtpService) Send() error {
	conn, err := s.s.dial()
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, s.s.Host)
	if err != nil {
		return err
	}

	if err = client.Auth(s.s.Auth); err != nil {
		return err
	}

	if err = client.Mail(s.m.Sender); err != nil {
		return err
	}

	receivers := append(s.m.To, s.m.Cc...)
	receivers = append(receivers, s.m.Bcc...)
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

	_, err = w.Write([]byte(s.m.BuildMessage()))
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
