package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/elangreza14/moviefestival/cmd/http/config"
	"github.com/elangreza14/moviefestival/cmd/http/routes"
	"github.com/elangreza14/moviefestival/controller"
	"github.com/elangreza14/moviefestival/middleware"
	"github.com/elangreza14/moviefestival/repository"
	"github.com/elangreza14/moviefestival/service"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// setup env

	cfg, err := config.LoadConfig()
	errChecker(err)

	// logger
	logger, err := Logger(cfg)
	errChecker(err)
	defer logger.Sync()

	ctx := context.Background()

	// db connection
	db, err := DB(ctx, cfg, logger)
	errChecker(err)

	// dependency injection
	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)
	movieRepository := repository.NewMovieRepository(db, db)
	movieViewRepository := repository.NewMovieViewRepository(db, db)
	genreRepository := repository.NewGenreRepository(db)

	authService := service.NewAuthService(userRepository, tokenRepository)
	movieService := service.NewMovieService(movieRepository, movieViewRepository)
	genreService := service.NewGenreService(genreRepository)

	authController := controller.NewAuthController(authService)
	movieController := controller.NewMovieController(movieService)
	genreController := controller.NewGenreController(genreService)

	if cfg.ENV != "DEVELOPMENT" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	// cors middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.ALLOWED_ORIGINS
	router.Use(cors.New(corsConfig))

	// logger middleware
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// pinger
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authMiddleware := middleware.NewAuthMiddleware(authService)

	// group api
	apiGroup := router.Group("/api")
	routes.AuthRoute(apiGroup, authController)
	routes.MovieRoute(apiGroup, movieController, authMiddleware)
	routes.GenreRoute(apiGroup, genreController, authMiddleware)

	srv := &http.Server{
		Addr:    cfg.HTTP_PORT,
		Handler: router.Handler(),
	}

	go func() {
		logger.Info("listening http", zap.String("port", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	waitOperations := gracefulShutdown(context.Background(), logger, time.Second*5,
		func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
		func(ctx context.Context) error {
			db.Close()
			return nil
		})

	<-waitOperations
}

func errChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func DB(ctx context.Context, cfg *config.Config, logger *zap.Logger) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.POSTGRES_USER,
		cfg.POSTGRES_PASSWORD,
		cfg.POSTGRES_HOSTNAME,
		cfg.POSTGRES_PORT,
		cfg.POSTGRES_DB,
		cfg.POSTGRES_SSL,
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	logLevel, err := tracelog.LogLevelFromString(logger.Level().String())
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxzap.NewLogger(logger),
		LogLevel: logLevel,
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func Logger(cfg *config.Config) (*zap.Logger, error) {
	logger := zap.NewExample(zap.IncreaseLevel(zap.InfoLevel))
	if cfg.ENV != "DEVELOPMENT" {
		return zap.NewProduction()
	}

	return logger, nil
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, logger *zap.Logger, timeout time.Duration, ops ...operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		logger.Info("shutting down")

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		go func() {
			<-ctx.Done()
			logger.Info("force quit the app")
			wait <- struct{}{}
		}()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			go func(key int, op operation) {
				defer wg.Done()
				processName := fmt.Sprintf("process %d", key)

				if err := op(ctx); err != nil {
					logger.Error(processName, zap.Error(err), zap.Bool("success", true))
					return
				}

				logger.Info(processName, zap.Bool("success", true))
			}(key, op)
		}

		wg.Wait()
		cancel()
	}()

	return wait
}
