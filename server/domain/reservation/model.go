package reservation

import "time"

type Reservation struct {
	ID                      *int64  `converson:"id"`
	BarberID                *int64  `converson:"barbeiro_id"`
	ClientID                *int64  `converson:"cliente_id"`
	BarberShopID            *int64  `converson:"barbearia_id"`
	DateReservation         *string `converson:"data_reserva"`
	DateReservationOriginal *string `converson:"data_reserva_original"`
	StartTime               *string `converson:"horario_inicial"`
	EndTime                 *string `converson:"horario_final"`
	Duration                *string `converson:"duracao"`
	Status                  *string `converson:"status"`
}

type FormartHours struct {
	StartTime       *string `conversor:"horario_inicial"`
	EndTime         *string `conversor:"horario_final"`
	DateReservation *string `converson:"data_reserva"`
	Duration        *string `converson:"duracao"`
}

type Reserva struct {
	DateReservation         *string    `json:"data_reserva" conversor:"data_reserva"`
	DateReservationOriginal *string    `json:"data_reserva_original" conversor:"data_reserva_original"`
	StartTime               *string    `json:"horario_inicial" conversor:"horario_inicial"`
	EndTime                 *string    `json:"horario_final" conversor:"horario_final"`
	Duration                *string    `json:"duracao" conversor:"duracao"`
	Status                  *string    `json:"status" conversor:"status"`
	CreatedAt               *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdatedAt               *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
	DataSuspensao           *time.Time `json:"data_suspensao" conversor:"data_suspensao"`
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
	Reservations []Reserva  `json:"reservas" conversor:"reservas"`
}
