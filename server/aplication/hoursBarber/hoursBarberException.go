package hoursBarber

import (
	"api/server/database"
	"api/server/domain/hoursBarber"
	"api/server/utils"
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

func ListHoursBarberException(ctx *gin.Context) (hoursBarberExceptions []*ListHoursBarberExeption, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	var (
		service = hoursBarber.GetService(hoursBarber.GetRepository(db))
		dados   []hoursBarber.ListHoursBarberExeption
	)
	if dados, err = service.ListExeption(ctx); err != nil {
		log.Printf("Failed to list hours exception: %v", err)
		return
	}

	for i := range dados {
		exeption := &ListHoursBarberExeption{
			ID:            dados[i].ID,
			BarberID:      dados[i].BarberID,
			DateException: utils.FormatDate(dados[i].DateException),
			Reason:        dados[i].Reason,
			CreatedAt:     dados[i].CreatedAt,
			UpdatedAt:     dados[i].UpdatedAt,
		}
		hoursBarberExceptions = append(hoursBarberExceptions, exeption)
	}
	return
}

func DeleteHoursBarberException(ctx *gin.Context, execptionID *int64) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	var (
		service    = hoursBarber.GetService(hoursBarber.GetRepository(db))
		BarberID   *int64
		DataExecao *string
	)

	BarberID, DataExecao, err = service.GetBarberIDByException(ctx, execptionID)
	if err != nil {
		return
	}

	if err = service.DeleteHoursBarberException(ctx, execptionID); err != nil {
		return err
	}

	marked, err := service.MarkReservationAsActive(ctx, BarberID, DataExecao)
	if err != nil {
		log.Printf("Failed to mark reservation as active: %v", err)
		return
	}
	if !marked {
		log.Println("Nenhuma reserva foi marcada como ativa.")
	}

	return

}
