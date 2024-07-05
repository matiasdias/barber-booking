package barber

import (
	"api/server/database"
	"api/server/domain/barber"
	"api/server/utils"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func BarberCreate(ctx *gin.Context, barbers *CreateBarber) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	service := barber.GetService(barber.GetRepository(db))

	dados := &barber.Barber{
		ID:      barbers.ID,
		Name:    barbers.Name,
		Contato: barbers.Contato,
	}
	if *dados.Name == "" || *dados.Contato == "" {
		log.Println("Failed to create barber: missing required fields")
		return errors.New("missing required fields")
	}
	contato, err := utils.FormatContact(dados.Contato)
	if err != nil {
		log.Printf("Failed to format contact: %v", err)
		return err
	}
	dados.Contato = contato

	if err := service.Create(ctx, dados); err != nil {
		log.Printf("Failed to create barber: %v", err)
		return err
	}

	return nil
}

func ListBarber(ctx *gin.Context) (barbers []barber.Barbers, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	defer db.Close()

	service := barber.GetService(barber.GetRepository(db))
	barbers, err = service.List(ctx)
	if err != nil {
		log.Printf("Failed to list barbers: %v", err)
		return nil, err
	}
	return
}
