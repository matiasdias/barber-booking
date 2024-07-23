package postgres

import (
	"api/server/infrastructure/persistence/hoursBarber"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type PGHoursBarberException struct {
	DB *sql.DB
}

func (pg *PGHoursBarberException) CreateHoursBarberException(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) error {
	query := `INSERT INTO horario_trabalho_excecao 
	( barbeiro_id, data_excecao, motivo ) 
	VALUES ( $1, $2, $3 ) RETURNING id`

	var hoursExeptionID int64
	err := pg.DB.QueryRowContext(ctx, query,
		hoursException.BarberID,
		hoursException.DateException,
		hoursException.Reason,
	).Scan(&hoursExeptionID)
	if err != nil {
		log.Println("Erro ao inserir e consultar o ID da exceção:", err)
		return err
	}
	return nil
}

func (pg *PGHoursBarberException) MarkReservationAsPending(ctx *gin.Context, barberID *int64, hoursExeptionID *string) error {
	query := `
	select id from reserva where barbeiro_id = $1 and data_reserva = $2 and status = 'ativo'
	`
	rows, err := pg.DB.QueryContext(ctx, query, barberID, hoursExeptionID)
	if err != nil {
		log.Println("Erro ao consultar reservas:", err)
	}
	defer rows.Close()

	var reservaIDs []int64
	for rows.Next() {
		var reservaID int64
		if err := rows.Scan(&reservaID); err != nil {
			log.Println("Failed to scan reservation id: %w:", err)
			return err
		}
		reservaIDs = append(reservaIDs, reservaID)
	}

	for _, reservaID := range reservaIDs {
		updateQuery := `
			update reserva set status = 'pendente' where id = $1
		`
		_, err := pg.DB.ExecContext(ctx, updateQuery, reservaID)
		if err != nil {
			log.Println("failed to update reservation status: %w", err)
		}
	}
	return nil
}

func (pg *PGHoursBarberException) HoursExecptionExists(ctx *gin.Context, hoursException *hoursBarber.HoursBarberException) (exists bool, err error) {
	query :=
		`
	select count(*) from horario_trabalho_excecao 
	where barbeiro_id = $1 and data_excecao = $2
	`
	err = pg.DB.QueryRowContext(ctx, query, hoursException.BarberID, hoursException.DateException).Scan(&exists)
	if err != nil {
		log.Println("Erro ao consultar se existe alguma exceção de horario:", err)
	}

	return
}
