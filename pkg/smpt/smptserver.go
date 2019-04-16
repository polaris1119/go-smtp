package smpt

import (
	"crypto/tls"
)

type SmtpServer struct {
	Host      string
	Port      string
	Email     string
	Password  string
	TlsConfig *tls.Config
}

func (s *SmtpServer) dial() (*tls.Conn, error) {
	return tls.Dial("tcp", s.serverName(), s.TlsConfig)
}

func (s *SmtpServer) serverName() string {
	return s.Host + ":" + s.Port
}
