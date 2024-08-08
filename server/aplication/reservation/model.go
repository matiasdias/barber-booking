package reservation

import "time"

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

type Reserva struct {
	DateReservation         *string    `conversor:"data_reserva" json:"data_reserva"`
	DateReservationOriginal *string    `conversor:"data_reserva_original" json:"data_reserva_original"`
	StartTime               *string    `conversor:"horario_inicial" json:"horario_inicial"`
	EndTime                 *string    `conversor:"horario_final" json:"horario_final"`
	Status                  *string    `conversor:"status" json:"status"`
	CreatedAt               *time.Time `conversor:"data_criacao" json:"data_criacao"`
	UpdatedAt               *time.Time `conversor:"data_atualizacao" json:"data_atualizacao"`
}

type Barber struct {
	Name    *string `json:"nome" conversor:"nome"`
	Contact *string `json:"contato" conversor:"contato"`
}

type Client struct {
	Name    *string `json:"nome" conversor:"nome"`
	Email   *string `json:"email" conversor:"email"`
	Contact *string `json:"contato" conversor:"contato"`
}

type Services struct {
	Name     *string  `json:"nome" conversor:"nome"`
	Price    *float64 `json:"preco" conversor:"preco"`
	Duration *string  `json:"duracao" conversor:"duracao"`
}

type BarberShop struct {
	Name             *string `json:"nome" conversor:"nome"`
	Cidade           *string `json:"cidade" conversor:"cidade"`
	Rua              *string `json:"rua" conversor:"rua"`
	NumeroResidencia *int64  `json:"numero_residencia" conversor:"numero_residencia"`
	PontoReferencia  *string `json:"ponto_referencia" conversor:"ponto_referencia"`
	Contact          *string `json:"contato" conversor:"contato"`
}

type ReservationList struct {
	Shop         BarberShop `json:"barbearia" conversor:"barbearia"`
	Barber       Barber     `json:"barbeiro" conversor:"barbeiro"`
	Client       Client     `json:"cliente" conversor:"cliente"`
	Service      Services   `json:"servico" conversor:"servico"`
	Reservations []Reserva  `json:"reservas" conversor:"reservas"`
}
