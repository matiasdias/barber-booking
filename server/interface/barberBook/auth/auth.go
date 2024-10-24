package auth

import (
	"api/server/aplication/client"
	"api/server/auth"
	jwt "api/server/token"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Handler para redirecionar o usuário para a página de autenticação do Google
func Login(c *gin.Context) {
	url := auth.GoogleOauthConfig.AuthCodeURL("random-state", oauth2.AccessTypeOffline)

	c.Redirect(http.StatusTemporaryRedirect, url)

}

// Handler para tratar o callback do Google
func CallBack(c *gin.Context) {
	// Captura o código da query string
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Code not found",
		})
		return
	}

	// troca o codigo pelo token
	token, err := auth.GoogleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange code for token" + err.Error(),
		})
		return
	}

	// Usa o token para acessar as informações do usuario
	userInfo, err := getUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user info:" + err.Error(),
		})
		return
	}
	err = client.CreateClientFromGoogle(c, &userInfo)
	if err != nil {
		if err == client.ErrCLientExists {
			c.JSON(http.StatusOK, gin.H{
				"message": "Session updated successfully",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save client: " + err.Error(),
		})
		return
	}

	jwtToken, err := jwt.GenerateJWT(userInfo.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate JWT: " + err.Error(),
		})
		return
	}

	// Enviar o token JWT de volta ao cliente
	c.JSON(http.StatusOK, gin.H{
		"message": "Client authenticated successfully",
		"token":   jwtToken,
	})
}

func getUserInfo(accessToken string) (client.CreateClient, error) {
	// Cria uma requisição HTTP para a API de userinfo do Google
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken)
	if err != nil {
		return client.CreateClient{}, err
	}
	defer resp.Body.Close()

	// Decodifica a resposta JSON
	var userInfo client.CreateClient
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return client.CreateClient{}, err
	}
	return userInfo, nil
}

func RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refreshToken"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Valida o refresh token
	email, err := client.IsRefreshTokenValid(c, request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	// Se o refresh token é válido, gera um novo access token
	newAcessToken, err := jwt.GenerateJWT(&email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": newAcessToken})
}

//func getUserPhone(accessToken string) (string, error) {
//	// Define o endpoint da API People para pegar o número de telefone do próprio usuário
//	url := "https://people.googleapis.com/v1/people/me?personFields=phoneNumbers"
//
//	// Cria a requisição HTTP com o token de acesso
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return "", err
//	}
//	req.Header.Set("Authorization", "Bearer "+accessToken)
//
//	// Faz a requisição
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	var userInfo struct {
//		PhoneNumbers []struct {
//			Value string `json:"value"`
//		} `json:"phoneNumbers"`
//	}
//	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
//		return "", err
//	}
//
//	// Verifica se há números de telefone e retorna o primeiro encontrado
//	if len(userInfo.PhoneNumbers) > 0 {
//		return userInfo.PhoneNumbers[0].Value, nil
//	}
//
//	// Retorna um erro ou uma mensagem indicando que nenhum número de telefone foi encontrado
//	return "", fmt.Errorf("nenhum número de telefone encontrado associado à conta do Google do usuário")
//}
//
