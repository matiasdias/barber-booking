package middleware

import (
	"api/server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RequestLogger: middleware para registrar as requisições
func RequestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.Int("StatusCode", statusCode),
			zap.Time("date", startTime),
			zap.String("IP", c.ClientIP()),
			zap.String("Method", c.Request.Method),
			zap.String("Path", c.Request.URL.String()),
			zap.String("User-agent", c.Request.UserAgent()),
		}

		// Verifica se houve um erro no contexto
		if errValue, set := c.Keys["error"]; set {
			if err, ok := errValue.(*utils.CustomError); ok {
				fields = append(fields, []zap.Field{
					zap.Int("error_code", err.Code),
					zap.String("error", err.Error()),
					zap.String("cause", err.Cause.Error()),
					zap.Strings("trace", err.Trace),
				}...)
				logger.Error("falha no tratamento da solicitação", fields...)
				return
			} else if err, ok := errValue.(error); ok {
				// Caso seja um erro comum, loga apenas a mensagem
				fields = append(fields, zap.String("error", err.Error()))
				logger.Error("falha no tratamento da solicitação", fields...)
				return
			}
		}
		if statusCode >= 400 {
			// Log de erro para status de erro
			logger.Warn("falha no tratamento da solicitação", fields...)
		} else {
			logger.Info("solicitação tratada", fields...)
		}
	}
}
