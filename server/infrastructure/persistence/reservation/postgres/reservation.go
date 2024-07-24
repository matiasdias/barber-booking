package postgres

import (
	"api/server/infrastructure/persistence/reservation"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type PGReservation struct {
	DB *sql.DB
}

func (pg *PGReservation) Create(ctx *gin.Context, reser *reservation.Reservation) error {

	query := `INSERT INTO reserva 
	( barbeiro_id, cliente_id, barbearia_id, data_reserva, horario_inicial_reserva, duracao ) 
	VALUES ( $1, $2, $3, $4, $5, $6 ) RETURNING id`

	var reservationID int64
	err := pg.DB.QueryRowContext(ctx, query,
		reser.BarberID,
		reser.ClientID,
		reser.BarberShopID,
		reser.DateReservation,
		reser.StartTime,
		reser.Duration,
	).Scan(&reservationID)
	if err != nil {
		log.Println("Erro ao inserir e consultar o ID da reserva:", err)
		return err
	}
	return nil
}

func (pg *PGReservation) CheckConflictReservation(ctx *gin.Context, reser *reservation.Reservation) error {
	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM reserva
            WHERE (barbeiro_id = $1 AND data_reserva = $2 AND ($3::time, $3::time + $4::interval) OVERLAPS (horario_inicial_reserva, horario_final))
            OR (cliente_id = $5 AND data_reserva = $2 AND ($3::time, $3::time + $4::interval) OVERLAPS (horario_inicial_reserva, horario_final))
        )
    `
	err := pg.DB.QueryRowContext(ctx, query,
		reser.BarberID,
		reser.DateReservation,
		reser.StartTime,
		reser.Duration,
		reser.ClientID,
	).Scan(&exists)
	if err != nil {
		log.Println("Erro ao consultar se existe conflito de reservas:", err)
		return err
	}
	if exists {
		return errors.New("Conflito de horário: já existe uma reserva para o barbeiro ou uma reserva para o cliente nesse horário especifico")
	}
	return nil
}

func (pg *PGReservation) List(ctx *gin.Context) (reservations []reservation.ReservationList, err error) {

	query := `
		SELECT r.data_reserva, r.horario_inicial_reserva, r.duracao, r.status, r.horario_final, r.data_criacao, r.data_atualizacao,
		r.data_suspensao,ba.nome, ba.cidade, ba.rua, ba.numero_residencia, ba.ponto_referencia, ba.contato, b.nome, b.contato,  
		c.nome, c.email, c.contato
		FROM reserva r
		JOIN barbeiro b ON b.id = r.barbeiro_id
		JOIN cliente c ON c.id = r.cliente_id
		JOIN barbearia ba ON ba.id = r.barbearia_id
		`

	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar reservas:", err)
		return
	}
	defer rows.Close()

	reservationMap := make(map[string]*reservation.ReservationList)
	for rows.Next() {
		var (
			dataReserva           *string
			horarioInicialReserva *string
			duracao               *string
			status                *string
			horarioFinal          *string
			criadoEm              *time.Time
			updatedEm             *time.Time
			dataSuspensao         *time.Time
			shopName              *string
			shopCidade            *string
			shopRua               *string
			shopNumeroResidencia  *int64
			shopPontoReferencia   *string
			shopContato           *string
			barberName            *string
			barberContact         *string
			clientName            *string
			clientEmail           *string
			clientContact         *string
		)

		err := rows.Scan(
			&dataReserva,
			&horarioInicialReserva,
			&duracao,
			&status,
			&horarioFinal,
			&criadoEm,
			&updatedEm,
			&dataSuspensao,
			&shopName,
			&shopCidade,
			&shopRua,
			&shopNumeroResidencia,
			&shopPontoReferencia,
			&shopContato,
			&barberName,
			&barberContact,
			&clientName,
			&clientEmail,
			&clientContact,
		)
		if err != nil {
			log.Println("Erro ao escanear linha:", err)
			return nil, err
		}

		shop := reservation.BarberShop{
			Name:             shopName,
			Cidade:           shopCidade,
			Rua:              shopRua,
			Contact:          shopContato,
			NumeroResidencia: shopNumeroResidencia,
			PontoReferencia:  shopPontoReferencia,
		}

		barber := reservation.Barber{
			Name:    barberName,
			Contact: barberContact,
		}

		client := reservation.Client{
			Contact: clientContact,
			Email:   clientEmail,
			Name:    clientName,
		}

		res := reservation.Reserva{
			DateReservation: dataReserva,
			StartTime:       horarioInicialReserva,
			Duration:        duracao,
			Status:          status,
			EndTime:         horarioFinal,
			CreatedAt:       criadoEm,
			UpdatedAt:       updatedEm,
			DataSuspensao:   dataSuspensao,
		}

		// Construa uma chave única para a combinação de Shop, Barber e Client
		key := fmt.Sprintf("%s-%s-%s", *shop.Name, *barber.Name, *client.Name)

		// Verifica se já existe uma entrada para a mesma combinação de Shop, Barber e Client
		if existing, ok := reservationMap[key]; ok {
			existing.Reservations = append(existing.Reservations, res)
		} else {
			// Se não encontrado, cria uma nova entrada
			reservationMap[key] = &reservation.ReservationList{
				Shop:         shop,
				Barber:       barber,
				Client:       client,
				Reservations: []reservation.Reserva{res},
			}
		}
	}

	// Converte o mapa para uma lista
	for _, v := range reservationMap {
		reservations = append(reservations, *v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
