package utils

import (
	"errors"
	"time"
)

// ParseStringFromTime converte string para time.Time no formato de horas
func ParseStringFromTime(timeStr *string) (*time.Time, error) {
	if timeStr == nil || *timeStr == "" {
		return nil, errors.New("time string is empty")
	}
	parseTime, err := time.Parse("15:04:05", *timeStr)
	if err != nil {
		return nil, err
	}
	return &parseTime, nil
}

// ParseStringFromDate converte string para time.Time no formato de data
func ParseStringFromDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, nil
	}
	parsedDate, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, err
	}
	return &parsedDate, nil
}

func ParseDuration(durationStr *string) (time.Duration, error) {
	duration, err := time.ParseDuration(*durationStr)
	if err != nil {
		return 0, errors.New("invalid duration format")
	}
	return duration, nil
}
