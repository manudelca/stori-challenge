package repository

import (
	"fmt"
	"github.com/manudelca/stori-challenge/internal/domain"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type ReportRepository interface {
	SendReport(accountInfo domain.AccountInfo, monthInfo []domain.MonthInfo)
}

type reportRepository struct {
}

func NewReportRepository() ReportRepository {
	return &reportRepository{}
}

func (r *reportRepository) SendReport(accountInfo domain.AccountInfo, monthInfo []domain.MonthInfo) {
	// Sender email credentials
	from := os.Getenv("EMAIL")
	if from == "" {
		log.Println("EMAIL environment variable not set")
	}
	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		log.Println("EMAIL_PASSWORD environment variable not set")
	}
	recipient := os.Getenv("EMAIL_RECIPIENT")
	if recipient == "" {
		log.Println("EMAIL_RECIPIENT environment variable not set")
	}
	to := []string{recipient}
	smtpHost := os.Getenv("EMAIL_HOST")
	if smtpHost == "" {
		log.Println("EMAIL_HOST environment variable not set")
	}
	smtpPort := os.Getenv("EMAIL_PORT")
	if smtpPort == "" {
		log.Println("EMAIL_PORT environment variable not set")
	}
	bodyLines := make([]string, 0, 13)
	bodyLines = append(bodyLines, fmt.Sprintf("Total Balance is %.2f\n", accountInfo.TotalBalance))
	log.Printf("accountInfo: %+v", accountInfo)
	log.Printf("monthInfo: %+v", monthInfo)
	for _, month := range monthInfo {
		monthName := time.Month(month.Month).String()
		numberTransactions := month.NumberCreditTransactions + month.NumberDebitTransactions
		avgDebitAmount := month.TotalDebitTransactions / float64(month.NumberDebitTransactions) * (-1)
		avgCreditAmount := month.TotalCreditTransactions / float64(month.NumberCreditTransactions)
		bodyLines = append(bodyLines, fmt.Sprintf("%s:\n", monthName))
		bodyLines = append(bodyLines, fmt.Sprintf("\tNumber of transactions:%d\n", numberTransactions))
		bodyLines = append(bodyLines, fmt.Sprintf("\tAverage debit amount:%.2f\n", avgDebitAmount))
		bodyLines = append(bodyLines, fmt.Sprintf("\tAverage credit amount:%.2f\n", avgCreditAmount))
	}
	log.Printf("Sending: %+v", bodyLines)
	// Email message
	message := []byte("Subject: Account Info report\n" +
		"\n" +
		strings.Join(bodyLines, ""))
	// Set up authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
