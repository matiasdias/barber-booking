package hoursBarber

import "time"

type HoursBarber struct {
	ID             *int64     `conversor:"id" `
	BarberID       *int64     `conversor:"barbeiro_id"`
	DayOfWeek      *string    `conversor:"dia_semana"`
	StartTime      *string    `conversor:"horario_inicial"`
	LunchStartTime *string    `conversor:"horario_almoco_inicial"`
	LunchEndTime   *string    `conversor:"horario_almoco_final"`
	EndTime        *string    `conevrsor:"horario_final"`
	CreatedAt      *time.Time `conversor:"criado_em"`
	UpdatedAt      *time.Time `conversor:"updated_em"`
}

type HoursBarbers struct {
	DayOfWeek      *string    `json:"dia_semana" conversor:"dia_semana"`
	StartTime      *string    `json:"horario_inicial" conversor:"horario_inicial"`
	LunchStartTime *string    `json:"horario_almoco_inicial" conversor:"horario_almoco_inicial"`
	LunchEndTime   *string    `json:"horario_almoco_final" conversor:"horario_almoco_final"`
	EndTime        *string    `json:"horario_final" conevrsor:"horario_final"`
	CreatedAt      *time.Time `json:"criado_em" conversor:"criado_em"`
	UpdatedAt      *time.Time `json:"updated_em" conversor:"updated_em"`
}
type Barber struct {
	Name    *string `conversor:"nome" json:"nome"`
	Contato *string `conversor:"contato" json:"contato"`
}

type ListHoursBarber struct {
	Barber      Barber         `conversor:"barbeiro" json:"barbeiro"`
	HourBarbers []HoursBarbers `conversor:"horario_trabalho" json:"horario_trabalho"`
}
