package client

import (
	"api/server/infrastructure/persistence/client"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo IClient
}

// GetService retorna um servico para acesso a funções de auxilio à lógica de negócio
func GetService(r IClient) *Service {
	return &Service{repo: r}
}

// GetRepository retorna um repositório para acesso à camada de dados
func GetRepository(db *sql.DB) IClient {
	return newRepository(db)
}

func (s *Service) Create(ctx *gin.Context, clients *Client) (err error) {
	dados := &client.Client{
		ID:       clients.ID,
		Name:     clients.Name,
		Email:    clients.Email,
		Contato:  clients.Contato,
		PassWord: clients.PassWord,
	}
	return s.repo.Create(ctx, dados)
}

func (s *Service) List(ctx *gin.Context) (clients []Clients, err error) {

	clientes, err := s.repo.List(ctx)
	if err != nil {
		return
	}
	clients = make([]Clients, len(clientes))
	for i := range clientes {
		var cliente Clients
		cliente.ID = clientes[i].ID
		cliente.Name = clientes[i].Name
		cliente.Email = clientes[i].Email
		cliente.Contato = clientes[i].Contato
		cliente.PassWord = clientes[i].PassWord
		cliente.CriadoEm = clientes[i].CriadoEm
		cliente.UpdateEm = clientes[i].UpdateEm
		clients[i] = cliente
	}
	return
}
