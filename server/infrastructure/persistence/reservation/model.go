package reservation

import "time"

type Reservation struct {
	ID                     *int64  `converson:"id"`
	BarberID               *int64  `converson:"barbeiro_id"`
	ClientID               *int64  `converson:"cliente_id"`
	BarberShopID           *int64  `converson:"barbearia_id"`
	DateReservation        *string `converson:"data_reserva"`
	DateRservationOriginal *string `converson:"data_reserva_original"`
	StartTime              *string `converson:"horario_inicial"`
	EndTime                *string `converson:"horario_final"`
	Duration               *string `converson:"duracao"`
	Status                 *string `converson:"status"`
}

type Reserva struct {
	DateReservation         *string    `conversor:"data_reserva" json:"data_reserva"`
	DateReservationOriginal *time.Time `conversor:"data_reserva_original" json:"data_reserva_original"`
	StartTime               *string    `conversor:"horario_inicial" json:"horario_inicial"`
	EndTime                 *string    `conversor:"horario_final" json:"horario_final"`
	Duration                *string    `conversor:"duracao" json:"duracao"`
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
