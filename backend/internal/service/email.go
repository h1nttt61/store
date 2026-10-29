package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

type EmailService struct {
	rdb *redis.Client
}

func NewEmailService(rdb *redis.Client) *EmailService {
	return &EmailService{rdb: rdb}
}

func (s *EmailService) SendCode(ctx context.Context, email string) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	err := s.rdb.Set(ctx, "email_code:"+email, code, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	htmlBody := GenerateEmailHTML(code)

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = email
	headers["Subject"] = "Код подтверждения - Kittens Shop"
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EmailService) VerifyCode(ctx context.Context, email, code string) error {
	stored, err := s.rdb.Get(ctx, "email_code:"+email).Result()
	if err != nil {
		return errors.New("code expired or not found")
	}

	if stored != code {
		return errors.New("invalid code")
	}

	return nil
}
