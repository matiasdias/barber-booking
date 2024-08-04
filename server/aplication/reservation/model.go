package reservation

type CreateReservation struct {
	BarberID        *int64  `json:"barbeiro_id" converson:"barbeiro_id" binding:"required"`
	ClientID        *int64  `json:"cliente_id" converson:"cliente_id" binding:"required"`
	BarberShopID    *int64  `json:"barbearia_id" converson:"barbearia_id" binding:"required"`
	ServiceID       *int64  `json:"servico_id" converson:"servico_id" binding:"required"`
	DateReservation *string `json:"data_reserva" converson:"data_reserva" binding:"required"`
	StartTime       *string `json:"horario_inicial" converson:"horario_inicial" binding:"required"`
}

type UpdateReservationReq struct {
	ID                     *int64  `json:"id" converson:"id"`
	BarberID               *int64  `json:"barbeiro_id" converson:"barbeiro_id"`
	DateReservation        *string `json:"data_reserva" converson:"data_reserva"`
	DateRservationOriginal *string `converson:"data_reserva_original"`
	StartTime              *string `json:"horario_inicial" converson:"horario_inicial"`
	Status                 *string `json:"status" converson:"status"`
	ServiceID              *int64  `json:"servico_id" converson:"servico_id"`
}
