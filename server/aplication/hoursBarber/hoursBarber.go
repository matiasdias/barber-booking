package hoursBarber

import (
	"api/server/database"
	"api/server/domain/hoursBarber"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, hours *CreateHoursBarber) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return errors.New("failed to connect to database")
	}
	defer db.Close()

	service := hoursBarber.GetService(hoursBarber.GetRepository(db))

	dados := &hoursBarber.HoursBarber{
		ID:             hours.ID,
		BarberID:       hours.BarberID,
		DayOfWeek:      hours.DayOfWeek,
		StartTime:      hours.StartTime,
		LunchStartTime: hours.LunchStartTime,
		LunchEndTime:   hours.LunchEndTime,
		EndTime:        hours.EndTime,
	}
	formatHours, err := service.ValidateHoursBarber(dados)
	if err != nil {
		log.Printf("Failed to validate reservation hours: %v", err)
		return err
	}

	dados.StartTime = formatHours.StartTime
	dados.LunchStartTime = formatHours.LunchStartTime
	dados.LunchEndTime = formatHours.LunchEndTime
	dados.EndTime = formatHours.EndTime

	conflit, err := service.CheckConflitHoursBarber(ctx, dados)
	if err != nil {
		log.Printf("Failed to check conflict hours: %v", err)
		return err
	}
	if conflit {
		return errors.New("horário de trabalho já existente para o barbeiro no dia da semana especificado")
	}

	if err = service.Create(ctx, dados); err != nil {
		log.Printf("Fails to add barber hours: %v", err)
		return err
	}

	return
}

func ListHourBarber(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarber, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, errors.New("failed to connect to database")
	}
	defer db.Close()

	service := hoursBarber.GetService(hoursBarber.GetRepository(db))
	hoursBarbers, err = service.List(ctx)
	if err != nil {
		log.Printf("Failed to list hours barbers: %v", err)
		return nil, err
	}
	return
}
