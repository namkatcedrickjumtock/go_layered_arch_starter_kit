package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Iknite-space/itamba-api/internal/api"
	"github.com/Iknite-space/itamba-api/internal/persistence"
	"github.com/Iknite-space/itamba-api/internal/services/echo"
	"github.com/ardanlabs/conf/v3"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func run() error {
	var cfg struct {
		API struct {
			ListenPort string `conf:"env:LISTEN_PORT,required"`
		}
		DB struct {
			User       string `conf:"env:DB_USER,mask,required"`
			Password   string `conf:"env:DB_PASSWORD,mask,required"`
			Host       string `conf:"env:DB_HOST,required"`
			Port       int    `conf:"env:DB_PORT,required"`
			Name       string `conf:"env:DB_NAME,required"`
			DisableTLS bool   `conf:"env:DB_DISABLE_TLS,default:false"`
		}
	}

			// loadDevEnv loads .env file if present
			if _, err := os.Stat(".env"); err == nil {
				err := godotenv.Load()
				if err != nil {
					log.Fatal("Error loading .env file")
				}
			}
			

	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		return fmt.Errorf("parsing config: %w", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	dbInstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer dbInstance.Close()

	repo, err := persistence.NewRepository(dbInstance)
	if err != nil {
		return err
	}

	echoer, err := echo.NewEchoer(repo)
	if err != nil {
		return err
	}

	listerner, err := api.NewAPIListener(echoer)
	if err != nil {
		return err
	}

	listenAddress := fmt.Sprintf("0.0.0.0:%s", cfg.API.ListenPort)

	return listerner.Run(listenAddress)
}
