package test

import (
	"api/server/aplication/hoursBarber"
	"api/server/database"
	domain "api/server/domain/hoursBarber"
	"bytes"
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCreateHoursBarber(t *testing.T) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	service := domain.GetService(domain.GetRepository(db))

	router := gin.Default()
	router.POST("hoursBarber/create", func(ctx *gin.Context) {
		var hours domain.HoursBarber
		if err := ctx.ShouldBindJSON(&hours); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.Create(ctx, &hours); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "barber hours added successfully"})
		}
	})

	// Caso de teste 01: Dados válidos
	hoursValid := hoursBarber.CreateHoursBarber{
		BarberID:       Int64Ptr(1),
		DayOfWeek:      StringPtr("Terça-Feira"),
		StartTime:      StringPtr("08:00:00"),
		LunchStartTime: StringPtr("12:00:00"),
		LunchEndTime:   StringPtr("13:00:00"),
		EndTime:        StringPtr("18:00:00"),
	}

	hoursValidJSON, _ := json.Marshal(hoursValid)
	req, _ := http.NewRequest("POST", "/hoursBarber/create", bytes.NewBuffer(hoursValidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "barber hours added successfully")

	// Caso de teste 02: Horário de início maior que horário de fim
	hoursInvalidTime := hoursBarber.CreateHoursBarber{
		BarberID:       Int64Ptr(1),
		DayOfWeek:      StringPtr("Terça-Feira"),
		StartTime:      StringPtr("08:00:00"),
		LunchStartTime: StringPtr("12:00:00"),
		LunchEndTime:   StringPtr("13:00:00"),
		EndTime:        StringPtr("07:00:00"),
	}

	hoursInvalidTimeJSON, _ := json.Marshal(hoursInvalidTime)
	req, _ = http.NewRequest("POST", "/hoursBarber/create", bytes.NewBuffer(hoursInvalidTimeJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "start time must be before end time")

	// Caso de teste 03: Falta o ID do barbeiro
	hoursMissingID := hoursBarber.CreateHoursBarber{
		DayOfWeek:      StringPtr("Terça-Feira"),
		StartTime:      StringPtr("08:00:00"),
		LunchStartTime: StringPtr("12:00:00"),
		LunchEndTime:   StringPtr("13:00:00"),
		EndTime:        StringPtr("07:00:00"),
	}

	hoursMissingIDJSON, _ := json.Marshal(hoursMissingID)
	req, _ = http.NewRequest("POST", "/hoursBarber/create", bytes.NewBuffer(hoursMissingIDJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "missing or invalid BarberID")

	// Caso de teste 04: Falta o dia da semana
	hoursMissingDayWeek := hoursBarber.CreateHoursBarber{
		BarberID:       Int64Ptr(1),
		StartTime:      StringPtr("08:00:00"),
		LunchStartTime: StringPtr("12:00:00"),
		LunchEndTime:   StringPtr("13:00:00"),
		EndTime:        StringPtr("07:00:00"),
	}

	hoursMissingDayWeekJSON, _ := json.Marshal(hoursMissingDayWeek)
	req, _ = http.NewRequest("POST", "/hoursBarber/create", bytes.NewBuffer(hoursMissingDayWeekJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "missing or invalid day of the week")

	// Caso de teste 05: Horário de inicio do almoço maior que o horario de fim do almoço
	hoursInvalidLunchTime := hoursBarber.CreateHoursBarber{
		BarberID:       Int64Ptr(1),
		DayOfWeek:      StringPtr("Terça-Feira"),
		StartTime:      StringPtr("08:00:00"),
		LunchStartTime: StringPtr("13:00:00"),
		LunchEndTime:   StringPtr("12:00:00"),
		EndTime:        StringPtr("07:00:00"),
	}

	hoursInvalidLunchTimeJSON, _ := json.Marshal(hoursInvalidLunchTime)
	req, _ = http.NewRequest("POST", "/hoursBarber/create", bytes.NewBuffer(hoursInvalidLunchTimeJSON))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "LunchStartTime must be before LunchEndTime")

}

func StringPtr(s string) *string {
	return &s
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
