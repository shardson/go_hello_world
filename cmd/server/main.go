/*package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GO")
}
*/

package main

import (
	"net/http"

	httpadapter "go_hello_world/internal/adapters/in/http"
	logadapter "go_hello_world/internal/adapters/out/logging"
	staticadapter "go_hello_world/internal/adapters/out/static"
	"go_hello_world/internal/core/service"
)

func main() {

	logger := logadapter.NewJSONLogger()
	logger.Info("main", "main", "Inicializando aplicação")

	logger.Info("main", "main", "Criando adaptador de saída")
	greetingProvider := staticadapter.NewGreetingProvider("Hello GO", logger)

	logger.Info("main", "main", "Criando serviço de negócio")
	greeterService := service.NewGreeterService(greetingProvider, logger)

	logger.Info("main", "main", "Criando adaptador HTTP")
	handler := httpadapter.NewHandler(greeterService, logger)

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	logger.Info("main", "main", "Servidor escutando na porta :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("main", "main", "Erro ao iniciar servidor: "+err.Error())
	}
}
