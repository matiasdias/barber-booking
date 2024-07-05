package utils

import (
	"errors"
	"regexp"
)

func FormatContact(contact *string) (*string, error) {
	re := regexp.MustCompile("\\D")
	*contact = re.ReplaceAllString(*contact, "")

	if len(*contact) != 11 {
		return nil, errors.New("Telephone must have exactly 11 digits")
	}
	formatTelephone := "(" + (*contact)[:2] + ")" + (*contact)[2:7] + "-" + (*contact)[7:]
	return &formatTelephone, nil
}
