package smpt

import (
	"crypto/tls"
	"net/smtp"
)

type Smtper interface {
	new() (*smtp.Client, error)
}

type SmtpServer struct {
	Host      string
	Port      string
	Email     string
	Password  string
	TlsConfig *tls.Config
}

var _ Smtper = (*SmtpServer)(nil)

func (s *SmtpServer) new() (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", s.serverName(), s.TlsConfig)
	if err != nil {
		return nil, err
	}
	return smtp.NewClient(conn, s.Host)
}

func (s *SmtpServer) serverName() string {
	return s.Host + ":" + s.Port
}
