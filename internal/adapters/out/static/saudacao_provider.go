package static

import (
	"go_hello_world/internal/core/ports/logger"
)

type SaudacaoProvider struct {
	message string
	logger  logger.Logger
}

func NewGreetingProvider(message string, logger logger.Logger) *SaudacaoProvider {
	logger.Info("static", "NewGreetingProvider", "Instanciando provider estático de saudação")
	return &SaudacaoProvider{
		message: message,
	}
}

func (p *SaudacaoProvider) GetGreeting() string {
	return p.message
}
