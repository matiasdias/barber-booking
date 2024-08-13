package utils

import (
	"errors"
	"log"
	"regexp"
	"time"
)

// FormatContact converte string para string no formato "(xx) xxxxx-xxxx"
func FormatContact(contact *string) (*string, error) {
	re := regexp.MustCompile("\\D")
	*contact = re.ReplaceAllString(*contact, "")

	if len(*contact) != 11 {
		return nil, errors.New("Telephone must have exactly 11 digits")
	}
	formatTelephone := "(" + (*contact)[:2] + ")" + (*contact)[2:7] + "-" + (*contact)[7:]
	return &formatTelephone, nil
}

// FormatDate converte time.Time para uma string no formato "2006-01-02"
func FormatDate(dateStr *string) *string {
	if dateStr == nil || *dateStr == "" {
		return dateStr
	}
	parsedDate, err := time.Parse(time.RFC3339, *dateStr)
	if err != nil {
		log.Printf("Failed to parse date: %v", err)
		return dateStr
	}
	formattedDate := parsedDate.Format("2006-01-02")
	return &formattedDate
}

// FormatTime converte time.Time para uma string no formato "15:04:05"
func FormatTime(timeStr *string) *string {
	if timeStr == nil || *timeStr == "" {
		return timeStr
	}
	parsedTime, err := time.Parse(time.RFC3339, *timeStr)
	if err != nil {
		log.Printf("Failed to parse time: %v", err)
		return timeStr
	}
	formattedTime := parsedTime.Format("15:04:05")
	return &formattedTime
}
