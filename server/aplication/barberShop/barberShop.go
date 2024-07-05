package barberShop

import (
	"api/server/database"
	"api/server/domain/barberShop"
	"api/server/utils"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func BarberShopCreate(c *gin.Context, req *CreateBarberShop) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	service := barberShop.GetService(barberShop.GetRepository(db))

	dados := &barberShop.BarberShop{
		ID:               req.ID,
		Name:             req.Name,
		Cidade:           req.Cidade,
		Rua:              req.Rua,
		NumeroResidencia: req.NumeroResidencia,
		PontoReferencia:  req.PontoReferencia,
		Contato:          req.Contato,
	}
	if *dados.Name == "" || *dados.Cidade == "" || *dados.Rua == "" || *dados.NumeroResidencia == 0 || *dados.PontoReferencia == "" || *dados.Contato == "" {
		log.Println("Failed to create barber-shop: missing required fields")
		return errors.New("missing required fields")
	}
	contato, err := utils.FormatContact(dados.Contato)
	if err != nil {
		log.Printf("Failed to format contact: %v", err)
		return err
	}
	dados.Contato = contato

	if err := service.Create(c, dados); err != nil {
		log.Printf("Failed to create barber-shop: %v", err)
		return err
	}

	return
}

func ListBshop(c *gin.Context) (barberShops []barberShop.ListBarberShop, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	service := barberShop.GetService(barberShop.GetRepository(db))
	barberShops, err = service.List(c)
	if err != nil {
		log.Printf("Failed to get barber-shops: %v", err)
		return
	}
	return
}
