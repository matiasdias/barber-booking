package reservation

import (
	"api/server/aplication/reservation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateReservation godoc
// @Summary Criação das reservas
// @Description Cria uma nova reserva para um barbeiro
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param barber body CreateReservation true "Create reservation"
// @Success 200 "Sem conteúdo"
// @Router /barber/reservation/create [post]
func Create(c *gin.Context) {
	var (
		err error
		req reservation.CreateReservation
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = reservation.Create(c.Copy(), &req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Reservation created successfully",
	})

}

func List(c *gin.Context) {
	var (
		err error
	)
	reservations, err := reservation.List(c.Copy())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, reservations)
}

// UpdateReservationReq godoc
// @Summary Atualização das reservas
// @Description Atualiza uma nova reserva para um barbeiro
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param id path int true "Reservation ID"
// @Param reservation body reservation.UpdateReservationReq true "Update reservation"
// @Success 200 "Sem conteúdo"
// @Router /barber/reservation/update/{id} [put]
func Update(c *gin.Context) {
	var (
		err error
		req reservation.UpdateReservationReq
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reservation ID",
		})
		return
	}

	// Call the application layer function
	if err = reservation.Update(c.Copy(), &id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Reservation updated successfully",
	})
}
