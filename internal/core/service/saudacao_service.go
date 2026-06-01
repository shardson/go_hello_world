package service

import (
	"go_hello_world/internal/core/ports/logger"
	"go_hello_world/internal/core/ports/outbound"
)

type SaudacaoService struct {
	provider outbound.SaudacaoProvider
	logger   logger.Logger
}

func NewGreeterService(
	provider outbound.SaudacaoProvider,
	logger logger.Logger,
) *SaudacaoService {
	return &SaudacaoService{
		provider: provider,
		logger:   logger,
	}
}

func (s *SaudacaoService) Greet() string {
	s.logger.Info("service", "Greet", "Iniciando caso de uso de saudação")
	message := s.provider.GetGreeting()
	s.logger.Info("service", "Greet", "Saudação obtida com sucesso")
	return message
}
