package hoursBarber

import (
	"api/server/database"
	"api/server/domain/hoursBarber"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateHoursBarberExecption(ctx *gin.Context, hoursBarberExecepion *CreateException) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	service := hoursBarber.GetService(hoursBarber.GetRepository(db))

	dados := &hoursBarber.HoursBarberException{
		BarberID:      hoursBarberExecepion.BarberID,
		DateException: hoursBarberExecepion.DateException,
		Reason:        hoursBarberExecepion.Reason,
	}

	formatDateExecption, err := service.ValidateHoursBarberExecption(dados)
	if err != nil {
		log.Printf("Failed to validate hours execption: %v", err)
		return
	}
	dados.DateException = formatDateExecption.DateException

	exists, err := service.HoursExecptionExists(ctx, dados)
	if err != nil {
		log.Printf("Failed to check if hours exception exists: %v", err)
		return
	}
	if exists {
		err = errors.New("exception already exists")
		return
	}

	marked, err := service.MarkReservationAsPending(ctx, dados.BarberID, dados.DateException)
	if err != nil {
		log.Printf("Failed to mark reservation as pending: %v", err)
		return
	}
	if !marked {
		log.Println("Nenhuma reserva foi marcada como pendente.")
	}

	if err = service.CreateHoursBarberException(ctx, dados); err != nil {
		log.Printf("Fails to add barber hours execption: %v", err)
		return
	}

	return
}
