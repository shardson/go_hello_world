package http

import (
	"fmt"
	nethttp "net/http"

	"go_hello_world/internal/core/ports/inbound"
	"go_hello_world/internal/core/ports/logger"
)

type Handler struct {
	greeter inbound.SaudacaoUseCase
	logger  logger.Logger
}

func NewHandler(greeter inbound.SaudacaoUseCase, logger logger.Logger) *Handler {
	return &Handler{
		greeter: greeter,
		logger:  logger,
	}
}

func (h *Handler) ServeHTTP(w nethttp.ResponseWriter, r *nethttp.Request) {
	h.logger.Info("http", "ServeHTTP", "Recebida requisição HTTP na rota /")
	message := h.greeter.Greet()
	h.logger.Info("http", "ServeHTTP", "Escrevendo resposta HTTP com sucesso")
	fmt.Fprint(w, message)
}
