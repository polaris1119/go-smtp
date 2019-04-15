package gosmtp

import (
	"crypto/tls"
	"net/smtp"
)

type SmtpServer struct {
	Host      string
	Port      string
	TlsConfig *tls.Config
	Auth      smtp.Auth
}

func (s *SmtpServer) dial() (*tls.Conn, error) {
	return tls.Dial("tcp", s.serverName(), s.TlsConfig)
}

func (s *SmtpServer) serverName() string {
	return s.Host + ":" + s.Port
}
