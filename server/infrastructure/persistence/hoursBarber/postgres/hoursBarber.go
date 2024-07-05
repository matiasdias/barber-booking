package postgres

import (
	"api/server/infrastructure/persistence/hoursBarber"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type PGHoursBarber struct {
	DB *sql.DB
}

func (pg *PGHoursBarber) Create(ctx *gin.Context, hours *hoursBarber.HoursBarber) error {

	query := `INSERT INTO horario_trabalho_barbeiro
	( barbeiro_id, dia_semana, horario_inicio, horario_almoco_inicio, horario_almoco_fim, horario_fim ) 
	VALUES ( $1, $2, $3, $4, $5, $6 ) RETURNING id`
	var hoursID int64
	err := pg.DB.QueryRowContext(ctx, query,
		hours.BarberID,
		hours.DayOfWeek,
		hours.StartTime,
		hours.LunchStartTime,
		hours.LunchEndTime,
		hours.EndTime).Scan(&hoursID)

	if err != nil {
		log.Println("Erro ao inserir e consultar o ID do horário do barbeiro:", err)
		return err
	}
	return nil
}

func (pg *PGHoursBarber) CheckConflitHoursBarber(ctx *gin.Context, hours *hoursBarber.HoursBarber) (bool, error) {
	query := `SELECT COUNT(*) FROM horario_trabalho_barbeiro 
	WHERE barbeiro_id = $1 AND dia_semana = $2
	AND (
		($3::time, $4::time) OVERLAPS (horario_inicio, horario_fim)
		OR ($3::time, $4::time) OVERLAPS (horario_almoco_inicio, horario_almoco_fim)
	)`
	var count int
	err := pg.DB.QueryRowContext(ctx, query, hours.BarberID, hours.DayOfWeek, hours.StartTime, hours.EndTime).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pg *PGHoursBarber) List(ctx *gin.Context) (hoursBarbers []hoursBarber.ListHoursBarber, err error) {
	query := `SELECT b.nome, b.contato, htb.dia_semana, htb.horario_inicio,
	 htb.horario_almoco_inicio, htb.horario_almoco_fim,
	 htb.horario_fim::time, htb.criadoem, htb.updatedem
	 FROM horario_trabalho_barbeiro htb
	 JOIN barbeiro b ON b.id = htb.barbeiro_id
	 `
	rows, err := pg.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("Erro ao consultar os horarios:", err)
		return
	}

	defer rows.Close()

	hours := make(map[string]*hoursBarber.ListHoursBarber)

	for rows.Next() {
		var (
			NomeBarberio       *string
			ContatoBarberbeiro *string
			diaSemana          *string
			horarioInicio      *string
			almocoInicio       *string
			almocoFim          *string
			horarioFim         *string
			criadoEm           *time.Time
			updatedEm          *time.Time
		)
		err = rows.Scan(
			&NomeBarberio,
			&ContatoBarberbeiro,
			&diaSemana,
			&horarioInicio,
			&almocoInicio,
			&almocoFim,
			&horarioFim,
			&criadoEm,
			&updatedEm,
		)
		if err != nil {
			log.Println("Erro ao consultar horários do barbeiro:", err)
			return
		}

		barber := hoursBarber.Barber{
			Name:    NomeBarberio,
			Contato: ContatoBarberbeiro,
		}

		hBarber := hoursBarber.HoursBarbers{
			DayOfWeek:      diaSemana,
			StartTime:      horarioInicio,
			LunchStartTime: almocoInicio,
			LunchEndTime:   almocoFim,
			EndTime:        horarioFim,
			CreatedAt:      criadoEm,
			UpdatedAt:      updatedEm,
		}

		key := fmt.Sprintf("%s", *barber.Name)

		if existingHours, ok := hours[key]; ok {
			existingHours.HourBarbers = append(existingHours.HourBarbers, hBarber)
		} else {
			hours[key] = &hoursBarber.ListHoursBarber{
				Barber:      barber,
				HourBarbers: []hoursBarber.HoursBarbers{hBarber},
			}
		}
	}

	for _, v := range hours {
		hoursBarbers = append(hoursBarbers, *v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hoursBarbers, nil
}
