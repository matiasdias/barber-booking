package hoursBarber

import (
	"api/server/database"
	"api/server/domain/hoursBarber"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

//-- Verifica se a data da reserva está marcada como uma exceção
//    IF EXISTS (
//        SELECT 1
//        FROM horario_excecao
//        WHERE id_barbeiro = NEW.id_barbeiro
//        AND data_excecao = NEW.data_reserva
//    ) THEN
//        RAISE EXCEPTION 'A data % está marcada como exceção para o barbeiro.', NEW.data_reserva;
//    END IF;

func CreateHoursBarberExecption(ctx *gin.Context, hoursBarberExecepion *CreateException) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	service := hoursBarber.GetService(hoursBarber.GetRepository(db))

	dados := &hoursBarber.HoursBarberException{
		ID:            hoursBarberExecepion.ID,
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

	if err = service.MarkReservationAsPending(ctx, dados.BarberID, dados.DateException); err != nil {
		log.Printf("Failed to mark reservation as pending: %v", err)
		return
	}

	exists, err := service.HoursExecptionExists(ctx, dados)
	if err != nil {
		log.Printf("Failed to check if hours exception exists: %v", err)
		return
	}
	if exists {
		err = errors.New("exception already exists")
		return
	}

	if err = service.CreateHoursBarberException(ctx, dados); err != nil {
		log.Printf("Fails to add barber hours execption: %v", err)
		return
	}

	return
}
