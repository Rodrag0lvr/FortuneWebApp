package main

import (
	uc "fortuna-express-web/pkg/domain/usecases"
	repository "fortuna-express-web/pkg/repositories"
	"fortuna-express-web/pkg/web"
	handlers "fortuna-express-web/pkg/web/handlers"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear una instancia de Gin
	r := gin.Default()
	logger := slog.Default()
	liquidationRepo := repository.NewInMemoryLiquidationRepository()
	descriptionRepo := repository.NewInMemoryDescriptionRepository()
	aditionRepo := repository.NewInMemoryAditionRepository()
	uc := uc.NewLiquidationUseCase(liquidationRepo, descriptionRepo, aditionRepo)
	handler := handlers.NewLiquidationsHandler(logger, uc)
	// Configurar el router desde infrastructure
	web.SetupRouter(r, handler)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
