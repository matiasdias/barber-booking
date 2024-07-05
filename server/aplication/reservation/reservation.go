package reservation

import (
	"api/server/database"
	"api/server/domain/reservation"
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
		DateReservation: reser.DateReservation,
		StartTime:       reser.StartTime,
		Duration:        reser.Duration,
		Status:          reser.Status,
	}

	if err = service.CheckConflictReservation(ctx, r); err != nil {
		log.Printf("Failed to check conflict reservation: %v", err)
		return err
	}

	formatHours, err := service.ValidateHoursRservation(r)
	if err != nil {
		log.Printf("Failed to validate reservation hours: %v", err)
		return err
	}

	r.StartTime = formatHours.StartTime
	r.DateReservation = formatHours.DateReservation
	r.Duration = formatHours.Duration

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
