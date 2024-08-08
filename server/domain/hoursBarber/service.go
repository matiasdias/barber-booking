package hoursBarber

import (
	"api/server/infrastructure/persistence/hoursBarber"
	"api/server/utils"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IHoursBarber
}

// GetService retorna um servico para acesso a funções de auxilio à lógica de negócio
func GetService(r IHoursBarber) *Service {
	return &Service{repo: r}
}

// GetRepository retorna um repositório para acesso à camada de hours
func GetRepository(db *sql.DB) IHoursBarber {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, hours *HoursBarber) (err error) {
	dados := &hoursBarber.HoursBarber{
		BarberID:       hours.BarberID,
		DayOfWeek:      hours.DayOfWeek,
		StartTime:      hours.StartTime,
		LunchStartTime: hours.LunchStartTime,
		LunchEndTime:   hours.LunchEndTime,
		EndTime:        hours.EndTime,
	}
	return s.repo.Create(ctx, dados)
}

func (s *Service) ValidateHoursBarber(hours *HoursBarber) (format *FormartHours, err error) {
	if hours.BarberID == nil || *hours.BarberID == 0 {
		return nil, errors.New("missing or invalid BarberID")
	}
	if hours.DayOfWeek == nil || *hours.DayOfWeek == "" {
		return nil, errors.New("missing or invalid day of the week")
	}

	startTime, err := utils.ParseStringFromTime(hours.StartTime)
	if err != nil {
		return nil, err
	}
	lunchStartTime, err := utils.ParseStringFromTime(hours.LunchStartTime)
	if err != nil {
		return nil, err
	}

	lunchEndTime, err := utils.ParseStringFromTime(hours.LunchEndTime)
	if err != nil {
		return nil, err
	}

	endTime, err := utils.ParseStringFromTime(hours.EndTime)
	if err != nil {
		return nil, err
	}

	if startTime != nil && endTime != nil {
		if !startTime.Before(*endTime) {
			return nil, errors.New("start time must be before end time")
		}
	}

	if lunchStartTime != nil && lunchEndTime != nil {
		if !lunchStartTime.Before(*lunchEndTime) {
			return nil, errors.New("LunchStartTime must be before LunchEndTime")
		}
	}

	if startTime != nil && endTime != nil && lunchStartTime != nil && lunchEndTime != nil {
		if !startTime.Before(*lunchStartTime) || !lunchEndTime.Before(*endTime) {
			return nil, errors.New("Invalid lunch break times")
		}
	}

	startTimeStr := startTime.Format("15:04:05")
	endTimeStr := endTime.Format("15:04:05")
	lunchStartTimeStr := lunchStartTime.Format("15:04:05")
	lunchEndTimeStr := lunchEndTime.Format("15:04:05")

	format = &FormartHours{
		StartTime:      &startTimeStr,
		LunchStartTime: &lunchStartTimeStr,
		LunchEndTime:   &lunchEndTimeStr,
		EndTime:        &endTimeStr,
	}

	return format, nil
}

func (s *Service) CheckConflitHoursBarber(ctx *gin.Context, hours *HoursBarber) (conflit bool, err error) {
	dados := &hoursBarber.HoursBarber{
		BarberID:       hours.BarberID,
		DayOfWeek:      hours.DayOfWeek,
		StartTime:      hours.StartTime,
		LunchStartTime: hours.LunchStartTime,
		LunchEndTime:   hours.LunchEndTime,
		EndTime:        hours.EndTime,
	}

	if conflit, err = s.repo.CheckConflitHoursBarber(ctx, dados); err != nil {
		return
	}
	return
}

func (s *Service) List(ctx *gin.Context) (hours []ListHoursBarber, err error) {
	hoursBarberS, err := s.repo.List(ctx)
	if err != nil {
		return
	}
	hours = make([]ListHoursBarber, len(hoursBarberS))
	for i := range hoursBarberS {
		var shop ListHoursBarber

		shop.Barber.Name = hoursBarberS[i].Barber.Name
		shop.Barber.Contato = hoursBarberS[i].Barber.Contato

		shop.HourBarbers = make([]HoursBarbers, len(hoursBarberS[i].HourBarbers))
		for j := range hoursBarberS[i].HourBarbers {
			var h HoursBarbers
			h.DayOfWeek = hoursBarberS[i].HourBarbers[j].DayOfWeek
			h.StartTime = hoursBarberS[i].HourBarbers[j].StartTime
			h.LunchStartTime = hoursBarberS[i].HourBarbers[j].LunchStartTime
			h.LunchEndTime = hoursBarberS[i].HourBarbers[j].LunchEndTime
			h.EndTime = hoursBarberS[i].HourBarbers[j].EndTime
			h.CreatedAt = hoursBarberS[i].HourBarbers[j].CreatedAt
			h.UpdatedAt = hoursBarberS[i].HourBarbers[j].UpdatedAt
			shop.HourBarbers[j] = h
		}

		hours[i] = shop
	}
	return
}

func (s *Service) CreateHoursBarberException(ctx *gin.Context, hoursException *HoursBarberException) (err error) {
	dados := &hoursBarber.HoursBarberException{
		BarberID:      hoursException.BarberID,
		DateException: hoursException.DateException,
		Reason:        hoursException.Reason,
	}
	return s.repo.CreateHoursBarberException(ctx, dados)
}

func (s *Service) ValidateHoursBarberExecption(hoursBarberExecption *HoursBarberException) (format *FormartHoursException, err error) {
	if hoursBarberExecption.BarberID == nil || *hoursBarberExecption.BarberID == 0 {
		return nil, errors.New("missing or invalid BarberID")
	}

	if hoursBarberExecption.Reason == nil || *hoursBarberExecption.Reason == "" {
		return nil, errors.New("missing or invalid Reason")
	}

	dateExecption, err := utils.ParseStringFromDate(hoursBarberExecption.DateException)
	if err != nil {
		return nil, err
	}

	dateExecptionFormat := dateExecption.Format("2006-01-02")

	format = &FormartHoursException{
		DateException: &dateExecptionFormat,
	}

	return format, nil

}

func (s *Service) MarkReservationAsPending(ctx *gin.Context, BarberID *int64, hoursExeptionID *string) (marked bool, err error) {
	return s.repo.MarkReservationAsPending(ctx, BarberID, hoursExeptionID)
}

func (s *Service) HoursExecptionExists(ctx *gin.Context, hoursException *HoursBarberException) (exists bool, err error) {
	dados := &hoursBarber.HoursBarberException{
		BarberID:      hoursException.BarberID,
		DateException: hoursException.DateException,
		Reason:        hoursException.Reason,
	}
	return s.repo.HoursExecptionExists(ctx, dados)
}
