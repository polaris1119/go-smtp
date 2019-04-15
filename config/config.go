package config

import (
	"os"
)

// Config contains all the environment varialbes.
type Config struct {
	Host     string
	Port     string
	Password string
	Sender   string
	To       string
	Subject  string
	Body     string
}

func New() Config {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "465"
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "smtp.gmail.com"
	}

	sender, ok := os.LookupEnv("SENDER")
	if !ok {
		sender = "account@gmail.com"
	}

	password, ok := os.LookupEnv("PASSWORD")
	if !ok {
		password = "password"
	}

	to, ok := os.LookupEnv("TO")
	if !ok {
		to = "poyatehu@maillink.top,poyatehu@maillink.top"
	}

	subject, ok := os.LookupEnv("SUBJECT")
	if !ok {
		subject = "Hello from StudyGolang"
	}

	body, ok := os.LookupEnv("BODY")
	if !ok {
		body = "Happy Coding!"
	}

	return Config{
		Port:     port,
		Host:     host,
		Password: password,
		Sender:   sender,
		To:       to,
		Subject:  subject,
		Body:     body,
	}
}
