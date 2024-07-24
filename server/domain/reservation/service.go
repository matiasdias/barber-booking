package reservation

import (
	"api/server/infrastructure/persistence/reservation"
	"api/server/utils"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IReservation
}

func GetService(r IReservation) *Service {
	return &Service{repo: r}
}

func GetRepository(db *sql.DB) IReservation {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, reser *Reservation) (err error) {
	dados := &reservation.Reservation{
		BarberID:        reser.BarberID,
		ClientID:        reser.ClientID,
		BarberShopID:    reser.BarberShopID,
		DateReservation: reser.DateReservation,
		StartTime:       reser.StartTime,
		EndTime:         reser.EndTime,
		Duration:        reser.Duration,
		Status:          reser.Status,
	}

	return s.repo.Create(ctx, dados)
}

func (s *Service) ValidateHoursRservation(reser *Reservation) (*FormartHours, error) {
	if reser.BarberID == nil || *reser.BarberID == 0 {
		return nil, errors.New("missing or invalid BarberID")
	}
	if reser.ClientID == nil || *reser.ClientID == 0 {
		return nil, errors.New("missing or invalid ClientID")
	}
	if reser.BarberShopID == nil || *reser.BarberShopID == 0 {
		return nil, errors.New("missing or invalid BarberShop")
	}
	if reser.Duration == nil || *reser.Duration == "" {
		return nil, errors.New("missing or invalid Duration")
	}

	startTime, err := utils.ParseStringFromTime(reser.StartTime)
	if err != nil {
		return nil, err
	}

	dateReservation, err := utils.ParseStringFromDate(reser.DateReservation)
	if err != nil {
		return nil, err
	}

	duration, err := utils.ParseDuration(reser.Duration)
	if err != nil {
		return nil, err
	}

	startTimeStrFormatted := startTime.Format("15:04:05")
	dateReservationStrFormatted := dateReservation.Format("2006-01-02")
	durationStrFormatted := duration.String()

	format := &FormartHours{
		StartTime:       &startTimeStrFormatted,
		DateReservation: &dateReservationStrFormatted,
		Duration:        &durationStrFormatted,
	}

	return format, nil

}

func (s *Service) CheckConflictReservation(ctx *gin.Context, reser *Reservation) (err error) {
	dados := &reservation.Reservation{
		BarberID:        reser.BarberID,
		ClientID:        reser.ClientID,
		BarberShopID:    reser.BarberShopID,
		DateReservation: reser.DateReservation,
		StartTime:       reser.StartTime,
		EndTime:         reser.EndTime,
		Duration:        reser.Duration,
		Status:          reser.Status,
	}

	return s.repo.CheckConflictReservation(ctx, dados)
}

func (s *Service) List(ctx *gin.Context) (reservations []ReservationList, err error) {
	res, err := s.repo.List(ctx)
	if err != nil {
		return
	}

	// Inicializa a lista de ReservationList com o mesmo comprimento da lista de resultados
	reservations = make([]ReservationList, len(res))
	for i := range res {
		var reser ReservationList

		reser.Barber.Name = res[i].Barber.Name
		reser.Barber.Contact = res[i].Barber.Contact

		reser.Shop.Name = res[i].Shop.Name
		reser.Shop.Cidade = res[i].Shop.Cidade
		reser.Shop.Rua = res[i].Shop.Rua
		reser.Shop.NumeroResidencia = res[i].Shop.NumeroResidencia
		reser.Shop.PontoReferencia = res[i].Shop.PontoReferencia
		reser.Shop.Contact = res[i].Shop.Contact

		reser.Client.Name = res[i].Client.Name
		reser.Client.Email = res[i].Client.Email
		reser.Client.Contact = res[i].Client.Contact

		// inicializa a lista de reservas
		reser.Reservations = make([]Reserva, len(res[i].Reservations))
		for j := range res[i].Reservations {
			var r Reserva
			r.DateReservation = res[i].Reservations[j].DateReservation
			r.StartTime = res[i].Reservations[j].StartTime
			r.EndTime = res[i].Reservations[j].EndTime
			r.Duration = res[i].Reservations[j].Duration
			r.Status = res[i].Reservations[j].Status
			r.CreatedAt = res[i].Reservations[j].CreatedAt
			r.UpdatedAt = res[i].Reservations[j].UpdatedAt
			r.DataSuspensao = res[i].Reservations[j].DataSuspensao
			reser.Reservations[j] = r
		}
		reservations[i] = reser
	}
	return
}
