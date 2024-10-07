package hoursBarber

import "time"

type CreateHoursBarber struct {
	BarberID       *int64  `conversor:"barbeiro_id" json:"barbeiro_id" binding:"required"`
	DayOfWeek      *string `conversor:"dia_semana" json:"dia_semana" binding:"required"`
	StartTime      *string `conversor:"horario_inicial" json:"horario_inicial" binding:"required"`
	LunchStartTime *string `conversor:"horario_almoco_inicial" json:"horario_almoco_inicial" binding:"required"`
	LunchEndTime   *string `conversor:"horario_almoco_final" json:"horario_almoco_final" binding:"required"`
	EndTime        *string `conversor:"horario_final" json:"horario_final" binding:"required"`
}

type CreateException struct {
	BarberID      *int64  `conversor:"barbeiro_id" json:"barbeiro_id" binding:"required"`
	DateException *string `conversor:"data_execeção" json:"data_execeção" binding:"required"`
	Reason        *string `conversor:"motivo" json:"motivo"`
}

type HoursBarbers struct {
	DayOfWeek      *string    `json:"dia_semana" conversor:"dia_semana"`
	StartTime      *string    `json:"horario_inicial" conversor:"horario_inicial"`
	LunchStartTime *string    `json:"horario_almoco_inicial" conversor:"horario_almoco_inicial"`
	LunchEndTime   *string    `json:"horario_almoco_final" conversor:"horario_almoco_final"`
	EndTime        *string    `json:"horario_final" conevrsor:"horario_final"`
	CreatedAt      *time.Time `json:"data_criacao" conversor:"data_criacao"`
	UpdatedAt      *time.Time `json:"data_atualizacao" conversor:"data_atualizacao"`
}
type Barber struct {
	Name    *string `conversor:"nome" json:"nome"`
	Contato *string `conversor:"contato" json:"contato"`
}

type ListHoursBarber struct {
	Barber      Barber         `conversor:"barbeiro" json:"barbeiro"`
	HourBarbers []HoursBarbers `conversor:"horario_trabalho" json:"horario_trabalho"`
}

type ListHoursBarberExeption struct {
	ID            *int64     `conversor:"id" json:"id"`
	BarberID      *int64     `conversor:"barbeiro_id" json:"barbeiro_id"`
	DateException *string    `conversor:"data_execeção" json:"data_execeção"`
	Reason        *string    `conversor:"motivo" json:"motivo"`
	CreatedAt     *time.Time `conversor:"data_criacao" json:"data_criacao"`
	UpdatedAt     *time.Time `conversor:"data_atualizacao" json:"data_atualizacao"`
}
