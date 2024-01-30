package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oluijks/golang-starter/server/api/handlers"
	"github.com/oluijks/golang-starter/server/internal/config"
	"go.uber.org/fx"
)

func registerRoutes(
	engine *gin.Engine,
	ph *handlers.PingHandler,
	uh *handlers.AccountHandler,
	ah *handlers.AuthHandler,
	config *config.Config) {
	apiV1 := engine.Group("/api/v1")
	{
		apiV1.GET("/ping", ph.Ping)

		auth := apiV1.Group("/auth")
		auth.POST("/", ah.SignIn)

		account := apiV1.Group("/account").Use(authMiddleware(config))
		account.GET("/", uh.ListAccounts)
		account.GET("/:id", uh.ListAccount)
		account.POST("/", uh.CreateAccount)
		account.PATCH("/:id", uh.UpdateAccount)
		account.DELETE("/:id", uh.DeleteAccount)
	}
}

func NewGinEngine(
	lc fx.Lifecycle,
	ph *handlers.PingHandler,
	uh *handlers.AccountHandler,
	ah *handlers.AuthHandler,
	config *config.Config,
) *gin.Engine {
	c, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	engine := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.FrondendURL}
	engine.Use(cors.New(corsConfig))

	httpServer := &http.Server{
		Addr:           net.JoinHostPort(config.HTTPServerHost, config.HTTPServerPort),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1mb
	}

	registerRoutes(engine, ph, uh, ah, config)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Listen: %s\n", err)
				}
			}()
			<-c.Done()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			stop()
			log.Println("Shutting down gracefully, press Ctrl+C again to force")
			c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := httpServer.Shutdown(c); err != nil {
				log.Fatal("Server forced to shutdown: ", err)
			}
			log.Println("Server exiting")

			return nil
		},
	})

	return engine
}
