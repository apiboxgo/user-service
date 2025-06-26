package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/api/user"
	"user-service/api_init"
	_ "user-service/docs"
)

func routes(config *api_init.InitGlobalStruct) *gin.Engine {
	r := gin.Default()

	user.InitUserRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func main() {

	fmt.Println("Init main ...")
	err := api_init.MainInit("")
	if err != nil {
		log.Fatal(err)
	}

	router := routes(api_init.InitGlobal)
	srv := &http.Server{
		Addr:    ":" + api_init.InitGlobal.Cfg.ServerPort,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	log.Println("Server started on port", api_init.InitGlobal.Cfg.ServerPort)

	<-quit // Ожидаем сигнала завершения

	log.Println("Shutting down server...")

	// Контекст с таймаутом (макс. 5 секунд на завершение)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
