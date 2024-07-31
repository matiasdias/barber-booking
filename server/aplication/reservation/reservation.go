package reservation

import (
	"api/server/database"
	"api/server/domain/reservation"
	"api/server/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, reser *CreateReservation) error {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()
	service := reservation.GetService(reservation.GetRepository(db))

	r := &reservation.Reservation{
		BarberID:        reser.BarberID,
		ClientID:        reser.ClientID,
		BarberShopID:    reser.BarberShopID,
		ServiceID:       reser.ServiceID,
		DateReservation: reser.DateReservation,
		StartTime:       reser.StartTime,
		Status:          reser.Status,
	}

	formatHours, err := service.ValidateHoursRservation(r)
	if err != nil {
		log.Printf("Failed to validate reservation hours: %v", err)
		return err
	}

	r.StartTime = formatHours.StartTime
	r.DateReservation = formatHours.DateReservation

	if err = service.CheckConflictReservation(ctx, r); err != nil {
		log.Printf("Failed to check conflict reservation: %v", err)
		return err
	}

	// Verifica se a data da reserva está marcada como exceção
	exists, err := service.CheckExceptionForBarber(ctx, r.BarberID, r.DateReservation)
	if err != nil {
		log.Printf("Failed to check exception for barber: %v", err)
		return err
	}
	if exists {
		err = fmt.Errorf("A data %s está marcada como exceção de trabalho para este barbeiro.", *r.DateReservation)
		log.Println(err)
		return err
	}

	if err = service.Create(ctx, r); err != nil {
		log.Printf("Fails to add reservation: %v", err)
		return err
	}

	return nil
}

func List(ctx *gin.Context) (reservations []reservation.ReservationList, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	server := reservation.GetService(reservation.GetRepository(db))
	reservations, err = server.List(ctx)
	if err != nil {
		log.Printf("Failed to list reservations: %v", err)
		return
	}
	return
}

func Update(ctx *gin.Context, reservationID *int64, reser *UpdateReservationReq) error {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()
	service := reservation.GetService(reservation.GetRepository(db))
	dados := &reservation.Reservation{
		ID:              reservationID,
		BarberID:        reser.BarberID,
		DateReservation: reser.DateReservation,
		StartTime:       reser.StartTime,
		Status:          reser.Status,
	}

	dateReservation, err := utils.ParseStringFromDate(reser.DateReservation)
	if err != nil {
		return err
	}
	dateReservationStrFormatted := dateReservation.Format("2006-01-02")
	dados.DateReservation = &dateReservationStrFormatted

	if err = service.CheckConflictReservation(ctx, dados); err != nil {
		log.Printf("Failed to check conflict reservation: %v", err)
		return err
	}

	exists, err := service.CheckExceptionForBarber(ctx, dados.BarberID, dados.DateReservation)
	if err != nil {
		log.Printf("Failed to check exception for barber: %v", err)
		return err
	}
	if exists {
		err = fmt.Errorf("A data %s está marcada como exceção de trabalho para este barbeiro.", *dados.DateReservation)
		log.Println(err)
		return err
	}

	if err = service.UpdateReservation(ctx, reservationID, dados); err != nil {
		log.Printf("Failed to update reservation: %v", err)
		return err
	}
	return nil
}
