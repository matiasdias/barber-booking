package reservation

type CreateReservation struct {
	ID                     *int64  `json:"id" converson:"id"`
	BarberID               *int64  `json:"barbeiro_id" converson:"barbeiro_id" binding:"required"`
	ClientID               *int64  `json:"cliente_id" converson:"cliente_id" binding:"required"`
	BarberShopID           *int64  `json:"barbearia_id" converson:"barbearia_id" binding:"required"`
	DateReservation        *string `json:"data_reserva" converson:"data_reserva" binding:"required"`
	DateRservationOriginal *string `json:"data_reserva_original" converson:"data_reserva_original"`
	StartTime              *string `json:"horario_inicial" converson:"horario_inicial" binding:"required"`
	EndTime                *string `json:"horario_final" converson:"horario_final"`
	Duration               *string `json:"duracao" converson:"duracao" binding:"required"`
	Status                 *string `json:"status" converson:"status"`
}
