package barberShop

import (
	"api/server/database"
	"api/server/domain/barberShop"
	"api/server/utils"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func BarberShopCreate(ctx *gin.Context, req *CreateBarberShop) (err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	service := barberShop.GetService(barberShop.GetRepository(db))

	dados := &barberShop.BarberShop{
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

	if err := service.Create(ctx, dados); err != nil {
		log.Printf("Failed to create barber-shop: %v", err)
		return err
	}

	return
}

func ListBshop(ctx *gin.Context) (barberShops []*ListBarbserShop, err error) {
	db, err := database.Connection()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	var (
		service = barberShop.GetService(barberShop.GetRepository(db))
		dados   []barberShop.ListBarberShop
	)

	if dados, err = service.List(ctx); err != nil {
		log.Printf("Failed to list barber-shops: %v", err)
		return
	}

	for i := range dados {
		barber := &ListBarbserShop{
			ID:               dados[i].ID,
			Name:             dados[i].Name,
			Cidade:           dados[i].Cidade,
			Rua:              dados[i].Rua,
			NumeroResidencia: dados[i].NumeroResidencia,
			PontoReferencia:  dados[i].PontoReferencia,
			Contato:          dados[i].Contato,
			CriadoEm:         dados[i].CriadoEm,
			UpdateEm:         dados[i].UpdateEm,
		}
		barberShops = append(barberShops, barber)
	}
	return
}
