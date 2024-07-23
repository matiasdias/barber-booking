package hoursBarber

type CreateHoursBarber struct {
	ID             *int64  `conversor:"id" json:"id"`
	BarberID       *int64  `conversor:"barbeiro_id" json:"barbeiro_id" binding:"required"`
	DayOfWeek      *string `conversor:"dia_semana" json:"dia_semana" binding:"required"`
	StartTime      *string `conversor:"horario_inicial" json:"horario_inicial" binding:"required"`
	LunchStartTime *string `conversor:"horario_almoco_inicial" json:"horario_almoco_inicial" binding:"required"`
	LunchEndTime   *string `conversor:"horario_almoco_final" json:"horario_almoco_final" binding:"required"`
	EndTime        *string `conversor:"horario_final" json:"horario_final" binding:"required"`
}

type CreateException struct {
	ID            *int64  `conversor:"id" json:"id"`
	BarberID      *int64  `conversor:"barbeiro_id" json:"barbeiro_id" binding:"required"`
	DateException *string `conversor:"data_execeção" json:"data_execeção" binding:"required"`
	Reason        *string `conversor:"motivo" json:"motivo"`
}
