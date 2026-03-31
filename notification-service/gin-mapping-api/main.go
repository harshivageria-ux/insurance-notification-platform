package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"probus-notification-system/gin-mapping-api/internal/infrastructure/db"
	"probus-notification-system/gin-mapping-api/internal/infrastructure/repository"
	"probus-notification-system/gin-mapping-api/internal/service"
	httpTransport "probus-notification-system/gin-mapping-api/internal/transport/http"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := db.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer pool.Close()

	mappingRepo := repository.NewMappingRepository(pool)
	mappingSvc := service.NewMappingService(mappingRepo)

	r := httpTransport.SetupRouter(mappingSvc)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "9100"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Mapping API server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	<-quit
	log.Println("shutting down mapping api server")
	_ = srv.Shutdown(ctx)
}
