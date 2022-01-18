package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/MSEarn/go-venue-finder/config"
	"github.com/MSEarn/go-venue-finder/db"
	"github.com/MSEarn/go-venue-finder/logz"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	GitCommit string
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to initConfig(), err: %s", err.Error()))
	}

	level := logz.ParseLevel(cfg.Logging.Level)
	logger, err := logz.Init(level, cfg.Logging.Format)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Warn(err.Error())
		}
	}()

	dbPool, err := db.NewPool(&cfg.Mysql)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if e := dbPool.Close(); e != nil {
			fmt.Println(e)
		}
	}()

	app := fiber.New(fiber.Config{
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
		IdleTimeout:           30 * time.Second,
		DisableStartupMessage: true,
		CaseSensitive:         true,
		StrictRouting:         true,
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"git_commit": GitCommit})
	})

	go func() {
		if err := app.Listen(":" + viper.GetString(cfg.Server.Port)); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	logger.Info(fmt.Sprintf("server gitcommit: %s loglevel: %s start: %s", GitCommit, level, viper.GetString(cfg.Server.Port)))
	<-s

	logger.Info("gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		logger.Fatal(err.Error())
	}

}

func initConfig() (*config.Config, error) {
	viper.SetDefault("SERVER.PORT", "9082")
	viper.SetDefault("LOGGING.LEVEL", "debug")
	viper.SetDefault("LOGGING.FORMAT", "json")
	viper.SetDefault("MYSQL.HOST", "localhost")
	viper.SetDefault("MYSQL.PORT", "3306")
	viper.SetDefault("MYSQL.USERNAME", "root")
	viper.SetDefault("MYSQL.PASSWORD", "keep1234")
	viper.SetDefault("MYSQL.DBNAME", "ambassador")
	viper.SetDefault("MYSQL.MAXOPENCONNS", 100)
	viper.SetDefault("MYSQL.MAXCONNLIFETIME", 300)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
