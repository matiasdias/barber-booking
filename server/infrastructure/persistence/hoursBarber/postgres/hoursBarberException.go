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

func (pg *PGHoursBarberException) MarkReservationAsPending(ctx *gin.Context, barberID *int64, hoursExeptionID *string) (bool, error) {
	query := `
	select id from reserva where barbeiro_id = $1 and data_reserva = $2 and status = 'ativo'
	`
	rows, err := pg.DB.QueryContext(ctx, query, barberID, hoursExeptionID) // alterar para data_reserva ao inves de hoursExeptionID
	if err != nil {
		log.Println("Erro ao consultar reservas:", err)
		return false, err
	}
	defer rows.Close()

	var reservaIDs []int64
	for rows.Next() {
		var reservaID int64
		if err := rows.Scan(&reservaID); err != nil {
			log.Println("Failed to scan reservation id: %w:", err)
			return false, err
		}
		reservaIDs = append(reservaIDs, reservaID)
	}

	if len(reservaIDs) == 0 {
		return false, nil
	}

	for _, reservaID := range reservaIDs {
		updateQuery := `
			update reserva set status = 'pendente' where id = $1
		`
		_, err := pg.DB.ExecContext(ctx, updateQuery, reservaID)
		if err != nil {
			log.Println("Failed to update reservation status: %w", err)
			return false, err
		}
	}
	return true, nil
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

func (pg *PGHoursBarberException) ListExeption(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarberExeption, err error) {

	query := `
	select id, barbeiro_id, data_excecao, motivo, data_criacao, data_atualizacao from horario_trabalho_excecao order by data_criacao ASC
	`
	rows, err := pg.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Erro ao consultar as exceções de horario:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var hoursBarber hoursBarber.ListHoursBarberExeption
		err = rows.Scan(&hoursBarber.ID, &hoursBarber.BarberID, &hoursBarber.DateException, &hoursBarber.Reason, &hoursBarber.CreatedAt, &hoursBarber.UpdatedAt)
		if err != nil {
			log.Println("Erro ao consultar as exceções de horario:", err)
			return
		}
		hoursBarbers = append(hoursBarbers, hoursBarber)
	}
	return
}

func (pg *PGHoursBarberException) DeleteExecption(ctx *gin.Context, execptionID *int64) error {

	query := `delete from horario_trabalho_excecao where id = $1`
	_, err := pg.DB.ExecContext(ctx, query, execptionID)
	if err != nil {
		log.Println("Erro ao deletar exceção de horario:", err)
		return err
	}
	return nil
}

func (pg *PGHoursBarberException) MarkReservationAsActive(ctx *gin.Context, barberID *int64, dataReservation *string) (bool, error) {
	query := `
	select id from reserva where barbeiro_id = $1 and data_reserva = $2 and status = 'pendente'
	`
	rows, err := pg.DB.QueryContext(ctx, query, barberID, dataReservation)
	if err != nil {
		log.Println("Erro ao consultar reservas:", err)
		return false, err
	}
	defer rows.Close()

	var reservaIDs []int64
	for rows.Next() {
		var reservaID int64
		if err := rows.Scan(&reservaID); err != nil {
			log.Println("Failed to scan reservation id: %w:", err)
			return false, err
		}
		reservaIDs = append(reservaIDs, reservaID)
	}

	if len(reservaIDs) == 0 {
		return false, nil
	}

	for _, reservaID := range reservaIDs {
		updateQuery := `
			update reserva set status = 'ativo' where id = $1
		`
		_, err := pg.DB.ExecContext(ctx, updateQuery, reservaID)
		if err != nil {
			log.Println("Failed to update reservation status: %w", err)
			return false, err
		}
	}
	return true, nil
}

func (pg *PGHoursBarberException) GetBarberIDByException(ctx *gin.Context, exceptionID *int64) (barberID *int64, dataExecao *string, err error) {

	query := `
	select barbeiro_id, data_excecao from horario_trabalho_excecao where id = $1
	`
	err = pg.DB.QueryRowContext(ctx, query, exceptionID).Scan(&barberID, &dataExecao)
	if err != nil {
		log.Println("Erro ao consultar o barbeiro:", err)
		return
	}
	return
}
