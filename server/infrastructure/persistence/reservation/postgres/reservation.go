package postgres

import (
	"api/server/infrastructure/persistence/reservation"
	"context"
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
	( barbeiro_id, cliente_id, barbearia_id, servico_id, data_reserva, horario_inicial_reserva ) 
	VALUES ( $1, $2, $3, $4, $5, $6 ) RETURNING id`

	var reservationID int64
	err := pg.DB.QueryRowContext(ctx, query,
		reser.BarberID,
		reser.ClientID,
		reser.BarberShopID,
		reser.ServiceID,
		reser.DateReservation,
		reser.StartTime,
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
            FROM reserva r
            JOIN servico s ON r.servico_id = s.id
            WHERE (r.barbeiro_id = $1 AND r.data_reserva = $2 AND ($3::time, $3::time + s.duracao) OVERLAPS (r.horario_inicial_reserva, r.horario_final))
            OR (r.cliente_id = $4 AND r.data_reserva = $2 AND ($3::time, $3::time + s.duracao) OVERLAPS (r.horario_inicial_reserva, r.horario_final))
        )
    `
	err := pg.DB.QueryRowContext(ctx, query,
		reser.BarberID,
		reser.DateReservation,
		reser.StartTime,
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
		SELECT r.data_reserva, r.horario_inicial_reserva, r.status, r.horario_final, r.data_criacao, r.data_atualizacao,
		r.data_reserva_original, ba.nome, ba.cidade, ba.rua, ba.numero_residencia, ba.ponto_referencia, ba.contato, b.nome, b.contato,  
		c.nome, c.email, c.contato, s.nome, s.preco, s.duracao
		FROM reserva r
		JOIN barbeiro b ON b.id = r.barbeiro_id
		JOIN cliente c ON c.id = r.cliente_id
		JOIN barbearia ba ON ba.id = r.barbearia_id
		JOIN servico s ON s.id = r.servico_id
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
			status                *string
			horarioFinal          *string
			criadoEm              *time.Time
			updatedEm             *time.Time
			dataReservaOriginal   *time.Time
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
			serviceName           *string
			servicePrice          *float64
			serviceDuration       *string
		)

		err := rows.Scan(
			&dataReserva,
			&horarioInicialReserva,
			&status,
			&horarioFinal,
			&criadoEm,
			&updatedEm,
			&dataReservaOriginal,
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
			&serviceName,
			&servicePrice,
			&serviceDuration,
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
		service := reservation.Service{
			Name:     serviceName,
			Price:    servicePrice,
			Duration: serviceDuration,
		}

		res := reservation.Reserva{
			DateReservation:         dataReserva,
			StartTime:               horarioInicialReserva,
			Status:                  status,
			EndTime:                 horarioFinal,
			CreatedAt:               criadoEm,
			UpdatedAt:               updatedEm,
			DateReservationOriginal: dataReservaOriginal,
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
				Service:      service,
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

func (pg *PGReservation) CheckExceptionForBarber(ctx *gin.Context, barberID *int64, dataReservation *string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1 
		FROM horario_trabalho_excecao 
		WHERE barbeiro_id = $1 
		AND data_excecao = $2)`
	var exists bool
	err := pg.DB.QueryRowContext(ctx, query, barberID, dataReservation).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		log.Printf("Failed to check exception for barber: %v", err)
		return false, err
	}
	return exists, nil

}

func (pg *PGReservation) UpdateReservation(ctx context.Context, reservationID *int64, updateReq *reservation.Reservation) error {
	query := `
        UPDATE reserva
        SET 
            barbeiro_id = COALESCE($2, barbeiro_id),
            data_reserva_original = CASE 
                WHEN data_reserva_original IS NULL THEN data_criacao 
                ELSE data_reserva_original 
            END,
            data_reserva = COALESCE($3, data_reserva),
            horario_inicial_reserva = COALESCE($4, horario_inicial_reserva),
            status = CASE 
                WHEN status = 'pendente' AND $3 IS NOT NULL THEN 'ativo'
                WHEN status = 'pendente' AND $4 IS NOT NULL THEN 'ativo'
                WHEN status = 'pendente' AND $2 IS NOT NULL THEN 'ativo'
                ELSE COALESCE($5, status)
            END,
            data_atualizacao = now()
        WHERE id = $1
        RETURNING status
    `
	var newStatus string
	err := pg.DB.QueryRowContext(ctx, query, reservationID,
		updateReq.BarberID,
		updateReq.DateReservation,
		updateReq.StartTime,
		updateReq.Status).Scan(&newStatus)
	if err != nil {
		return err
	}

	return nil
}
